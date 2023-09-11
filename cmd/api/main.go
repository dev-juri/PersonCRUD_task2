package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type config struct {
	port int
}
type application struct {
	config   config
	logger   *log.Logger
	dbClient *mongo.Client
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.Parse()

	dbString := os.Getenv("DB_STRING")
	if dbString != "" {
		dbString = "mongodb://localhost:27017"
	}
	fmt.Println(dbString)

	client, ctx, cancel, err := connect(dbString)
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	ping(client, ctx)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		cfg,
		logger,
		client,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}
