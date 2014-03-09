package loop

import (
  "dapplebeforedawn/disk-jock/system"
  "dapplebeforedawn/disk-jock/dsp"
)

const SAMPLE_SIZE int = 1024
const BUFFER_SIZE int = 100000

type Loop struct {
  callback   system.Callback
  channel    chan []int32
  loopback   Loopback
  Filledback Filledback
}

type Loopback   func([]float64)
type Filledback func([]int32)

func NewLoop(cb system.Callback, lb Loopback) Loop {
  l := Loop{
    channel:  make(chan []int32, BUFFER_SIZE),
    loopback: lb,
  }
  l.callback = l.decorateCallback(cb)

  return l
}

func (l *Loop) decorateCallback(cb system.Callback) system.Callback {
  return func(in, out []int32) {
    l.channel <- in
    cb(in, out)
  }
}

func (l *Loop) runLoop(){
  for {
    fft     := dsp.NewFFT()
    done    := make(chan []int32)

    go l.collectSamples(done)

    fullData    := <-done
    mags        := fft.FFT(fullData)
    half        := len(mags)/2
    useableMags := mags[:half]

    l.loopback(useableMags)
  }
}

func (l *Loop) collectSamples(done chan[]int32) {
  collect := dsp.NewCollector(SAMPLE_SIZE, done)
  for {
    data := <-l.channel
    if !collect.Add(data) { break }
  }
  l.Filledback(collect.Samples)
}

func (l *Loop) Start() {
  go l.runLoop()

  sys := system.NewSystem(l.callback)
  sys.Start()
}
