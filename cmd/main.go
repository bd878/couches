package main

import "flag"
import "fmt"
import "github.com/bd878/couches/internal/couches"

var p = flag.String("vol", "./couches", "path to volume dir")

func main() {
  flag.Parse()

  vol := couches.NewVolume(*p)

  err := vol.Read()
  if err != nil {
    panic(err)
  }

  for series := range vol.Scan() {
    fmt.Println(series.Count())
  }

  fmt.Println("ok")
}