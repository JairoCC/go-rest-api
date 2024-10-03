package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/JairoCC/go-rest-api/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "object structure is incorrect", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "an error occurred when creating person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "person created successfully", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "method not allowed"}`))
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "ID must be a positive integer"}`))
		return
	}
	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "object structure is incorrect"}`))
		return
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "an error occurred when updating person"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message_type": "message", "message": "person updated successfully"}`))

}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID must be a positive integer", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "person ID does not exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "an error occurred when deleting a person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "an error occurred when trying to get all persons", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
	// w.Write([]byte(`{"message_type": "message", "message": "Ok"}`))
}
