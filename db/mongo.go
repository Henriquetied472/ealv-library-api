package db

import (
	"context"
	"ealv-library-api/book"
	"encoding/json"
	"os"
	"time"

	"github.com/henriquetied472/logplus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var cancel context.CancelFunc
var client mongo.Client

func connect() mongo.Client {
	ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("LIB_URI")))
	if err != nil {
		logplus.Fatalf("Can't connect to the server: %v | %v", err, os.Getenv("LIB_URI"))
		os.Exit(1)
	}

	return *client
}

func AddBook(b book.Book) {
	client = connect()
	defer client.Disconnect(ctx)
	defer cancel()

	books := client.Database("library").Collection("books")

	books.InsertOne(ctx, b)
}

func RemoveBook(b book.Book) {
	client = connect()
	defer client.Disconnect(ctx)
	defer cancel()

	books := client.Database("library").Collection("books")

	books.DeleteOne(ctx, b)
}

func EditBook(ab book.Book, bb book.Book) {
	client = connect()
	defer client.Disconnect(ctx)
	defer cancel()

	books := client.Database("library").Collection("books")

	books.ReplaceOne(ctx, ab, bb)
}

func ListBooksJSON() string {
	client = connect()
	defer client.Disconnect(ctx)
	defer cancel()

	books := client.Database("library").Collection("books")

	cursor, err := books.Find(ctx, book.Book{})
	if err != nil {
		logplus.Fatal(err.Error())
		os.Exit(0)
	}

	var lBook []book.Book
	var jBooks string

	if err = cursor.All(ctx, &lBook); err != nil {
		logplus.Fatal(err.Error())
		os.Exit(0)
	}

	for _, b := range lBook {
		jb, _ := json.Marshal(b)
		jBooks += string(jb)
	}

	return jBooks
}
