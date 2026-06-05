/*
Autor: Luis Miguel Vasconez
Fecha: 28/05/2026
Descripcion: Incio desarrollo clase user
*/

package models

type User struct {
	id    int
	name  string
	email string
}

// Constructor
func NewUser(id int, name string, email string) User {
	return User{
		id:    id,
		name:  name,
		email: email,
	}
}

// Getters
func (u User) GetName() string {
	return u.name
}

func (u User) GetEmail() string {
	return u.email
}
