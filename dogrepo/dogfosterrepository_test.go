package dogrepo

import "fmt"

func ExampleDogFosterRepository_ExampleDog() {
	repo := DogFosterRepository{}
	fmt.Println(repo.ExampleDog())
	// Output:
	// &{Joey 13}
}

func ExampleDogFosterRepository_ExampleDogs() {
	repo := DogFosterRepository{}
	fmt.Println(repo.ExampleDogs())
	// Output:
	// &[{Joey 13} {Woody 1}]
}
