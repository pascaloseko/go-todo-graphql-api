package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pascaloseko/go-todo-graphql-api/graph"
	"github.com/pascaloseko/go-todo-graphql-api/graph/generated"
	"github.com/pascaloseko/go-todo-graphql-api/graph/model"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+
			" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+
			os.Getenv("PASSWORD"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to db")
	}

	db.LogMode(false)

	db.AutoMigrate(&model.Todo{})
}

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors(srv.ServeHTTP))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
