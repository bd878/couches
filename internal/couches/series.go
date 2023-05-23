package couches

import "fmt"
import "os"
import "strconv"
import "encoding/csv"

type Series struct {
  path string
  heading []string
  records [][]string
}

func NewSeries() *Series {
  return &Series{}
}

func (s *Series) Path() string {
  return s.path
}

func (s *Series) Load(path string) {
  s.path = path
  fd, err := os.Open(path)
  if err != nil {
    panic(err)
  }

  r := csv.NewReader(fd)
  records, err := r.ReadAll()
  if err != nil {
    panic(err)
  }

  _, err = strconv.ParseInt(records[0][0], 10, 64) // first is a heading
  if err != nil {
    s.heading = records[0]
    s.records = records[1:]
  } else {
    s.records = records
  }
}

func (s *Series) Count() int {
  return len(s.records)
}

func (s *Series) FirstRow() Row {
  return NewRow(s.records[1])
}

func (s *Series) LastRow() Row {
  return NewRow(s.records[s.Count() - 1])
}

func (s *Series) TsFrom() int64 {
  return s.FirstRow().Ts()
}

func (s *Series) TsTo() int64 {
  return s.LastRow().Ts()
}

func (s *Series) Slice(path string, fromTs int64, toTs int64) interface{} {
  var records [][]string

  for _, rec := range s.records {
    row := NewRow(rec)
    ts := row.Ts()
    if fromTs <= ts && toTs >= ts {
      records = append(records, []string(row))
    }
  }

  result := &Series{}
  result.heading = s.heading
  result.records = records
  result.path = path

  return result
}

func (s *Series) Save(fd *os.File) {
  w := csv.NewWriter(fd)

  err := w.Write(s.heading)
  if err != nil {
    panic(err)
  }

  err = w.WriteAll(s.records)
  if err != nil {
    panic(err)
  }

  err = w.Error()
  if err != nil {
    panic(err)
  }

  err = fd.Sync()
  if err != nil {
    panic(err)
  }
}