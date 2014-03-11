package peak

import (
  "math"
)

func FindPeaks(samples []float64) (peaks []int) {
  min    := math.Inf(1)
  max    := math.Inf(-1)
  thresh := 2000000000.0

  lookForMax := true

  for i, val := range samples {
    if val > max {
      max = val
    }

    if val < min {
      min = val
    }

    if lookForMax {
      if val < max - thresh {
        peaks  = append(peaks, i)
        min    = val
        lookForMax = false
      }
    } else {
      if val > min + thresh {
        max        = val
        lookForMax = true
      }
    }
  }

  return
}
