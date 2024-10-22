package storage

import (
	"testing"

	"github.com/JairoCC/go-rest-api/model"
)

func TestCreate(t *testing.T) {
	table := []struct {
		name        string
		person      *model.Person
		wantedError error
	}{
		{"empty_person", nil, model.ErrPersonCannotBeNil},
		{"Jairo", &model.Person{Name: "Jairo"}, nil},
		{"Matthew", &model.Person{Name: "Matthew"}, nil},
	}

	m := NewMemory()

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			gotErr := m.Create(v.person)
			if gotErr != v.wantedError {
				t.Errorf("%v was expected nut %v was returned", v.wantedError, gotErr)
			}
		})
	}
	wantedCount := len(table) - 1
	if m.currentID != wantedCount {
		t.Errorf("%d was expected, but %d was returned", wantedCount, m.currentID)
	}
}
