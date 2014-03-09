package main

import (
  "dapplebeforedawn/disk-jock/matchtable"
  "dapplebeforedawn/disk-jock/loop"
  "fmt"
  "os"
)

func main() {
  file, _ := os.Create("match-hashes.dat")

  callback := func(in, out []int32) {
    copy(out, in)
  }

  loopback := func(mags []float64){
    matchSet := matchtable.NewMatchMaker(mags).Extract()
    fmt.Fprintln(file, matchSet.String())
  }

  l := loop.NewLoop(callback, loopback)
  l.Start()
}
