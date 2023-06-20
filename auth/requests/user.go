package requests

type PaginPara struct {
	Pages   string
	PerPage string
}

type UserLoginRequest struct {
	UserName string
	Password string
}

type UserRequest struct {
	Name     string
	Email    string
	Phone    string
	Password string
}
