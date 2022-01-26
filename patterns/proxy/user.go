package proxy

// todo 使用newXXX创建和直接使用结构体创建有什么区别？

type IUser interface {
	Login(username, password string) error
}

type User struct {
}

func (u *User) Login(username, password string) error {
	// ...
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (p *UserProxy) Login(username, password string) error {
	//....before

	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// ...after

	return nil
}
