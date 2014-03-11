package main

import (
  "dapplebeforedawn/disk-jock/peak"
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

  peaks := peak.FindPeaks(values)

  for _, freq := range peaks {
    fmt.Println(freq)
  }
}
