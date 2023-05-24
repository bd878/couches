package main

import "flag"
import "path/filepath"
import "os"
import "strconv"
import "fmt"
import "github.com/bd878/couches/internal/couches"

var (
  d = flag.String("in", "./testdata/records", "volumes dir")
  res = flag.String("out", "./testdata/result", "volumes result dir")
  from = flag.String("from", "", "from ts")
  to = flag.String("to", "", "to ts")
)

func main() {
  flag.Parse()

  var err error
  var fromTs int64
  var toTs int64
  var entries []os.DirEntry

  if fromTs, err = strconv.ParseInt(*from, 10, 64); err != nil {
    panic(err)
  }

  if toTs, err = strconv.ParseInt(*to, 10, 64); err != nil {
    panic(err)
  }

  if entries, err = os.ReadDir(*d); err != nil {
    panic(err)
  }

  err = os.MkdirAll(*res, 0750)
  if err != nil && !os.IsExist(err) {
    panic(err)
  }

  rvol := couches.NewVolume(*res)

  for _, entry := range entries {
    vol := couches.NewVolume(filepath.Join(*d, entry.Name()))

    vol.Read()

    for series := range vol.Scan() {
      v := series.Slice(
        fmt.Sprintf("%s_%d", entry.Name(), series.TsFrom()),
        fromTs, toTs,
      )

      s, ok := v.(*couches.CSVSeries)
      if !ok {
        panic("not a series")
      }

      if s.Count() > 0 {
        rvol.Write(s)
      }
    }
  }

  fmt.Println("ok")
}