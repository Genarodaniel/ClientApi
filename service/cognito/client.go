package cognito

import (
	"clientApi/config/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func NewClient() (*cognitoidentityprovider.CognitoIdentityProvider, error) {
	session, err := createStaticSession()
	if err != nil {
		return nil, err
	}

	return cognitoidentityprovider.New(session), nil

}

func createStaticSession() (*session.Session, error) {
	creds := credentials.NewStaticCredentials(env.AwsAccessKeyId, env.AwsSecretKeyId, "")

	return session.NewSession(&aws.Config{
		Region:      aws.String(env.AwsUserPoolRegion),
		Credentials: creds,
	})
}
