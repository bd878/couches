package couches

import "os"
import "strconv"
import "encoding/csv"

type CSVSeries struct {
  path string
  heading []string
  records [][]string
  fd *os.File
}

func NewCSVSeries() *CSVSeries {
  return &CSVSeries{}
}

func (s *CSVSeries) Path() string {
  return s.path
}

func (s *CSVSeries) Load(path string) {
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

  s.fd = fd
}

func (s *CSVSeries) Close() {
  s.fd.Close()
}

func (s *CSVSeries) Count() int {
  return len(s.records)
}

func (s *CSVSeries) FirstRow() LogRow {
  return NewCSVRow(s.records[1])
}

func (s *CSVSeries) LastRow() LogRow {
  return NewCSVRow(s.records[s.Count() - 1])
}

func (s *CSVSeries) TsFrom() int64 {
  return s.FirstRow().Ts()
}

func (s *CSVSeries) TsTo() int64 {
  return s.LastRow().Ts()
}

func (s *CSVSeries) Slice(path string, fromTs int64, toTs int64) interface{} {
  var records [][]string

  for _, rec := range s.records {
    row := NewCSVRow(rec)
    ts := row.Ts()
    if fromTs <= ts && toTs >= ts {
      records = append(records, []string(row))
    }
  }

  result := &CSVSeries{}
  result.heading = s.heading
  result.records = records
  result.path = path

  return result
}

func (s *CSVSeries) Save(fd *os.File) {
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