package app

import (
	"fmt"

	"mysvc/src/controller/entity_controller"
	"mysvc/src/domain/entity"
	"mysvc/src/repository/db/memstore"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	fmt.Println("Starting app")
	memRepo := memstore.NewMemStoreRepository()
	entitySvc := entity.NewService(memRepo)
	entityHandler := entity_controller.NewEntityController(entitySvc)

	router.GET("/entity/:entity_id", entityHandler.GetByID)
	router.POST("/entity/", entityHandler.CreateEntity)

	router.Run(":8080")
}
