package models

import (
	"fmt"
	"log"
	"time"

	"comuniapp/config"
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
	db := config.GetDbConnection()
	err := db.Insert(u)
	if err != nil {
		log.Fatalf("Error al insertar al usuario %s.\n%v", u, err)
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	db := config.GetDbConnection()
	err := db.Model(&users).Select()
	if err != nil {
		log.Fatalf("Erro al conseguir la lista de usuarios.\n%v", err)
		return nil, err
	}
	return users, nil
}

func GetUserById(id int64) (*User, error) {
	user := &User{ID: id}
	db := config.GetDbConnection()
	err := db.Select(user)
	if err != nil {
		log.Fatalf("Error al encontrar el usuario con id: %d.\n%v", id, err)
		return nil, err
	}
	return user, nil
}

func (u User) Update() error {
	db := config.GetDbConnection()
	u.ModifyDateTime = time.Now()
	err := db.Update(u)
	if err != nil {
		log.Fatalf("Error al modificar el usuario %s.\n%v", u, err)
		return err
	}
	return nil
}

func (u User) Delete() error {
	db := config.GetDbConnection()
	u.Active = false
	u.ModifyDateTime = time.Now()
	err := db.Update(u)
	if err != nil {
		log.Fatalf("Error al eliminar al usuario %s.\n%v", u, err)
		return err
	}
	return nil
}
