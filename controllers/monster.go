// defines monster CRUD API

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	t "github.com/twosevenska/forge/types"
)

// CreateMonster adds a monster to DB
func CreateMonster(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "monster created successfully",
	})
}

// UpdateMonster adds an existing monster in DB
func UpdateMonster(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "monster updated successfully",
	})
}

// DeleteMonster adds an existing monster in DB
func DeleteMonster(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// FetchMonsters gets existing monsters from DB
func FetchMonsters(c *gin.Context) {
	// TODO: insert DB fetch logic here
	monsters := t.MonsterResult{
		Items: []t.MonsterEntity{},
		BaseResult: t.BaseResult{
			TotalCount:  0,
			CurrentPage: 1,
			TotalPages:  1,
		},
	}

	c.JSON(http.StatusOK, monsters)
}
