package entity_controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"mysvc/src/domain/entity"
)

type EntityController interface {
	GetByID(c *gin.Context)
	CreateEntity(c *gin.Context)
}

type entityController struct {
	service entity.Service
}

func NewEntityController(svc entity.Service) EntityController {
	return &entityController{
		service: svc,
	}
}

func (e *entityController) GetByID(c *gin.Context) {
	idStr := c.Param("entity_id")
	id, errC := strconv.ParseInt(idStr, 10, 64)
	if errC != nil {
		fmt.Printf("Invalid id :: %v\n", errC)
		c.JSON(400, "Invalid_data")
	}
	ent, err := e.service.GetByID(id)
	if err != nil {
		fmt.Printf("Failed to get entity :: %v\n", err)
		c.JSON(500, "Internal_server_error")
	}
	c.JSON(200, ent)
}

func (e *entityController) CreateEntity(c *gin.Context) {
	var ent entity.Entity
	err := c.BindJSON(&ent)
	if err != nil {
		fmt.Printf("Json bind error :: %v\n", err)
		c.JSON(400, "Invalid_input")
	}
	errC := e.service.CreateEntity(ent)
	if errC != nil {
		fmt.Printf("Entry %v creation failed :: %v\n", ent, errC)
		c.JSON(500, "Internal_error")
	}
	fmt.Printf("Entity %v posted\n", ent)
	c.JSON(201, "Entity_created")
}
