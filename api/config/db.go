package config

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"
	"os"

	"comuniapp/models"
)

func GetDbConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DBUser"),
		Password: os.Getenv("DBDBPASS"),
	})
	defer db.Close()
	return db
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*models.User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{Temp: false, IfNotExists: true})
		if err != nil {
			log.Fatalf("Error al crear la tabla de usuario.\n%v", err)
			return err
		}
	}
	return nil
}
