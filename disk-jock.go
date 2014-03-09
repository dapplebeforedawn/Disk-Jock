package main

import (
  "dapplebeforedawn/disk-jock/matchtable"
  "dapplebeforedawn/disk-jock/loop"
  "fmt"
)

func main() {
  matchTable := matchtable.NewMatchTable()

  callback := func(in, out []int32) {
    copy(out, in)
  }

  loopback := func(mags []float64) {
    wasMatched := matchTable.HasMatch(mags[10:len(mags)/2])
    fmt.Println(wasMatched)
  }

  l := loop.NewLoop(callback, loopback)
  l.Start()
}
