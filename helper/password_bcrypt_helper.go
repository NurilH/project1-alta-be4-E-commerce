package helper

import "golang.org/x/crypto/bcrypt"

// function untuk generate plan password menjadi hash
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// func CheckHashPassword(password, hash string) (bool, error) {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
