package users

import "slices"

type Users []User

type UserService struct {
	users Users
}

func (service *UserService) Add(user User) {
	if service.users == nil {
		service.users = Users{}
	}
	service.users = append(service.users, user)
}

func (service *UserService) Delete(user User) {
	if service.users != nil {
		service.users = slices.DeleteFunc(service.users, func(u User) bool {
			return u == user
		})
	}
}

func (service *UserService) FindAll() Users {
	return service.users
}

func (service *UserService) FindByEmail(email string) (bool, User) {
	index := slices.IndexFunc(service.users, func(u User) bool {
		return u.Email == email
	})
	if index == -1 {
		return false, User{}
	} else {
		return true, service.users[index]
	}
}
