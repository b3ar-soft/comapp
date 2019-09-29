package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Active         bool      `json:"active"`
	CreateDateTime time.Time `json:"create_date_time"`
	ModifyDateTime time.Time `json:"modify_date_time"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.ID, u.Name, u.Email)
}

func (u User) Insert() error {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DBUser"),
		Password: os.Getenv("DBDBPASS"),
	})
	defer db.Close()
	err := db.Insert(u)
	if err != nil {
		log.Fatalf("Error al insertar al usuario %s.\n%v", u, err)
		return err
	}
	return nil
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{Temp: false})
		if err != nil {
			log.Fatalf("Error al crear la tabla de usuario.\n%v", err)
			return err
		}
	}
	return nil
}
