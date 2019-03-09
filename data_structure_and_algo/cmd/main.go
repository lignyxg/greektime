package main

import (
	"fmt"
	"greektime/data_structure_and_algo/recursion"
)

func main() {
	// quick sort
	//a1 := []int{2,4,1,7,3,8,5}
	//sort.QuickSort(a1)
	//fmt.Println("quickSort:\n", a1)
	//
	////merge sort
	//a2 := []int{2,4,1,7,3,8,5}
	//a3 := sort.MergeSort(a2)
	//fmt.Println("mergeSort:\n", a3)
	//
	////merge sort2
	//a4 := []int{2,4,1,7,3,8,5}
	//sort.MergeSort2(a4)
	//fmt.Println("mergeSort2:\n", a4)

	//count sort
	//a5 := []int{2,4,1,7,3,8,5}
	//sort.CountSort(a5)
	//fmt.Println("countSort:\n", a5)

	//b := []int{4,5,6,7,0,1,2}
	////c := []int{3,1}
	//k := 0
	//i := biserach.FindInCyclingArray(b, len(b), k)
	//fmt.Printf("index of %d is %d", k, i)

	fibNum := 7
	fmt.Printf("Fibonacci(%d)=%d\n",fibNum,recursion.Fibonacci(fibNum))
	fmt.Printf("Fibonacci2(%d)=%d\n",fibNum,recursion.Fibonacci2(fibNum))

	factor := 4
	fmt.Printf("Factorial(%d)=%d\n",factor, recursion.Factorial(factor))

	a6 := []int{1,2,3}
	p := recursion.Permutation(a6)
	for _, x := range p {
		fmt.Println(x)
	}
}
