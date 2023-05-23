package couches

type LogSeries interface {
  TsFrom() int64
  TsTo() int64
}

type Slicer interface {
  Slice(fromTs int64, toTs int64) (interface{}, error)
}