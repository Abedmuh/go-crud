package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"belajar-go/go-crud/exception"
	"belajar-go/go-crud/helper"
	"belajar-go/go-crud/model/domain"
	"belajar-go/go-crud/model/web"
	"belajar-go/go-crud/repository"
)

type CollectionServiceImpl struct {
	CollectionRepository repository.CollectionRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewCollectionService(collectionRepository repository.CollectionRepository, DB *sql.DB, validate *validator.Validate) CollectionService {
	return &CollectionServiceImpl{
		CollectionRepository: collectionRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CollectionServiceImpl) Create(ctx context.Context, request web.CollectionCreateRequest) web.CollectionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	collection := domain.Collection{
		Name: request.Name,
	}

	collection = service.CollectionRepository.Save(ctx, tx, collection)

	return helper.ToCollectionResponse(collection)
}

func (service *CollectionServiceImpl) Update(ctx context.Context, request web.CollectionUpdateRequest) web.CollectionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	collection, err := service.CollectionRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	collection.Name = request.Name

	collection = service.CollectionRepository.Update(ctx, tx, collection)

	return helper.ToCollectionResponse(collection)
}

func (service *CollectionServiceImpl) Delete(ctx context.Context, collectionId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	collection, err := service.CollectionRepository.FindById(ctx, tx, collectionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CollectionRepository.Delete(ctx, tx, collection)
}

func (service *CollectionServiceImpl) FindById(ctx context.Context, collectionId int) web.CollectionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	collection, err := service.CollectionRepository.FindById(ctx, tx, collectionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCollectionResponse(collection)
}

func (service *CollectionServiceImpl) FindAll(ctx context.Context) []web.CollectionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CollectionRepository.FindAll(ctx, tx)

	return helper.ToCollectionResponses(categories)
}
