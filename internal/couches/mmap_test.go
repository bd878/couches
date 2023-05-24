package couches

import "testing"
import "path/filepath"

func TestMmap(t *testing.T) {
  s := NewMmapSeries()
  s.Load(filepath.Join("./testdata/test.csv"))

  reader := s.ReaderAt()
  t.Log(reader.Len())
  t.Log(s.RowSize())
  lr := s.LastRow()

  data := make([]byte, 10)
  _, err := lr.Read(data)
  if err != nil {
    panic(err)
  }
  for _, v := range data {
    t.Log(v)
  }

  t.Logf("%d - %d\n", s.TsFrom(), s.TsTo())
}