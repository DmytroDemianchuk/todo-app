// sobirae vsi service vmeste
package service

import "github.com/dmytrodemianchuk/todo-app/pkg/repository"

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
