// sobirae vsi service vmeste
package service

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// service obraschaetka k baze damym
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
