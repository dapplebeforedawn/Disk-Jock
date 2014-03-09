package main

import (
  "dapplebeforedawn/disk-jock/matchtable"
  "dapplebeforedawn/disk-jock/loop"
  "fmt"
  "os"
)

func main() {
  hashFile, _ := os.Open("match-hashes.dat")
  matchTable := matchtable.NewMatchTable(hashFile)

  callback := func(in, out []int32) {
    copy(out, in)
  }

  loopback := func(mags []float64) {
    wasMatched  := matchTable.HasMatch(mags)
    fmt.Println(wasMatched)
  }

  l := loop.NewLoop(callback, loopback)
  l.Start()
}
