package couches

import "os"

type LogSeries interface {
  TsFrom() int64
  TsTo() int64

  Path() string
  Save(fd *os.File)
  Load(path string)
  Close()
  FirstRow() *LogRow
  LastRow() *LogRow
}

type Slicer interface {
  Slice(path string, fromTs int64, toTs int64) interface{}
}

type LogRow interface {
  Ts() int64
}