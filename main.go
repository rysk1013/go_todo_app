package main

import (
	"log"
	"strconv"

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

// DB delete
// func dbDelete(id int) {
// 	db, err := gorm.Open("sqlite3", "test.sqlite3")
// 	if err != nil {
// 		log.Fatalln("can't open dbDelete", err)
// 	}
// 	defer db.Close()
// 	var todo Todo
// 	db.First(&todo, id)
// 	db.Delete(&todo)
// }

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	//Index
	router.GET("/", func (ctx *gin.Context)  {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	//Create
	router.POST("/new", func (ctx *gin.Context)  {
		text := ctx.PostForm("text")
		dbInsert(text)
		ctx.Redirect(302, "/")
	})

	//Get show
	router.GET("/todos/:id", func (ctx *gin.Context)  {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln("can't find show", err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "show.html", gin.H{"todo": todo})
	})

	//Update
	router.POST("/update/:id", func (ctx *gin.Context)  {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln("can't update", err)
		}
		text := ctx.PostForm("text")
		dbUpdate(id, text)
		ctx.Redirect(302, "/")
	})

	router.Run() // localhost:8080
}