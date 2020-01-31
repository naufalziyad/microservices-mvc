package services

import (
	"../domain"
	"../utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)

}
