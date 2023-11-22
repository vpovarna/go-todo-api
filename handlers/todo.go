package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vpovarna/go-todo-api/domain"
	"github.com/vpovarna/go-todo-api/repository"
	"net/http"
	"strconv"
)

type TodoHandlers struct {
	todoRepository *repository.TodoRepository
}

func NewTodoHandlers(todoRepository *repository.TodoRepository) *TodoHandlers {
	return &TodoHandlers{todoRepository: todoRepository}
}

func (h *TodoHandlers) errorMessageResponse(message string) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": message,
	}
}

func (h *TodoHandlers) GetTodoById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warn("Invalid id: ", idStr)
		_ = c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid id"))
		return err
	}

	// id validation
	if id <= 0 {
		log.Warn("Invalid id: ", id)
		_ = c.Status(http.StatusNotFound).JSON(h.errorMessageResponse("Invalid id"))
		return errors.New("invalid id")
	}

	todo, err := h.todoRepository.GetTodoById(id)
	if err != nil {
		log.Warn("Unable to get a todo from the specified id: ", id)
		_ = c.Status(http.StatusNotFound).JSON(h.errorMessageResponse("can't get any todo item"))
		return err
	}

	todoDAO := domain.TodoDAO{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	_ = c.Status(http.StatusOK).JSON(&fiber.Map{
		"todo": todoDAO,
	})

	return nil
}

func (h *TodoHandlers) CreateTodo(c *fiber.Ctx) error {
	t := new(domain.TodoDAO)

	if err := c.BodyParser(t); err != nil {
		log.Info("Unable to parse input body")
		_ = c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid body"))
		return err
	}

	_, err := h.todoRepository.CreateTodo(*t)
	if err != nil {
		log.Error("Unable to create a new todo")
		_ = c.Status(http.StatusInternalServerError).JSON(h.errorMessageResponse("Unable to create a new todo"))
		return err
	}

	err = c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  true,
		"message": "Todo created",
	})

	return err
}

func (h *TodoHandlers) CompleteTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warn("Invalid id: ", idStr)
		_ = c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid id"))
		return err
	}

	err = h.todoRepository.CompleteTodo(id)
	if err != nil {
		log.Error("Unable to complete a new todo")
		_ = c.Status(http.StatusInternalServerError).JSON(h.errorMessageResponse("Unable to complete a new todo"))
		return err
	}

	return nil
}
