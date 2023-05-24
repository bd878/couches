package couches

import "os"
import "encoding/csv"
import "golang.org/x/exp/mmap"

var newLine byte = byte('\n')

type MmapSeries struct {
  path string
  heading []string
  headingOffset int
  rowSize int
  reader *mmap.ReaderAt
}

func NewMmapSeries() *MmapSeries {
  return &MmapSeries{}
}

func (s *MmapSeries) Path() string {
  return s.path
}

func (s *MmapSeries) RowSize() int64 {
  return int64(s.rowSize)
}

func (s *MmapSeries) HeadingOffset() int64 {
  return int64(s.headingOffset)
}

func (s *MmapSeries) Close() {
  s.reader.Close()
}

func (s *MmapSeries) ReaderAt() *mmap.ReaderAt {
  return s.reader
}

func (s *MmapSeries) TsFrom() int64 {
  return s.FirstRow().Ts()
}

func (s *MmapSeries) TsTo() int64 {
  return s.LastRow().Ts()
}

func (s *MmapSeries) FirstRow() *MmapRow {
  return NewMmapRow(s.reader, s.headingOffset)
}

func (s *MmapSeries) LastRow() *MmapRow {
  return NewMmapRow(s.reader, s.reader.Len() - s.rowSize)
}

func (s *MmapSeries) Load(path string) {
  s.path = path

  fd, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer fd.Close()

  cr := csv.NewReader(fd)
  heading, err := cr.Read()
  if err != nil {
    panic(err)
  }
  s.heading = heading

  r, err := mmap.Open(path)
  if err != nil {
    panic(err)
  }

  for 
    s.headingOffset = 0;
    r.At(s.headingOffset) != newLine;
    s.headingOffset++ {
  }
  s.headingOffset += 1 // newLine

  for
    s.rowSize = 0;
    r.At(s.headingOffset + s.rowSize) != newLine;
    s.rowSize++ {
  }
  s.rowSize += 1 // newLine

  s.reader = r
}
