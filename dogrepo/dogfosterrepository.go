package dogrepo

// Dog represents a dog in foster
type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// DogFosterRepository contains functions to query a repository of dogs
type DogFosterRepository struct{}

// ExampleDog returns a sample Dog
func (dfr *DogFosterRepository) ExampleDog() *Dog {
	return &Dog{
		Name: "Joey",
		Age:  13,
	}
}

// ExampleDogs returns a slice of sample Dogs
func (dfr *DogFosterRepository) ExampleDogs() *[]Dog {
	return &[]Dog{
		Dog{
			Name: "Joey",
			Age:  13,
		},
		Dog{
			Name: "Woody",
			Age:  1,
		},
	}
}
