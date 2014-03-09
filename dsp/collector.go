package dsp

type Collector struct {
  Size    int
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
  c.Samples = append(c.Samples, data...)

  if len(c.Samples) >= c.Size {
    c.Samples = c.Samples[:c.Size]
    c.Done <- c.Samples
    return false
  }

  return true
}
