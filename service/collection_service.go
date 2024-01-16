package service

import (
	"context"
	"belajar-go/go-crud/model/web"
)

type CollectionService interface {
	Create(ctx context.Context, request web.CollectionCreateRequest) web.CollectionResponse
	Update(ctx context.Context, request web.CollectionUpdateRequest) web.CollectionResponse
	Delete(ctx context.Context, collectionId int)
	FindById(ctx context.Context, collectionId int) web.CollectionResponse
	FindAll(ctx context.Context) []web.CollectionResponse
}
