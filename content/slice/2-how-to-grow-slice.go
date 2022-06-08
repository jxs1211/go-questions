package main

import "fmt"

func main() {
	s := make([]int, 0)
	oldCap := cap(s)
	for i := 0; i < 2048; i++ {
		s = append(s, i)
		newCap := cap(s)
		if oldCap != newCap {
			fmt.Printf("[slice's cap: %4d], [after append %-4d], the new cap: %d, the rate of growth: %.2f\n", oldCap, i, newCap, float64(newCap)/float64(oldCap))
			oldCap = newCap
		}
	}

}

func f() {
}
