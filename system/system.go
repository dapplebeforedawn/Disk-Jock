package system

import (
  "code.google.com/p/portaudio-go/portaudio"
)

type Callback func(in, out []int32)

type System struct {
  Api           *portaudio.HostApiInfo
  StreamParams  portaudio.StreamParameters
  InParams      portaudio.StreamDeviceParameters
  OutParams     portaudio.StreamDeviceParameters
  Callback      Callback
}

func NewSystem(callback Callback) *System {
  s := &System{
    Callback: callback,
  }

  portaudio.Initialize()

  api, err := portaudio.DefaultHostApi()
  chk(err)

  s.Api = api

  s.deviceNames()
  s.initParams()

  return s
}

func (s *System) Start() {
  stream, err := portaudio.OpenStream(s.StreamParams, s.Callback)
  chk(err)

  defer stream.Close()
  chk(stream.Start())
  select {}
  chk(stream.Stop())
  portaudio.Terminate()
}

func (s *System) initParams() {
  inDev := s.Api.Devices[4] //6
  s.InParams = portaudio.StreamDeviceParameters{
    Device: inDev,
    Channels: 2,
    Latency: inDev.DefaultLowOutputLatency,
  }

  outDev := s.Api.Devices[0] //1
  s.OutParams = portaudio.StreamDeviceParameters{
    Device: outDev,
    Channels: 2,
    Latency: outDev.DefaultLowOutputLatency,
  }

  s.StreamParams = portaudio.StreamParameters{
    Input: s.InParams,
    Output: s.OutParams,
    SampleRate: s.Api.Devices[4].DefaultSampleRate,
  }
}

func (s *System) deviceNames() {
  for _, device := range s.Api.Devices {
    println(device.Name)
  }
}

func chk(err error) {
  if err != nil { panic(err)
  }
}
