package opts

import (
  "flag"
  "fmt"
  "os"
)

type Options struct {
  FftOutput string
  WavOutput string
  HashFile  string
}

func (o *Options) Parse() {
  flag.Usage  = usage
  fftOutput  := ""
  waveOut    := ""
  hashFile   := "hashes.dat"

  flag.StringVar(&fftOutput, "fft", fftOutput, "output file for the FFT")
  flag.StringVar(&waveOut, "wav", waveOut, "output file for the raw wave")
  flag.StringVar(&hashFile, "hash", hashFile, "the file containing the audio fingerprints")
  flag.Parse()

  o.FftOutput = fftOutput
  o.WavOutput = waveOut
  o.HashFile  = hashFile
}

func usage() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "Example: ")
  fmt.Fprintln(os.Stderr, "  disk-jock -fft=fft-output.dat")
}
