package main

import "flag"
import "path/filepath"
import "os"
import "strconv"
import "fmt"
import "github.com/bd878/couches/internal/couches"

var (
  d = flag.String("dir", "./testdata/records", "volumes dir")
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

  var totalCount int

  for _, entry := range entries {
    vol := couches.NewVolume(filepath.Join(*d, entry.Name()))

    err := vol.Read()
    if err != nil {
      panic(err)
    }

    for series := range vol.Scan() {
      v, err := series.Slice(fromTs, toTs)
      if err != nil {
        panic(err)
      }

      res, ok := v.(*couches.Series)
      if !ok {
        panic("not a series")
      }

      totalCount += res.Count()
    }
  }

  fmt.Println(totalCount)
  fmt.Println("ok")
}