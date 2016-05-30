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
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	ibbdSecretKey, err = ioutil.ReadFile(os.Getenv("IBBD_SECRET_KEY_FILE"))
	if err != nil {
		panic(err.Error())
	}
	/*ibbdPublicKey, err = ioutil.ReadFile(os.Getenv("IBBD_PUBLIC_KEY_FILE"))
	if err != nil {
		panic(err.Error())
	}*/
}

func GetUserToken(userId int32, expTime int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	//token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["iss"] = userId
	token.Claims["exp"] = expTime
	token.Claims["iat"] = time.Now().Unix()
	tokenString, err := token.SignedString(ibbdSecretKey)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return tokenString, nil
}

func ParseUserToken(tokenString string) (int32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		/*if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}*/
		if token.Header["alg"] != "HS256" {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return ibbdSecretKey, nil
	})

	if err == nil && token.Valid {
		//fmt.Printf("%+v\n", token.Claims)
		userId := int32(token.Claims["iss"].(float64))
		//fmt.Printf("%T", userId)
		return userId, nil
	}
	return 0, err
}
