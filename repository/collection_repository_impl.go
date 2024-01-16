package repository

import (
	"context"
	"database/sql"
	"errors"
	"belajar-go/go-crud/helper"
	"belajar-go/go-crud/model/domain"
)

type CollectionRepositoryImpl struct {
}

func NewCollectionRepository() CollectionRepository {
	return &CollectionRepositoryImpl{}
}

func (repository *CollectionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, collection domain.Collection) domain.Collection {
	SQL := "insert into collection(name,point,status) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, collection.Name, collection.Point, collection.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	collection.Id = int(id)
	return collection
}

func (repository *CollectionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, collection domain.Collection) domain.Collection {
	SQL := "update collection set name = ?, point = ?,status = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, collection.Name, collection.Point, collection.Status, collection.Id)
	helper.PanicIfError(err)

	return collection
}

func (repository *CollectionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, collection domain.Collection) {
	SQL := "delete from collection where id = ?"
	_, err := tx.ExecContext(ctx, SQL, collection.Id)
	helper.PanicIfError(err)
}

func (repository *CollectionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, collectionId int) (domain.Collection, error) {
	SQL := "select id, name,point,status from collection where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, collectionId)
	helper.PanicIfError(err)
	defer rows.Close()

	collection := domain.Collection{}
	if rows.Next() {
		err := rows.Scan(&collection.Id, &collection.Name, &collection.Point, &collection.Status)
		helper.PanicIfError(err)
		return collection, nil
	} else {
		return collection, errors.New("collection is not found")
	}
}

func (repository *CollectionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Collection {
	SQL := "select id, name, point, status from collection"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var collections []domain.Collection
	for rows.Next() {
		collection := domain.Collection{}
		err := rows.Scan(&collection.Id, &collection.Name, &collection.Point, &collection.Status)
		helper.PanicIfError(err)
		collections = append(collections, collection)
	}
	return collections
}
