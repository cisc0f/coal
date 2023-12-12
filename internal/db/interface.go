package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	dbUrl            = "https://coal-f8d25-default-rtdb.firebaseio.com/"
	firebaseAuthPath = "serviceAccount.json"
)

// Firebase client abstraction
type DB struct {
	*db.Client
}

type Storage struct {
	*storage.Client
}

// Firebase response structure
type Song struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	SongUrl string `json:"song"`
}

// Firebase DB setup
func (db *DB) SetupDB() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile(firebaseAuthPath)
	config := &firebase.Config{DatabaseURL: dbUrl}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing firebase app: %v", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing firebase database: %v", err)
	}

	db.Client = client
	return nil
}

// Firebase Storage setup
func (s *Storage) SetupStorage() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile(firebaseAuthPath)
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return fmt.Errorf("error initializing firebase storage: %v", err)
	}

	s.Client = client
	return nil
}
