package tasks

import "fmt"

func Task15() {
	someFunc()
}

func someFunc() {
	arr := createHugeString()
	fmt.Printf("Task15: %c\n", arr)
}

func createHugeString() []byte {
	hugeString := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	byteString := []byte(hugeString)
	arr := make([]byte, len(byteString[:100]))
	copy(arr, byteString[:100])
	return arr
}
