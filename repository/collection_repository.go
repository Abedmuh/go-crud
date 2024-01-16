package repository

import (
	"context"
	"database/sql"
	"belajar-go/go-crud/model/domain"
)

type CollectionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, collection domain.Collection) domain.Collection
	Update(ctx context.Context, tx *sql.Tx, collection domain.Collection) domain.Collection
	Delete(ctx context.Context, tx *sql.Tx, collection domain.Collection)
	FindById(ctx context.Context, tx *sql.Tx, collectionId int) (domain.Collection, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Collection
}
