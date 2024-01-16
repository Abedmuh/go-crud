package app

import (
	"github.com/julienschmidt/httprouter"
	"belajar-go/go-crud/controller"
	"belajar-go/go-crud/exception"
)

func NewRouter(collectionController controller.CollectionController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/collection", collectionController.FindAll)
	router.GET("/collection/:collectionId", collectionController.FindById)
	router.POST("/collection", collectionController.Create)
	router.PUT("/collection/:collectionId", collectionController.Update)
	router.DELETE("/collection/:collectionId", collectionController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
