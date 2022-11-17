package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	GetAll() (map[int]Employee, error)
	Update(id int, e Employee)
	Delete(id int) error
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		counter: 1,
		data:    make(map[int]Employee),
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()
	e.ID = s.counter
	s.data[e.ID] = *e
	s.counter++
	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	e, exist := s.data[id]
	if !exist {
		return e, errors.New("сотрудник не найден")
	}

	return e, nil
}

func (s *MemoryStorage) GetAll() (map[int]Employee, error) {
	s.Lock()
	defer s.Unlock()
	return s.data, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
	return nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}
