package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jt "github.com/jwt-example/internal/jwt"
	"github.com/jwt-example/internal/jwt/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (srv *Service) CreateToken(userID uint64) (*model.TokenDetails, error) {
	return jt.CreateToken(userID)
}

func (srv *Service)  ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func (srv *Service)  VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := srv.ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (srv *Service)  TokenValid(r *http.Request) error {
	token, err := srv.VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func (srv *Service)  ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := srv.VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:   userId,
		}, nil
	}

	return nil, err
}

