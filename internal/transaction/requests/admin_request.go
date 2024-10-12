package requests

type RegisterAdminRequest struct{
    Email    string `json:"email"`
    Password string `json:"password"`
}