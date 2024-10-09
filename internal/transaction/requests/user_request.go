package requests


type CreateUserRequest struct{
	Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
    Phone    string `json:"phone"`
}
