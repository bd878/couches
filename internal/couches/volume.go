package couches

import "os"
import "path/filepath"

type Volume struct {
  path string
  entries []os.DirEntry
}

func NewVolume(path string) *Volume {
  return &Volume{
    path: path,
  }
}

func (v *Volume) Read() error {
  entries, err := os.ReadDir(v.path)
  if err != nil {
    return err
  }
  v.entries = entries
  return nil
}

func (v *Volume) Scan() chan *Series {
  ch := make(chan *Series)

  go func() {
    defer close(ch)

    for _, entry := range v.entries {
      series := NewSeries()

      err := series.Load(filepath.Join(v.path, entry.Name()))
      if err != nil {
        panic(err)
      }

      ch <- series
    }
  }()

  return ch
}
