package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashPass, pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass)); err != nil {
		return false
	}

	return true
}
