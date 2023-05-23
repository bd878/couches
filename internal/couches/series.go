package couches

import "os"
import "strconv"
import "encoding/csv"

type Series struct {
  path string
  records [][]string
}

func NewSeries() *Series {
  return &Series{}
}

func (s *Series) Load(path string) error {
  s.path = path
  fd, err := os.Open(path)
  if err != nil {
    return err
  }

  r := csv.NewReader(fd)
  records, err := r.ReadAll()
  if err != nil {
    return err
  }

  _, err = strconv.ParseInt(records[0][0], 10, 64) // first is a heading
  if err != nil {
    s.records = records[1:]
  } else {
    s.records = records
  }

  return nil
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

func (s *Series) Slice(fromTs int64, toTs int64) (interface{}, error) {
  var records [][]string

  for _, rec := range s.records {
    row := NewRow(rec)
    ts := row.Ts()
    if fromTs <= ts && toTs >= ts {
      records = append(records, []string(row))
    }
  }

  result := &Series{}
  result.records = records
  result.path = ""

  return result, nil
}