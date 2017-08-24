// defines monster CRUD API

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/twosevenska/forge/mongo"
	t "github.com/twosevenska/forge/types"
)

const (
	monsterCollection = "critters"
)

// UpsertMonster adds a monster to DB
func UpsertMonster(c *gin.Context) {
	m := c.MustGet("mongo").(mongo.Client)

	e := t.MonsterEntity{
		BaseEntity: t.BaseEntity{
			ID:   uuid.NewV4(),
			Name: "test",
		},
	}

	if err := m.UpsertMonster(e, monsterCollection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "monster created successfully",
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
