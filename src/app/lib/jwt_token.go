package lib

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"time"
	//"strconv"
)

var (
	ibbdSecretKey []byte = nil
	ibbdPublicKey []byte = nil
	//defaultKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return ibbdPublicKey, nil }
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file!")
	}

	ibbdSecretKey, err = ioutil.ReadFile(os.Getenv("IBBD_SECRET_KEY_FILE"))
	if err != nil {
		panic("error")
	}
	ibbdPublicKey, err = ioutil.ReadFile(os.Getenv("IBBD_PUBLIC_KEY_FILE"))
	if err != nil {
		panic("error")
	}
}

func GetUserToken(userId int32, expTime int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["sub"] = userId
	token.Claims["exp"] = expTime
	token.Claims["iat"] = time.Now().Unix()
	tokenString, err := token.SignedString(ibbdSecretKey)
	if err != nil {
		return "", fmt.Errorf("Error: create Token")
	}

	return tokenString, nil
}

func ParseUserToken(tokenString string) (int32, error) {
	//token, err := jwt.Parse(tokenString, defaultKeyFunc)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return ibbdSecretKey, nil
	})

	if err == nil && token.Valid {
		//return int32(token.Claims["user"]), nil
		return 0, nil
	}
	return 0, fmt.Errorf("Error: parse token")
}
