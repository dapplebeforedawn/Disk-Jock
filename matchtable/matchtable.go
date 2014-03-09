package matchtable

import (
)

type MatchTable struct {
  Table map[string]string
}

func NewMatchTable() *MatchTable {
  return &MatchTable{ Table: make(map[string]string) }
}

func (mt *MatchTable) HasMatch(data []float64) (found bool) {
  matchMaker := NewMatchMaker(data)
  matchSet   := matchMaker.Extract()

  _, found = mt.Table[matchSet.String()]
  return
}
