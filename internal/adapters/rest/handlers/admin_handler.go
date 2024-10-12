package handlers

import (
	"github.com/B1gdawg0/se-project-backend/internal/usecases/admin"
)

type AdminRestHandler struct {
	usecase admin.AdminUseCase
}


func ProvideAdminRestHandler(usecase admin.AdminUseCase) *AdminRestHandler {
    return &AdminRestHandler{
        usecase: usecase,
    }
}

func (arh *AdminRestHandler) InitializeAdminAccount() error{
	_, err := arh.usecase.InitializeAdminAccount()	
	if err != nil{
		return err
	}
	return nil
}
