package main

import (
  "dapplebeforedawn/disk-jock/matchtable"
  "dapplebeforedawn/disk-jock/loop"
  "fmt"
  "os"
)

func main() {
  callback := func(in, out []int32) {
    copy(out, in)
  }

  loopback := func(mags []float64){}

  l := loop.NewLoop(callback, c, loopback)
  l.Start()
}
