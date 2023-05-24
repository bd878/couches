package main

import "flag"
import "path/filepath"
import "os"
import "strconv"
import "fmt"
import "github.com/bd878/couches/internal/couches"

var (
  outf = flag.String("in", "./testdata/records", "volumes dir")
  inf = flag.String("out", "./testdata/result", "volumes result dir")
  from = flag.String("from", "", "from ts")
  to = flag.String("to", "", "to ts")
)

func main() {
  flag.Parse()

  if *from == "" {
    fmt.Println("no \"from\" flag provided\n")
    flag.Usage()
    os.Exit(1)
  }

  if *to == "" {
    fmt.Println("no \"out\" flag provided\n")
    flag.Usage()
    os.Exit(1)
  }

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

  if entries, err = os.ReadDir(*outf); err != nil {
    panic(err)
  }

  err = os.MkdirAll(*inf, 0750)
  if err != nil && !os.IsExist(err) {
    panic(err)
  }

  rvol := couches.NewVolume(*inf)

  for _, entry := range entries {
    vol := couches.NewVolume(filepath.Join(*outf, entry.Name()))

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
