package response

type LoginResponse struct{
	Token string `json:"token"`
}

type RegisterResponse struct{
	ID 		 string `json:"ID"`
	Email    string `json:"email"`
    Name     string `json:"name"`
}