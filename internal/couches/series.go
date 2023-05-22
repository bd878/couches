package couches

import "os"
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

  s.records = records
  return nil
}

func (s *Series) Count() int {
  return len(s.records)
}
