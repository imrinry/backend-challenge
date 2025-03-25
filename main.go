package main

import "fmt"

func main() {
	// โจทย์ข้อ 1
	fmt.Println("Problem 1: Maximum Path Sum")
	test1 := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	ans1 := findSum(test1)
	fmt.Println("Example result:", ans1)

	// โจทย์ข้อ 2
	fmt.Println("\nProblem 2: Decode LR= Pattern")
	fmt.Print("Enter encoded string (L,R,=): ")
	var input string
	fmt.Scanln(&input)
	result := decode(input)
	fmt.Println("Result:", result)

	// โจทย์ข้อ 3
	fmt.Println("\nProblem 3: Meat Counting API")
	fmt.Println("Starting API server on :8080")
	startMeatAPI()
}
