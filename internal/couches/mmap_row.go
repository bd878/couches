package couches

import "encoding/binary"
import "golang.org/x/exp/mmap"

type MmapRow struct {
  reader *mmap.ReaderAt
  off int
}

func NewMmapRow(r *mmap.ReaderAt, off int) *MmapRow {
  return &MmapRow{
    reader: r,
    off: off,
  }
}

func (r *MmapRow) Ts() int64 {
  b := make([]byte, 0, 8) // int64

  n, err := r.reader.ReadAt(b, int64(r.off))
  if err != nil {
    panic(err)
  }

  if n < len(b) {
    panic("read less bytes")
  }

  v, n := binary.Varint(b)
  if n != len(b) {
    panic("Varint did not consume all of input")
  }

  return v
}