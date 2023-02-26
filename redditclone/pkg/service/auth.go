package service

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/repository"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"time"
)

const (
	Day             = time.Hour * 24
	salt            = "random"
	signingKey      = "heheKey"
	tokenTTL        = 7 * Day
	expiredTokenMSG = "token is expired"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenUser struct {
	Username string `json:"username"`
	ID       int    `json:"id,,string"`
}

func newTokenUser(user structs.User) tokenUser {
	return tokenUser{
		Username: user.Username,
		ID:       user.ID,
	}
}

type tokenClaims struct {
	User tokenUser `json:"user"`
	IAT  int64     `json:"iat"`
	EXP  int64     `json:"exp"`
}

func (t tokenClaims) Valid() error {
	if t.EXP > time.Now().Unix() {
		return nil
	}
	err := structs.NewErrorMap()
	err["msg"] = expiredTokenMSG
	msg, er := json.Marshal(err)
	if er != nil {
		return errors.New(expiredTokenMSG)
	}
	return errors.New(string(msg))
}

func newTokenClaims(user structs.User, iat, exp int64) *tokenClaims {
	return &tokenClaims{
		User: newTokenUser(user),
		IAT:  iat,
		EXP:  exp,
	}
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structs.User) (string, error) {
	tmpPass := user.Password
	user.Password = generatePasswordHash(user.Password)
	_, err := s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}
	return s.GenerateToken(user.Username, tmpPass)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	iat := time.Now().Unix()
	exp := time.Now().Add(tokenTTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newTokenClaims(user, iat, exp))
	err = s.repo.GenerateToken(username, user.ID, iat, exp)
	if err != nil {
		return "", err
	}
	return token.SignedString([]byte(signingKey))
}

func hashSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != "HS256" {
		return nil, fmt.Errorf("bad sign method")
	}
	return []byte(signingKey), nil
}

func (s *AuthService) CheckIdentity(username string, id int) error {
	return s.repo.CheckIdentity(username, id)
}
func (s *AuthService) ParseToken(accessToken string) (structs.Author, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, hashSecretGetter)

	if err != nil {
		return structs.Author{}, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return structs.Author{}, errors.New("token claims are not of type *tokenClaims")

	}
	if s.CheckIdentity(claims.User.Username, claims.User.ID) != nil {
		return structs.Author{}, errors.New("this token doesn`t match any user")
	}
	return structs.Author{
		ID:       claims.User.ID,
		Username: claims.User.Username,
	}, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
