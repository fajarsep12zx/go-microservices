package service

import (
	"context"

	coreclient "zebrax.id/product/dmaa/core/client"
	servicemodel "zebrax.id/product/dmaa/core/proto"
)

// GetValidateTokenSession ...
func GetValidateTokenSession(token string) (*servicemodel.AuthData, error) {
	service := coreclient.GetAuthService()
	response, err := service.ValidateAccessTokenSession(
		context.Background(),
		&servicemodel.AuthRequest{
			AccessToken: token,
		},
	)

	if err != nil {
		return nil, err
	}

	return response.AuthToken, nil
}
