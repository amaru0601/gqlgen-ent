package db

import (
	"context"
	"log"

	"entgo.io/bug/ent"
	_ "github.com/lib/pq"
)

func GetClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5434 dbname=test user=postgres password=pass sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
		return nil, err
	}

	return client, nil
}
