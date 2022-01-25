package service

import (
	"github.com/Calavrat/todo-app/model"
	"github.com/Calavrat/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listid int) (model.TodoList, error) {
	return s.repo.GetById(userId, listid)
}

func (s *TodoListService) Delete(userId, listid int) error {
	return s.repo.Delete(userId, listid)
}

func (s *TodoListService) Update(userid, listid int, input model.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userid, listid, input)
}
