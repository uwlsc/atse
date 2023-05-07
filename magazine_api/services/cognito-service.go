package services

import (
	"context"
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/utils"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var jwksURL = ""
var issuer = ""
var keySet jwk.Set = jwk.NewSet()

type CognitoAuthService struct {
	client *cognitoidentityprovider.Client
	env    lib.Env
	logger lib.Logger
}

func NewCognitoAuthService(
	client *cognitoidentityprovider.Client,
	env lib.Env,
	logger lib.Logger,
) CognitoAuthService {
	issuer = "https://cognito-idp." + env.AWSRegion + ".amazonaws.com/" + env.PoolID
	jwksURL = issuer + "/.well-known/jwks.json"

	keySet, _ = jwk.Fetch(context.Background(), jwksURL)

	return CognitoAuthService{
		client: client,
		env:    env,
		logger: logger,
	}
}

func (cg *CognitoAuthService) CreateUser(email string, password string) (*cognitoidentityprovider.SignUpOutput, error) {
	cognitoUser, err := cg.client.SignUp(
		context.Background(), &cognitoidentityprovider.SignUpInput{
			ClientId: &cg.env.ClientID,
			Username: &email,
			Password: &password,
			UserAttributes: []types.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(email),
				},
			},
			ValidationData: []types.AttributeType{},
		},
	)

	if err != nil {
		if awsErr := utils.MapAWSError(cg.logger, err); awsErr != nil {
			return nil, awsErr
		}
		return nil, err
	}

	return cognitoUser, nil
}

func (cg *CognitoAuthService) ConfirmSignUp(email string) error {
	_, err := cg.client.AdminConfirmSignUp(
		context.Background(), &cognitoidentityprovider.AdminConfirmSignUpInput{
			UserPoolId: &cg.env.PoolID,
			Username:   &email,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// AdminCreateUser creates normal user by admin
func (cg *CognitoAuthService) AdminCreateUser(id, email, contact_number, password *string) (
	*cognitoidentityprovider.AdminCreateUserOutput, error,
) {
	var attribute []types.AttributeType

	if email != nil {
		att := []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(*email),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		}
		attribute = append(attribute, att...)
	}

	if contact_number != nil {
		att := []types.AttributeType{
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(*contact_number),
			},
			{
				Name:  aws.String("phone_number_verified"),
				Value: aws.String("true"),
			},
		}
		attribute = append(attribute, att...)
	}

	user, err := cg.client.AdminCreateUser(
		context.Background(), &cognitoidentityprovider.AdminCreateUserInput{
			UserPoolId:        &cg.env.PoolID,
			Username:          id,
			TemporaryPassword: password,
			MessageAction:     types.MessageActionTypeSuppress,
			UserAttributes:    attribute,
		},
	)

	if err != nil {
		return nil, err
	}

	_, err = cg.client.AdminSetUserPassword(
		context.Background(), &cognitoidentityprovider.AdminSetUserPasswordInput{
			Username:   id,
			Password:   password,
			Permanent:  true,
			UserPoolId: &cg.env.PoolID,
		},
	)

	if err != nil {
		cg.DeleteUser(*id)
		return nil, err
	}

	return user, nil
}

func (cg *CognitoAuthService) SetEmployeeRoleToUser(id *string, emp_profile string) error {
	err := cg.SetCustomClaimToOneUser(*id, map[string]string{"employee_profile": emp_profile})
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) SetUserRoleToUser(id *string, user_profile string) error {
	err := cg.SetCustomClaimToOneUser(*id, map[string]string{"user_profile": user_profile})
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) SetRoleToUser(id *string, roles models.UserRole) error {
	err := cg.SetCustomClaimToOneUser(*id, map[string]string{"role": strings.Join(roles[:], ",")})
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) SetStoreEmployeeRoleToUser() {

}

func (cg *CognitoAuthService) CreateStoreEmployee() {

}

func (cg *CognitoAuthService) CreateAdminUser(email, password string) error {
	_, err := cg.client.SignUp(
		context.Background(), &cognitoidentityprovider.SignUpInput{
			ClientId: &cg.env.ClientID,
			Password: &password,
			Username: &email,
			UserAttributes: []types.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(email),
				},
			},
			ValidationData: []types.AttributeType{},
		},
	)
	if err != nil {
		return err
	}

	err = cg.SetCustomClaimToOneUser(
		email, map[string]string{
			"role": "admin",
		},
	)
	if err != nil {
		return err
	}

	err = cg.ConfirmSignUp(email)
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) UpdateUserAttribute(user string, c map[string]string) error {
	var claim []types.AttributeType

	for key, val := range c {
		temp := types.AttributeType{
			Name:  aws.String(key),
			Value: aws.String(val),
		}
		claim = append(claim, temp)
	}

	_, err := cg.client.AdminUpdateUserAttributes(
		context.Background(),
		&cognitoidentityprovider.AdminUpdateUserAttributesInput{
			UserAttributes: claim, Username: &user, UserPoolId: &cg.env.PoolID,
		},
	)

	return err
}

