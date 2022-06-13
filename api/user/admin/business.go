package admin

import (
	cognitoAdmin "clientApi/service/cognito/admin"
)

type IAdminBusiness interface {
	CreateUser() (int, error)
}

type AdminBusiness struct {
	cognito *cognitoAdmin.AdminImpl
}

func NewAdminBusiness(cognito *cognitoAdmin.AdminImpl) IAdminBusiness {
	return &AdminBusiness{cognito}
}

func (ab *AdminBusiness) CreateUser() (int, error) {
	statusCode := ab.cognito.CreateUser()
	return statusCode, nil
}
