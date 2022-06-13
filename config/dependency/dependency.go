package dependency

import (
	"clientApi/service/cognito"
	"clientApi/service/cognito/admin"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
	CognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	CognitoAdmin  *admin.AdminImpl
)

func Load() error {
	var err error

	CognitoClient, err = cognito.NewClient()
	if err != nil {
		return err
	}
	CognitoAdmin = admin.NewCognitoAdmin(CognitoClient)

	return nil
}
