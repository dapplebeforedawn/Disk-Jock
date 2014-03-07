
package main

import (
  "code.google.com/p/portaudio-go/portaudio"
)

func main() {
  portaudio.Initialize()
  defer portaudio.Terminate()
  h, err := portaudio.DefaultHostApi()
  chk(err)
  for _, d := range h.Devices {
    println(d.Name)
  }

  callback := func(in, out []int32) {
      for i := range in {
        out[i] = in[i]
      }

  }

  inParam := portaudio.StreamDeviceParameters{
    Device: h.Devices[6],
    Channels: 2,
  }

  outParam := portaudio.StreamDeviceParameters{
    Device: h.Devices[1],
    Channels: 2,
  }

  streamParams := portaudio.StreamParameters{
    Input: inParam,
    Output: outParam,
    SampleRate: h.Devices[1].DefaultSampleRate,
  }

  stream, err := portaudio.OpenStream(streamParams, callback)
  // stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(h.Devices[6], h.Devices[1]), callback)
  chk(err)

  defer stream.Close()
  chk(stream.Start())
  select {}
  chk(stream.Stop())
}

func chk(err error) {
  if err != nil {
    panic(err)
  }
}
