package service

import "github.com/neverhover/Go-000/Week02/storage"

func GetUser(id string) (interface{} ,error) {
	return storage.FindUserById(id)
}

func GetUserError(id string) (interface{} ,error) {
	return storage.FindUserByIdError(id)
}