func (cg *CognitoAuthService) GetUserByEmail(email string) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	user, err := cg.client.AdminGetUser(
		context.Background(), &cognitoidentityprovider.AdminGetUserInput{
			Username:   &email,
			UserPoolId: &cg.env.PoolID,
		},
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (cg *CognitoAuthService) VerifyUser(email string) error {
	_, err := cg.client.AdminConfirmSignUp(
		context.Background(), &cognitoidentityprovider.AdminConfirmSignUpInput{
			UserPoolId: &cg.env.PoolID,
			Username:   &email,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) VerifyToken(tokenString string) (jwt.Token, error) {
	parsedToken, err := jwt.Parse(
		[]byte(tokenString),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
		jwt.WithIssuer(issuer),
	)

	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func (cg *CognitoAuthService) DeleteUser(user string) error {
	_, err := cg.client.AdminDeleteUser(
		context.Background(), &cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: &cg.env.PoolID,
			Username:   &user,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) DisableUser(user string) error {
	_, err := cg.client.AdminDisableUser(
		context.Background(), &cognitoidentityprovider.AdminDisableUserInput{
			UserPoolId: &cg.env.PoolID,
			Username:   &user,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (cg *CognitoAuthService) EnableUser(user string) error {
	_, err := cg.client.AdminEnableUser(
		context.Background(), &cognitoidentityprovider.AdminEnableUserInput{
			UserPoolId: &cg.env.PoolID,
			Username:   &user,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (cg *CognitoAuthService) ListUsers() (*cognitoidentityprovider.ListUsersOutput, error) {
	users, err := cg.client.ListUsers(
		context.Background(), &cognitoidentityprovider.ListUsersInput{
			UserPoolId: &cg.env.PoolID,
		},
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (cg *CognitoAuthService) PasswordReset(user string) error {
	_, err := cg.client.AdminResetUserPassword(
		context.Background(), &cognitoidentityprovider.AdminResetUserPasswordInput{
			UserPoolId: &cg.env.PoolID,
			Username:   &user,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (cg *CognitoAuthService) SetCustomClaimToOneUser(user string, c map[string]string) error {
	var claim []types.AttributeType
	var create []types.SchemaAttributeType

	for key, val := range c {
		temp := types.AttributeType{
			Name:  aws.String("custom:" + key),
			Value: aws.String(val),
		}
		temp2 := types.SchemaAttributeType{
			AttributeDataType: "String",
			// DeveloperOnlyAttribute: false,
			// Mutable:                true,
			Name: aws.String(key),
			// Required:               false,
		}
		claim = append(claim, temp)
		create = append(create, temp2)
	}

	_, _ = cg.client.AddCustomAttributes(
		context.Background(), &cognitoidentityprovider.AddCustomAttributesInput{
			CustomAttributes: create,
			UserPoolId:       &cg.env.PoolID,
		},
	)

	_, err := cg.client.AdminUpdateUserAttributes(
		context.Background(), &cognitoidentityprovider.AdminUpdateUserAttributesInput{
			UserAttributes: claim,
			UserPoolId:     &cg.env.PoolID,
			Username:       &user,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
