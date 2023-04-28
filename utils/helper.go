package utils

import (
	"crypto/sha256"
	"math/rand"

	"github.com/pkg/errors"
)

// GeneratePid returns a pid of length n with a given prefix
func GeneratePid(prefix string, length int64) string {
	pid := prefix + "_" + GenerateRandomString(int(length))
	return pid
}

// GenerateRandomString returns a alphanmeric string of length n
func GenerateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return salt, errors.Wrap(err, "[GenerateSalt]")
	}
	return salt, nil
}

func GenerateHashedPassword(salt []byte, password string) []byte {
	saltedPassword := append(salt, []byte(password)...)
	passwordHash := sha256.Sum256(saltedPassword)
	return passwordHash[:]
}
