package main

import (
  "dapplebeforedawn/disk-jock/system"
  "dapplebeforedawn/disk-jock/dsp"
  "dapplebeforedawn/disk-jock/matchtable"
  "fmt"
  // "os"
)

func main() {
  matchTable := matchtable.NewMatchTable()
  c := make(chan []int32, 100000)
  callback := func(in, out []int32) {
    c <- in
    copy(out, in)
  }

  go func(){
    // file, _ := os.Create("fft.dat")
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
      // fmt.Fprintln(file, mags)

      wasMatched := matchTable.HasMatch(mags[10:len(mags)/2])
      fmt.Println(wasMatched)
    }
  }()

  sys := system.NewSystem(callback)
  sys.Start()
}
