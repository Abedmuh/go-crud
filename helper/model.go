package helper

import (
	"belajar-go/go-crud/model/domain"
	"belajar-go/go-crud/model/web"
)

func ToCollectionResponse(collection domain.Collection) web.CollectionResponse {
	return web.CollectionResponse{
		Id:   collection.Id,
		Name: collection.Name,
	}
}

func ToCollectionResponses(categories []domain.Collection) []web.CollectionResponse {
	var collectionResponses []web.CollectionResponse
	for _, collection := range categories {
		collectionResponses = append(collectionResponses, ToCollectionResponse(collection))
	}
	return collectionResponses
}
