package models

import "fmt"

var currentId int

var users Users

// Give us some seed data
func init() {
	CreateUser("Jack Thibaut")
	CreateUser("Sae Barceló")
}

func FindUser(id int) User {
	for _, t := range users {
		if t.Id == id {
			return t
		}
	}
	return User{}
}

func FindAllUsers() Users {
	return users
}

func CreateUser(name string) User {
	user := User{Name: name}
	currentId += 1
	user.Id = currentId
	users = append(users, user)
	return user
}

func RepoDestroyUser(id int) error {
	for i, t := range users {
		if t.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}