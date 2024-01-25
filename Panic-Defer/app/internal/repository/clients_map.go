package repository

import (
	"errors"
	"fmt"

	"github.com/jcanonbenavi/app/internal"
)

var (
	ErrConstraintViolation = errors.New("Error: client already exists")
	ErrInvalidClient       = errors.New("Error: invalid client")
)

type clientMap struct {
	client map[int]internal.Client
	lastId int // auto-increment
}

// NewClientMap returns a new clientMap
func NewClientMap() *clientMap {
	return &clientMap{
		client: make(map[int]internal.Client),
	}
}

// GetAll returns all clients
func (m *clientMap) GetAll() (clients map[int]internal.Client) {
	clients = m.client
	return
}

func (m *clientMap) Save(cliente *internal.Client) (err error) {
	// verify constraints
	err = verifyConstraints(m.client, cliente)
	if err != nil {
		return
	}
	// business logic
	err = bussinessLogic(cliente)
	if err != nil {
		return
	}
	// save
	// - increment id
	(*m).lastId++
	// - set id
	(*cliente).ID = (*m).lastId
	// - save
	m.client[cliente.ID] = *cliente
	return
}

// verifyConstraints verifies that the client does not exist
func verifyConstraints(m map[int]internal.Client, cliente *internal.Client) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Several errors were detected at runtime: %v", r)
		}
	}()
	for _, m := range m {
		if m.ID == cliente.ID {
			panic(ErrConstraintViolation)
		}
		if m.Name == cliente.Name {
			panic(ErrConstraintViolation)
		}
	}
	return
}

// bussinessLogic validates the business logic of the client
func bussinessLogic(cliente *internal.Client) (err error) {
	if cliente.Name == "" {
		err = fmt.Errorf("%w. field: name", ErrInvalidClient)
		return
	}
	if cliente.File == "" {
		err = fmt.Errorf("%w. field: file", ErrInvalidClient)
		return
	}
	if cliente.Phone == 0 {
		err = fmt.Errorf("%w. field: phone", ErrInvalidClient)
		return
	}
	if cliente.Home == "" {
		err = fmt.Errorf("%w. field: home", ErrInvalidClient)
		return
	}
	return
}
