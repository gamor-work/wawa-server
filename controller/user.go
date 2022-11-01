package controller

import (
	Msql "gin/databases"
	response "gin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UserHandler(c *gin.Context) {
	var persons []Person
	rows, _ := Msql.DB.Query("select * from user")
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err)
		}

		persons = append(persons, person)
	}

	c.JSON(http.StatusOK, response.OK.WithData(persons))
}
