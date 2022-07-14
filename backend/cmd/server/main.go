package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/graphql/gql"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://postgres:faroukhamadi@127.0.0.1/likude")
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	client.User.Create()
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)

	srv := handler.NewDefaultServer(gql.NewSchema(client))
	http.Handle("/",
		playground.Handler("Likude", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
