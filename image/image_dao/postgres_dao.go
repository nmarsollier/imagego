package image_dao

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nmarsollier/imagego/tools/db"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/log"
)

var (
	instance *PostgressDao
)

func Get(deps ...interface{}) ImageDao {
	for _, o := range deps {
		if client, ok := o.(ImageDao); ok {
			return client
		}
	}

	if instance == nil {
		client, err := db.GetPostgresClient()
		if err == nil {
			_, err := client.Exec(context.Background(), "SET search_path TO imagego")
			if err != nil {
				log.Get(deps...).Error(err)
			}

			instance = &PostgressDao{
				client: client,
			}
		}
	}

	return instance
}

type PostgressDao struct {
	client *pgxpool.Pool
}

func (r *PostgressDao) Get(key string) (image string, err error) {
	err = r.client.QueryRow(context.Background(), "SELECT image FROM image WHERE id=$1", key).Scan(&image)
	if err != nil {
		return "", errs.NotFound
	}
	return
}

func (r *PostgressDao) Set(key string, image string) (err error) {
	_, err = r.client.Exec(context.Background(), "INSERT INTO image (id, image) VALUES ($1, $2)", key, image)
	if err != nil {
		log.Get().Error(err)
	}

	return
}
