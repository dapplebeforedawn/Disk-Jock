package matchtable

import (
  "fmt"
)

type MatchMaker struct {
  Data []float64
}

type MatchSet [8]MatchPair
type MatchPair struct {
  Magnitude float64
  Frequency int
}

func NewMatchMaker(data []float64) *MatchMaker {
  return &MatchMaker{ Data: data }
}

func (m *MatchMaker) Extract() MatchSet {
  ms := MatchSet{}

  // for each in data, if its bigger than anyone
  // in matchSet then, replace the match with the
  // datum.  Give the datum a chance to get back
  // into the matchSet.

  for i, data := range m.Data {
    ejected := ms.TryAdd(i, data)
    if(ejected != MatchPair{}) {
      ms.TryAdd(ejected.Frequency, ejected.Magnitude)
    }
  }

  return ms
}

func (ms *MatchSet) TryAdd(freq int, mag float64) MatchPair {
  ejected := MatchPair{}
  for i, matchPair := range ms {
    if mag > matchPair.Magnitude {
      ejected = matchPair
      ms[i] = MatchPair{
        Magnitude: mag,
        Frequency: freq,
      }
      break
    }
  }
  return ejected
}

func (ms *MatchSet) String() (s string) {
  for _, matchPair := range ms {
    s = s + fmt.Sprintf("%05d", matchPair.Frequency)
  }
  return
}
