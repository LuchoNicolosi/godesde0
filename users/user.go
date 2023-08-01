package users

import (
	"fmt"
	"time"

	"github.com/luchonicolosi/godesde0/modelos"
)

func AltaUsuario() {
	user := new(modelos.User)
	user.AddUser(10, "Lucho", time.Now(), true)
	fmt.Println(user)
}
