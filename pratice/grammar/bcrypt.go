package grammar

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptPassword() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	login_password := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(login_password))
	if err != nil {
		println("You can't login")
		return
	}
	println("You're logged in")

}
