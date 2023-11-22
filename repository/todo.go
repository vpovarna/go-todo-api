package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/vpovarna/go-todo-api/domain"
	"time"
)

type TodoRepository struct {
	conn *sqlx.DB
}

func NewTodoStorage(conn *sqlx.DB) *TodoRepository {
	return &TodoRepository{conn: conn}
}

type Todo struct {
	Id          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Completed   bool      `db:"completed"`
	CreatedAt   time.Time `db:"created_at"`
	CompletedAt time.Time `db:"completed_at"`
}

func (t *TodoRepository) CreateTodo(todoDAO domain.TodoDAO) (int, error) {
	stmt := "INSERT INTO todos (title, description, completed, created_at, completed_at) VALUES (?,?,?,?,?)"

	exec, err := t.conn.Exec(stmt, todoDAO.Title, todoDAO.Description, false, time.Now(), time.Now())
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

func (t *TodoRepository) GetTodoById(todoId int) (Todo, error) {
	todo := Todo{}
	stmt := "SELECT id, title, description, completed, created_at, completed_at FROM todos WHERE id = ?"

	err := t.conn.GetContext(context.Background(), &todo, stmt, todoId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return Todo{}, err
		}
		return Todo{}, errors.New("record not found")
	}

	return todo, nil
}

func (t *TodoRepository) CompleteTodo(todoId int) error {
	stmt := "UPDATE todos SET completed = true, completed_at = NOW() WHERE id = ?"

	exec, err := t.conn.Exec(stmt, todoId)
	if err != nil {
		log.Warn("Unable to complete todo with id: ", todoId)
		return err
	}

	_, err = exec.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoRepository) DeleteTodo(todoId int) error {
	stmt := "DELETE todos WHERE id = ?"
	exec, err := t.conn.Exec(stmt, todoId)

	if err != nil {
		log.Warn("Unable to delete todo with id: ", todoId)
		return err
	}

	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	log.Info("Successfully deleted: ", rowsAffected)
	return nil
}
