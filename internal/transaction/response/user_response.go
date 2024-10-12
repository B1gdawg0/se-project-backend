package response

import "github.com/jackc/pgx/v5/pgtype"

type CreateUserResponse struct{
	ID 		 pgtype.UUID `json:"ID"`
	Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
    Phone    string `json:"phone"`
}

type FindUserResponse CreateUserResponse

type FindUsersResponse struct{
	Users []FindUserResponse
}