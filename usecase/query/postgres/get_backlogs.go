package postgres

import (
	"context"
	"database/sql"

	backlogitem "github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type PostgresQueries struct {
	db sql.DB
}

func (p PostgresQueries) GetBacklogItems(ctx context.Context, query *backlogitem.BacklogItemQuery) (backlogitem.BacklogItemNodeList, error) {
	//queries := sqlc.New(&p.db)
	//rows, err := queries.GetBacklogItems(ctx)
	//if err != nil {
	//	return nil, err
	//}
	panic("todo: impl")

}
