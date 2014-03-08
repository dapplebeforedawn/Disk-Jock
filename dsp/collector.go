package dsp

type Collector struct {
  Size    int
  count   int
  Samples []int32
  Done    chan []int32
}

func NewCollector(size int, done chan []int32) *Collector {
  return &Collector{
    Size: size,
    Done: done,
  }
}

func (c *Collector) Add(data []int32) bool {
  c.count++
  c.Samples = append(c.Samples, data...)

  if c.count == c.Size {
    c.Done <- c.Samples
    return false
  }

  return true
}
