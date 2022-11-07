package models

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

const TokenString = "secret"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`
}

func (u *User) Validate() error {
	fmt.Println(u)
	if len(u.FullName) < 3 || u.Age < 18 || strings.Contains("@", u.Email) || len(u.Email) < 3 || len(u.Password) < 4 {
		return errors.New("bad request")
	}
	return nil
}

func (u *User) EncryptPassword() {
	u.Password, _ = HashPassword(u.Password)
}

func (u *User) CheckPassword(password string) bool {
	return CheckPasswordHash(password, u.Password)
}

type CustomClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func (u *User) CreateClaim() string {
	claims := CustomClaim{
		Name: u.FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour * 24),
			Issuer:    strconv.Itoa(int(u.ID)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(TokenString))
	return ss
}
