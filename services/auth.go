package services

type AuthService struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (this AuthService) Login() (err error) {
	if this.Account == "elsa" && this.Password == "12345" {

	}
	return
}
