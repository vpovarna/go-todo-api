package handlers

import (
	"fmt"
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
		"status": false,
		"errors": message,
	}
}

func (h *TodoHandlers) GetTodoById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warn("Invalid id: ", idStr)
		return c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid provided id"))
	}

	// id validation
	if id <= 0 {
		log.Warn("Invalid provided id: ", id)
		return c.Status(http.StatusNotFound).JSON(h.errorMessageResponse("Invalid provided id"))
	}

	todo, err := h.todoRepository.GetTodoById(id)
	if err != nil {
		log.Warn("Unable to get a todo from the specified id: ", id)
		return c.Status(http.StatusNotFound).JSON(h.errorMessageResponse("Can't get a todo for the specified id"))
	}

	todoDAO := domain.TodoDAO{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
	}

	log.Infof("Fetched todo: %v from the repository", todoDAO)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "true",
		"todo":   todoDAO,
	})
}

func (h *TodoHandlers) CreateTodo(c *fiber.Ctx) error {
	todo := new(domain.TodoDAO)

	if err := c.BodyParser(todo); err != nil {
		log.Warn("Unable to parse input body. Error: ", err.Error())
		err = c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid body"))
		return err
	}

	todoId, err := h.todoRepository.CreateTodo(*todo)
	if err != nil {
		log.Error("Unable to create a new todo")
		err = c.Status(http.StatusInternalServerError).JSON(h.errorMessageResponse("Unable to create a new todo"))
		return err
	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("A todo with the id %d has been successfully created!", todoId),
	})
}

func (h *TodoHandlers) CompleteTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warn("Invalid id: ", idStr)
		return c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid id"))
	}

	err = h.todoRepository.CompleteTodo(id)
	if err != nil {
		log.Error("Unable to complete a new todo")
		return c.Status(http.StatusInternalServerError).JSON(h.errorMessageResponse("Unable to complete a new todo"))
	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Todo id: %d completed successfully", id),
	})

}

func (h *TodoHandlers) DeleteTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warn("Invalid id: ", idStr)
		return c.Status(http.StatusBadRequest).JSON(h.errorMessageResponse("Invalid id"))
	}

	if err := h.todoRepository.DeleteTodo(id); err != nil {
		log.Error("Unable to delete a new todo")
		return c.Status(http.StatusInternalServerError).JSON(h.errorMessageResponse("Unable to delete a new todo"))
	}

	return c.Status(http.StatusNoContent).JSON(&fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Todo: %d successfully deleted", id),
	})
}

func (h *TodoHandlers) GetAllTodos(c *fiber.Ctx) error {

	todos, err := h.todoRepository.GetAllTodos()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(h.errorMessageResponse("Unable to fetch all todos from the repository"))
	}

	var todoDAOs []domain.TodoDAO
	for _, todo := range todos {
		todoDAOs = append(todoDAOs, domain.TodoDAO{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": true,
		"todos":  &todoDAOs,
	})
}
