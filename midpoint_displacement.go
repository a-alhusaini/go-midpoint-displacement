package main

import (
  "fmt"
  "math/rand"
)

func main() {
  hill := gen_hill(make([]int, 17))
  head := get_head(hill)
  head = populate_terrain(head)

  tail := get_tail(hill)
  tail = populate_terrain(tail)

  fmt.Println(hill)
}

func gen_hill(hill []int) []int {
  hill[0] = rand.Intn(99) + 1
  hill[len(hill) - 1] = rand.Intn(99) + 1
  hill = set_midpoint(hill)

  return hill
}

func get_tail(slice []int) []int {
  return slice[len(slice) / 2:]
}

func get_head(slice []int) []int {
  return slice[:(len(slice) + 1) / 2]
}

func set_midpoint(slice []int) []int {
  slice[len(slice) / 2] = (slice[0] + slice[len(slice) - 1]) / 2
  return slice
}

func populate_terrain(hill []int) []int {
  lower := hill[0]
  higher := hill[len(hill) - 1]
  if (lower > higher) {
    higher = hill[0]
    lower = hill[len(hill) - 1]
  }

  for i := range(hill) {
    if i == 0 || i == len(hill) - 1 {
      continue
    }

    hill[i] = rand.Intn(higher - lower + 1) + lower
  }

  return hill
}
