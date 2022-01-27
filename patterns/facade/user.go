package facade

import "strconv"

type UserService interface {
	Login(phone, code int) (*User, error)
	Register(phone, code int) (*User, error)
}

type UserFacade interface {
	LoginOrRegister(phone, code int) (*User, error)
}

type User struct {
	Username    string
	Password    string
	PhoneNumber string
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) Login(phone, code int) (*User, error) {
	return &User{}, nil
}

func (u *UserServiceImpl) Register(phone, code int) (*User, error) {
	return &User{
		Username:    "test",
		PhoneNumber: strconv.Itoa(phone),
	}, nil
}

var _ UserService = &UserServiceImpl{}

func (s *UserServiceImpl) LoginOrRegister(phone, code int) (*User, error) {
	u, err := s.Login(phone, code)
	if err != nil {
		return nil, err
	}
	if u != nil && u.Username != "" {
		return u, nil
	}
	return s.Register(phone, code)
}

var _ UserFacade = &UserServiceImpl{}
