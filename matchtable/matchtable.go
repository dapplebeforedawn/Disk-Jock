package matchtable

import (
  "os"
  "bufio"
)

type MatchTable struct {
  Table map[string]string
}

func NewMatchTable(file *os.File) *MatchTable {
  table   := make(map[string]string)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    table[scanner.Text()] = "exists"
  }

  return &MatchTable{ Table: table }
}

func (mt *MatchTable) HasMatch(data []float64) (found bool) {
  matchMaker := NewMatchMaker(data)
  matchSet   := matchMaker.Extract()

  _, found = mt.Table[matchSet.String()]
  return
}
