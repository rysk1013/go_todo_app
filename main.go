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
		log.Fatalln("can't open DB", err)
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
}

// DB insert
func dbInsert(text string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		log.Fatalln("can't open DB", err)
	}
	defer db.Close()
	db.Create(&Todo{Text: text})
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