package service

import (
	"time"

	"github.com/tweeeety/go-gin-dir-sample/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoService struct {
	Db *gorm.DB
}

func NewTodoService() (ts TodoService) {
	todoService := TodoService{}
	todoService.Db = DbOpen()
	return todoService
}

func DbOpen() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil || db == nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Todo{})
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func (ts TodoService) GetAll() []model.Todo {

	var todos []model.Todo

	if ts.Db != nil {
		ts.Db.Order("created_at desc").Find(&todos)
	}
	return todos
}

func (ts TodoService) Add(text string, status string) {
	ts.Db.Create(&model.Todo{Text: text, Status: status})
}

func (ts TodoService) GetOne(id int) model.Todo {
	var todo model.Todo
	ts.Db.First(&todo, id)
	return todo
}

func (ts TodoService) Update(id int, text string, status string) {
	var todo model.Todo
	ts.Db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	ts.Db.Save(&todo)
}

func (ts TodoService) Delete(id int) {
	var todo model.Todo
	ts.Db.First(&todo, id)
	ts.Db.Delete(&todo)
}
