package dogrepo

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DogFosterRepository struct{}

func (dfr *DogFosterRepository) ExampleDog() *Dog {
	return &Dog{
		Name: "Joey",
		Age:  13,
	}
}

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
