package couches

import "strconv"

type Row []string

func (r Row) Ts() int64 {
  res, err := strconv.ParseInt(r[0], 10, 64)
  if err != nil {
    panic(err)
  }
  return res
}