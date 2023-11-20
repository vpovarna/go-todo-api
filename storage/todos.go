package storage

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/vpovarna/go-todo-api/domain"
	"time"
)

type TodoStorage struct {
	Conn *sqlx.DB
}

func NewTodoStorage(conn *sqlx.DB) *TodoStorage {
	return &TodoStorage{Conn: conn}
}

type Todo struct {
	id          int
	title       string
	description string
	completed   bool
	createdAt   time.Time
	completedAt time.Time
}

func (t *TodoStorage) CreateTodo(todoDAO domain.TodoDAO) (int, error) {
	stmt := "INSERT INTO todos (title, description, completed, created_at, completed_at) VALUES (?,?,?,?,?)"

	exec, err := t.Conn.Exec(stmt, todoDAO.Title, todoDAO.Description, false, time.Now(), time.Now())
	if err != nil {
		log.Info("Unable to create todo: ", todoDAO)
		return 0, err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *TodoStorage) GetTodoById(todoId int) (Todo, error) {
	todo := Todo{}
	stmt := "SELECT id, title, description, completed, created_at, completed_at FROM todos WHERE id = ?"

	err := t.Conn.Get(&todo, stmt, todoId)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (t *TodoStorage) CompleteTodo(todoId int) error {
	stmt := "UPDATE todos SET completed = true, completed_at = NOW() WHERE id = ?"

	exec, err := t.Conn.Exec(stmt, todoId)
	if err != nil {
		log.Info("Unable to complete todo with id: ", todoId)
		return err
	}

	_, err = exec.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
