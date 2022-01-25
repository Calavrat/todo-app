package service

import (
	"github.com/Calavrat/todo-app/model"
	"github.com/Calavrat/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, pasword string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listid int) (model.TodoList, error)
	Delete(userid, listid int) error
	Update(userid, listid int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item model.TodoItem) (int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userid, itenId int, input model.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
