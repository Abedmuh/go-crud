package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"belajar-go/go-crud/helper"
	"belajar-go/go-crud/model/web"
	"belajar-go/go-crud/service"
	"strconv"
)

type CollectionControllerImpl struct {
	CollectionService service.CollectionService
}

func NewCollectionController(collectionService service.CollectionService) CollectionController {
	return &CollectionControllerImpl{
		CollectionService: collectionService,
	}
}

func (controller *CollectionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	collectionCreateRequest := web.CollectionCreateRequest{}
	helper.ReadFromRequestBody(request, &collectionCreateRequest)

	collectionResponse := controller.CollectionService.Create(request.Context(), collectionCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   collectionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CollectionControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	collectionUpdateRequest := web.CollectionUpdateRequest{}
	helper.ReadFromRequestBody(request, &collectionUpdateRequest)

	collectionId := params.ByName("collectionId")
	id, err := strconv.Atoi(collectionId)
	helper.PanicIfError(err)

	collectionUpdateRequest.Id = id

	collectionResponse := controller.CollectionService.Update(request.Context(), collectionUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   collectionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CollectionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	collectionId := params.ByName("collectionId")
	id, err := strconv.Atoi(collectionId)
	helper.PanicIfError(err)

	controller.CollectionService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CollectionControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	collectionId := params.ByName("collectionId")
	id, err := strconv.Atoi(collectionId)
	helper.PanicIfError(err)

	collectionResponse := controller.CollectionService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   collectionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CollectionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	collectionResponses := controller.CollectionService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   collectionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
