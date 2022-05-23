package utils

import (
	"crypto/rand"
	"math/big"
	math "math/rand"
	"time"

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

const otpChars = "1234567890"

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func GenerateId() string {
	r := math.New(math.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}
