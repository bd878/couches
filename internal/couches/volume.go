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

func (v *Volume) Path() string {
  return v.path
}

func (v *Volume) Read() {
  entries, err := os.ReadDir(v.path)
  if err != nil {
    panic(err)
  }
  v.entries = entries
}

func (v *Volume) Scan() chan *Series {
  ch := make(chan *Series)

  go func() {
    defer close(ch)

    for _, entry := range v.entries {
      series := NewSeries()

      series.Load(filepath.Join(v.path, entry.Name()))

      ch <- series
    }
  }()

  return ch
}

func (v *Volume) Write(s *Series) {
  fd, err := os.OpenFile(filepath.Join(v.Path(), s.Path()), os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    panic(err)
  }

  s.Save(fd)
}