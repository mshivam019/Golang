package utils

import (
	"errors"
	"gin/basic/model"
	"log"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var tokenBlacklist = struct {
	sync.RWMutex
	tokens map[string]struct{}
}{
	tokens: make(map[string]struct{}),
}

func GenerateToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	myToken, err := token.SignedString([]byte("gotutorial"))
	if err != nil {
		return "", err
	}
	log.Println(myToken)
	// Store the generated token in the blacklist
	tokenBlacklist.Lock()
	defer tokenBlacklist.Unlock()
	tokenBlacklist.tokens[myToken] = struct{}{}
	return myToken, nil
}

func InvalidateToken(userToken string) {
	tokenBlacklist.Lock()
	defer tokenBlacklist.Unlock()
	tokenBlacklist.tokens[userToken] = struct{}{}
}

func IsTokenBlacklisted(userToken string) bool {
	tokenBlacklist.RLock()
	defer tokenBlacklist.RUnlock()
	_, blacklisted := tokenBlacklist.tokens[userToken]
	return blacklisted
}

func ValidateToken(userToken string) (bool, error) {
	if IsTokenBlacklisted(userToken) {
		return false, errors.New("Token is Invalid")
	}

	jwtKeyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Some unexpected error")
		}
		return []byte("gotutorial"), nil
	}
	token, err := jwt.Parse(userToken, jwtKeyfunc)
	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, errors.New("Token is Invalid")
}
