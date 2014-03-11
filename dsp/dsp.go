package dsp

import (
  "github.com/mjibson/go-dsp/fft"
  // "github.com/mjibson/go-dsp/window"
  // "github.com/runningwild/go-fftw"
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
  // window.Apply(data64, window.Hamming)
  return fft.FFTReal(data64)
}

// func (d *Dsp) fft(data []int32) []complex128 {
//   ffData   := fftw.Alloc1d(1024)  // Similar to calling make([]complex128, 64)
//   forward  := fftw.PlanDft1d(ffData, ffData, fftw.Forward, fftw.Estimate)
//
//   for i := range ffData {
//     ffData[i] = complex(float64(data[i]), float64(0))
//   }
//   forward.Execute()
//   return ffData
// }
