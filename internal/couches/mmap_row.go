package couches

import "strings"
import "strconv"
import "golang.org/x/exp/mmap"

var sep byte = byte(',')

type MmapRow struct {
  reader *mmap.ReaderAt
  off int
  ts int64
}

func NewMmapRow(r *mmap.ReaderAt, off int) *MmapRow {
  return &MmapRow{
    reader: r,
    off: off,
  }
}

func (r *MmapRow) Read(p []byte) (int, error) {
  return r.reader.ReadAt(p, int64(r.off))
}

func (r *MmapRow) Ts() int64 {
  if r.ts != 0 {
    return r.ts
  }

  var builder strings.Builder
  for i := 0; r.reader.At(r.off + i) != sep; i++ {
    err := builder.WriteByte(r.reader.At(r.off + i))
    if err != nil {
      panic(err)
    }
  }

  var err error
  r.ts, err = strconv.ParseInt(builder.String(), 10, 64)
  if err != nil {
    panic(err)
  }

  return r.ts
}