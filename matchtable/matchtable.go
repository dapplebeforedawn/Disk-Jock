package matchtable

import (
  "os"
  "bufio"
  "fmt"
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

  m := fmt.Sprintf("%06d", matchSet.Finger())
  _, found = mt.Table[m]
  if found { return true }

  return false
}
