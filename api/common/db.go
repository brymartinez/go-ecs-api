package common

import (
	"context"
	"fmt"
	"os"

	pg "github.com/go-pg/pg/v10"
)

var storage *pg.DB

func ConnectToDB() (*pg.DB, error) {
	if storage != nil {
		return storage, nil
	}
	fmt.Println("Creating new db connection...")
	opt, err := pg.ParseURL(os.Getenv("DB_CONNSTRING"))
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)

	err = db.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	storage = db

	return db, nil
}
