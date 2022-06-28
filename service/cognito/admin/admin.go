package admin

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Admin interface {
	CreateUser() int
}

type AdminImpl struct {
	cli *cognitoidentityprovider.CognitoIdentityProvider
}

func NewCognitoAdmin(cli *cognitoidentityprovider.CognitoIdentityProvider) *AdminImpl {
	return &AdminImpl{cli}
}

func (adm *AdminImpl) CreateUser() int {
	return http.StatusOK
}
