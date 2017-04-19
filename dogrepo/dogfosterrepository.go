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
