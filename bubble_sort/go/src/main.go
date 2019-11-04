package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	timestart := time.Now()

	array := []int16{3, 4, 1, 3, 5, 1, 92, 2, 4124, 424, 52, 12}

	for c := 0; c < 1000000; c++ {

		for i := 0; i < len(array); i++ {
			for y := 0; y < len(array)-1; y++ {
				if array[y+1] < array[y] {
					t := array[y]
					array[y] = array[y+1]
					array[y+1] = t
				}
			}
		}

	}

	fmt.Print(array)
	fmt.Print("\n")

	timeend := time.Now()

	fmt.Println(timeend.Sub(timestart))
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v KB", bToKb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v KB", bToKb(m.TotalAlloc))
}

func bToKb(b uint64) uint64 {
	return b / 1024
}
