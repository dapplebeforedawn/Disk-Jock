package dsp

import (
  "github.com/mjibson/go-dsp/fft"
  "math"
)

type Dsp struct {
}

func NewFFT() *Dsp {
  return &Dsp{}
}

func (d* Dsp) FFT(data []int32) []float64 {
  mags := make([]float64, len(data))
  for i, comp := range d.fft(data) {
    rel := math.Pow(real(comp), 2)
    img := math.Pow(imag(comp), 2)
    mags[i] = math.Sqrt(rel + img)
  }
  return mags
}

func (d *Dsp) fft(data []int32) []complex128 {
  data64 := make([]float64, len(data))
  for i := range data {
    data64[i] = float64(data[i])
  }
  return fft.FFTReal(data64)
}
