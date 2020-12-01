package storage

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	TableUser = "user"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}


func FindUserById(id string) (*User,error){
	u := User{}
	db := GetDB()
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = '%s'", TableUser, id)
	err := db.QueryRow(sql).Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "Find user with id %s error", id)
	}
	return &u, nil
}

func FindUserByIdError(id string) (*User,error){
	u := User{}
	db := GetDB()
	sql := fmt.Sprintf("SELECT * FROM %s WHERESSSSSSSS id = '%s'", TableUser, id)
	err := db.QueryRow(sql).Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "Find user with id %s error", id)
	}
	return &u, nil
}