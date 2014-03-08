package main

import (
  "dapplebeforedawn/disk-jock/system"
  "dapplebeforedawn/disk-jock/dsp"
  "fmt"
  "os"
)
// need 5 chunks to make a decent FFT

func main() {
  c := make(chan []int32, 100000)
  callback := func(in, out []int32) {
    c <- in
    copy(out, in)
  }

  go func(){
    file, _ := os.Create("fft.dat")
    for {
      fft     := dsp.NewFFT()
      done    := make(chan []int32)
      collect := dsp.NewCollector(8, done)

      go func(){
        for {
          data := <-c
          if !collect.Add(data) { break }
        }
      }()

      fullData := <-done
      mags     := fft.FFT(fullData)
      fmt.Fprintln(file, mags[10:len(mags)/2])
    }
  }()

  sys := system.NewSystem(callback)
  sys.Start()
}
