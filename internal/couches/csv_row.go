package couches

import "strconv"

type CSVRow []string

func NewCSVRow(str []string) CSVRow {
  var res CSVRow = str
  return res
}

func (r CSVRow) Ts() int64 {
  res, err := strconv.ParseInt(r[0], 10, 64)
  if err != nil {
    panic(err)
  }
  return res
}