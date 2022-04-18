package utils

import (
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a hashed password
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// CheckPasswordHash validates hashed passwords
func VerifyPassword(hash []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

// CheckPasswordHash validates hashed passwords
func GeneratePassword() string {
	p := make([]byte, 16)
	p[10] = 0xFF
	return new(big.Int).SetBytes(p).Text(62)
}
