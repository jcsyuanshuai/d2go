package facebook

// User

type UserStatus int

type User struct {
	Username   string
	Password   string
	Email      string
	Company    string
	Address    string
	UserStatus UserStatus
}

type AdminBehavior interface {
}

type Admin User

type MemberBehavior interface {
	SendMessage(msg []byte)
	CreatePost(post string)
}

type Member User
