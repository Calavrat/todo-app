package service

import (
	"github.com/Calavrat/todo-app/model"
	"github.com/Calavrat/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item model.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err

	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]model.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (model.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userid, itemId int, input model.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userid, itemId, input)
}