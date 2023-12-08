package repository

import (
	"errors"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("Error user not found")
	ErrInvalidLogin = errors.New("Error user invalid login")
)

func Authenticate(username, password string) error {
	hash, err := Client.Get("user:" + username).Bytes()
	if err == redis.Nil {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return ErrInvalidLogin
	}
	return nil

}

func Register(username, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	return Client.Set("users:"+username, hash, 0).Err()
}
