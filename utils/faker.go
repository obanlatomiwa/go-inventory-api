package utils

import "github.com/bxcodec/faker/v3"

// CreateFaker returns faker data with the type of T
func CreateFaker[T any]() (T, error) {
	var fakerData *T = new(T)

	err := faker.FakeData(fakerData)
	if err != nil {
		return *fakerData, err
	}
	return *fakerData, nil
}
