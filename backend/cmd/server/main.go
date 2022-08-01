package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Faroukhamadi/likude/auth"
	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/migrate"
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
	router := chi.NewRouter()

	client := Open("postgresql://postgres:faroukhamadi@127.0.0.1/likude")
	defer client.Close()

	router.Use(auth.Middleware(client))

	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}
	client.User.Create()
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)

	srv := handler.NewDefaultServer(gql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	router.Handle("/",
		playground.Handler("Likude", "/query"),
	)
	router.Handle("/query", srv)
	log.Println("listening on :4000")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatal("http server terminated", err)
	}
}
