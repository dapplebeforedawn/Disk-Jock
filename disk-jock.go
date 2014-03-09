package main

import (
  "dapplebeforedawn/disk-jock/matchtable"
  "dapplebeforedawn/disk-jock/loop"
  "dapplebeforedawn/disk-jock/options"
  "fmt"
  "os"
)

func main() {
  opts := opts.Options{}
  opts.Parse()

  hashFile, _ := os.Open(opts.HashFile)
  matchTable  := matchtable.NewMatchTable(hashFile)

  callback := func(in, out []int32) {
    copy(out, in)
  }

  waveFile, _ := os.OpenFile(opts.WavOutput, os.O_RDWR|os.O_CREATE, 0660);
  defer waveFile.Close()
  collectionFilled := func(samples []int32) {
    fmt.Fprintln(waveFile, samples)
  }

  fftFile, _ := os.OpenFile(opts.FftOutput, os.O_RDWR|os.O_CREATE, 0660);
  defer fftFile.Close()
  loopback := func(mags []float64) {
    fmt.Fprintln(fftFile, mags)

    // wasMatched  := matchTable.HasMatch(mags)
    _  = matchTable.HasMatch(mags)
  }

  l := loop.NewLoop(callback, loopback)
  l.Filledback = collectionFilled
  l.Start()
}
