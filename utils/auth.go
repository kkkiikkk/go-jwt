package utils

import (
	"errors"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/kkkiikkk/go-jwt/database"
	"github.com/kkkiikkk/go-jwt/model"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}