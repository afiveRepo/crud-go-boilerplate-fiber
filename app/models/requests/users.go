package requests

type UserRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
}
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
