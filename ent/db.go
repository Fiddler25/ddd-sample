package ent

import (
	"context"
	"ddd-sample/ent/migrate"
	"log"
)

func New() (context.Context, *Client) {
	client, err := Open("mysql", "root@tcp(127.0.0.1:13306)/ddd-sample?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return ctx, client
}
