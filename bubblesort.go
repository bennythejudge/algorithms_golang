package main

import "fmt"

func bubble_sort(x []int) ([]int, int, int) {
  swapped := 1
  loop_count := 0
  swap_counter := 0

  for {
    if swapped == 0 {
      break
    }
    swapped = 0
    for i := 0; i < len(x)-1; i++ {
      if x[i] > x[i+1] {
        x[i], x[i+1] = x[i+1], x[i]
        swapped = 1
        swap_counter++
      }
    }
    loop_count++
    fmt.Println(x)
  }
  return x, loop_count, swap_counter
}

func main() {
  x := []int{
    48, 96, 86, 68,
    57, 82, 63, 70,
    37, 34, 83, 27,
    19, 97, 9, 17,
  }

  var sorted_array []int
  var loop_count int
  var swap_counter int

  worst_case := []int{97, 96, 86, 83, 82, 70, 68, 63, 57, 48, 37, 34, 27, 19, 17, 9}

  fmt.Println("Bubble sort examples in Go\n\n")

  fmt.Println("unsorted array ")
  fmt.Println(x)
  fmt.Println()

  sorted_array, loop_count, swap_counter = bubble_sort(x)

  fmt.Println("sorted in ", loop_count, " loops and ", swap_counter, " swaps")
  fmt.Println(sorted_array)
  fmt.Println()
  fmt.Println()

  fmt.Println("reverse sorted array - worst case")
  fmt.Println(worst_case)

  sorted_array, loop_count, swap_counter = bubble_sort(worst_case)

  fmt.Println("sorted in ", loop_count, " loops and ", swap_counter, " swaps")
  fmt.Println(sorted_array)

}
