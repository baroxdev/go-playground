package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (TodoItem) TableName() string { return "todo_items" }

func main() {
	dsn := "host=localhost user=todo_admin password=admin dbname=tododb port=2345 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connecto Postgres:", err)
	}

	log.Println("Connected:", db)

	// Auto-migrate the schema
	db.AutoMigrate(&TodoItem{})
	log.Println("Database migrated successfully")

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", createItem(db))  // create item
		v1.GET("/items", getListItems(db)) // get list items
	}

	router.Run()
}

func createItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem TodoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dataItem.Title = strings.TrimSpace((dataItem.Title))

		if dataItem.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
			return
		}

		dataItem.Status = "Doing"

		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}

func getListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Paging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"total"`
		}

		var paging Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []TodoItem

		if err := db.Table(TodoItem{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
