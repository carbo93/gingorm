package controllers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/alexandercrosson/gingorm/db"
    "github.com/alexandercrosson/gingorm/models"
)

var err error

func DeletePerson(c *gin.Context) {
    id := c.Params.ByName("id")
    var person models.Person
    var getDB = db.GetDB()
    d := getDB.Where("id = ?", id).Delete(&person)
    fmt.Println(d)
    c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {
    var person models.Person
    id := c.Params.ByName("id")
    var getDB = db.GetDB()
    if err := getDB.Where("id = ?", id).First(&person).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    }
    c.BindJSON(&person)
    getDB.Save(&person)
    c.JSON(200, person)
}

func CreatePerson(c *gin.Context) {
    var person models.Person
    var getDB = db.GetDB()

    c.BindJSON(&person)
    getDB.Create(&person)
    c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
    id := c.Params.ByName("id")
    var person models.Person
    var getDB = db.GetDB()
    if err := getDB.Where("id = ?", id).First(&person).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, person)
    }
}

func GetPeople(c *gin.Context) {
    var people []models.Person
    var getDB = db.GetDB()
    if err := getDB.Find(&people).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, people)
    }
}