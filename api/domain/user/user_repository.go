package user

type UserRepository interface {
	Find() (*User, error)
	Login() (*User, error)
}
