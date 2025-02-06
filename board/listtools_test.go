package board

import "fmt"

func ExampleContainsOnly() {
	fmt.Println(ContainsOnly([]string{"O", "O", "O"}, "O"))
	fmt.Println(ContainsOnly([]string{"O", "O", "O"}, "X"))
	fmt.Println(ContainsOnly([]string{"O", "X", "O"}, "O"))
	fmt.Println(ContainsOnly([]string{"O", "X", "O"}, "X"))
	fmt.Println(ContainsOnly([]string{}, "X"))
	fmt.Println(ContainsOnly([]string{}, "O"))

	// Output:
	// true
	// false
	// false
	// false
	// true
	// true
}

func ExampleContainsAny() {
	fmt.Println(ContainsAny([]string{"O", "X", "O"}, "O"))
	fmt.Println(ContainsAny([]string{"O", "X", "O"}, "X"))
	fmt.Println(ContainsAny([]string{"O", "X", "O"}, " "))
	fmt.Println(ContainsAny([]string{}, "X"))
	fmt.Println(ContainsAny([]string{}, "O"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}
