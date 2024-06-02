package user

type User struct {
	ID    string
	Email string
	Type  string
}
type UserDomainService struct {
	ur UserRepository
}
