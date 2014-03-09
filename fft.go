package main

import (
  "github.com/mjibson/go-dsp/fft"
  "github.com/mjibson/go-dsp/window"
  "math"
  "os"
  "bufio"
  "strconv"
  "fmt"
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil { panic(err) }

  scanner   := bufio.NewScanner(file)
  values    := make([]float64, 0)
  for scanner.Scan() {
    line     := scanner.Text()
    value, e := strconv.ParseFloat(line, 64)
    if e != nil { panic(e) }
    values = append(values, value)
  }

  fftComplex := doFft(values)
  fftReal    := realize(fftComplex)

  for _, freq := range fftReal {
    fmt.Println(freq)
  }
}

func doFft(data []float64) []complex128 {
  window.Apply(data, window.Hamming)
  return fft.FFTReal(data)
}

func realize(data []complex128) []float64 {
  realVals := make([]float64, len(data))
  for i, comp := range data {
    rel := math.Pow(real(comp), 2)
    img := math.Pow(imag(comp), 2)
    realVals[i] = math.Sqrt(rel + img)
  }
  return realVals
}
