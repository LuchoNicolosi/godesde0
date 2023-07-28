package modelos

import (
	"time"
)

//En el paquete models, o modelos se colocan todas las definiciones o todas las estructuras en un solo lugar

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (u *User) AddUser(id int, name string, createdAt time.Time, status bool) { // -> * es un puntero de memoria
	u.Id = id
	u.Name = name
	u.CreatedAt = createdAt
	u.Status = status
}
