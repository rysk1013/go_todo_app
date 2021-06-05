package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text string
}

// DB init
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Fatalln("can't open dbInit", err)
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
}

// DB insert
func dbInsert(text string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Fatalln("can't open dbInsert", err)
	}
	defer db.Close()
	db.Create(&Todo{Text: text})
}

// DB get all
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Fatalln("can't open dbGetAll", err)
	}
	defer db.Close()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	return todos
}

// DB get one
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Fatalln("can't open dbGetOne", err)
	}
	defer db.Close()
	var todo Todo
	db.First(&todo, id)
	return todo
}

// DB update
func dbUpdate(id int, text string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Fatalln("can't open dbUpdate", err)
	}
	defer db.Close()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	db.Save(&todo)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin"

	router.GET("/", func (ctx *gin.Context)  {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run() // localhost:8080
}