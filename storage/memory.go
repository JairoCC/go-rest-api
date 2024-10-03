package storage

import (
	"fmt"

	"github.com/JairoCC/go-rest-api/model"
)

type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

func NewMemory() Memory {
	persons := make(map[int]model.Person)
	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCannotBeNil
	}
	m.currentID++
	m.Persons[m.currentID] = *person
	return nil
}

func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCannotBeNil
	}
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}
	m.Persons[ID] = *person
	return nil
}

func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}
	delete(m.Persons, ID)
	return nil
}

func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}
	return person, nil
}

func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}
	return result, nil
}
