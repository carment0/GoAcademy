// Suppose we wanted the sum of 1 to 8000000000
// Besides iterating through each number, what is a faster way of tackling this problem? By spilting the work load!
// We will create three threads: main fn, worker1 and worker2
// worker1 will calculate from 1 to 4000000000, worker2 will calculate from 4000000001 to 8000000000

package main

import (
  "fmt"
  "time"
)

func SumRange(start int, end int) int {
  sum := 0
  for i := start; i <= end; i += 1 {
    sum += i
  }

  return sum
}

func Worker(output chan int, start int, end int) {
  sum := SumRange(start, end)
  output <- sum
}


func main(){

  workerOutput := []int{}

  outputChan := make(chan int)

  start := time.Now()
  go Worker(outputChan, 1, 4000000000)
  go Worker(outputChan, 4000000001, 8000000000)

  for len(workerOutput) != 2 {
    partialSum := <- outputChan
    workerOutput = append(workerOutput, partialSum)
  }

  t := time.Now()
  elapsed := t.Sub(start)

  fmt.Printf("The sum is %v\n", workerOutput[0] + workerOutput[1])
  fmt.Printf("The time it took to compute: %v\n", elapsed)


  start = time.Now()
  sum := SumRange(1, 8000000000)
  t = time.Now()
  elapsed = t.Sub(start)

  fmt.Printf("The sum is %v\n", sum)
  fmt.Printf("The time it took to compute: %v\n", elapsed)
}
