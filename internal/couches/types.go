package couches

type LogSeries interface {
  TsFrom() int64
  TsTo() int64
}