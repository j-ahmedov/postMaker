package postgres

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"postMaker/internal/entity"
)

func NewDB() *bun.DB {
	dsn := "postgres://postgres:myp0stgr3sql@localhost:5432/post_maker_db?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(sqlDb, pgdialect.New())
}

func CreateAllTables(db *bun.DB, ctx context.Context) {
	modelsCreate := []interface{}{
		(*entity.Comment)(nil),
		(*entity.CommentLike)(nil),
		(*entity.Post)(nil),
		(*entity.PostFile)(nil),
		(*entity.PostLike)(nil),
		(*entity.User)(nil),
	}

	for _, model := range modelsCreate {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
