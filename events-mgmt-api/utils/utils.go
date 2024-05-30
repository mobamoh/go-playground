package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(pwdHash), err
}

func CheckPassword(hashedPwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
