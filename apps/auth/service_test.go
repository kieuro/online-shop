package auth

import (
	"context"
	"fmt"
	"kieuro-online-shop/external/database"
	"kieuro-online-shop/infra/response"
	"kieuro-online-shop/internal/config"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	email := fmt.Sprintf("%v@example.com", uuid.NewString())
	req := RegisterRequestPayload{
		Email:    email,
		Password: "secretpassword",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)

}
func TestRegister_Fail(t *testing.T) {
	t.Run("email already used", func(t *testing.T) {
		// preparation for duplicate email
		email := fmt.Sprintf("%v@example.com", uuid.NewString())
		req := RegisterRequestPayload{
			Email:    email,
			Password: "secretpassword",
		}
		err := svc.register(context.Background(), req)
		require.Nil(t, err)

		// end preparation

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})

}

func TestLoginSuccess(t *testing.T) {
	email := fmt.Sprintf("%v@example.com", uuid.NewString())
	pass := "secretpassword"
	req := RegisterRequestPayload{
		Email:    email,
		Password: pass,
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email:    email,
		Password: pass,
	}

	token, err := svc.login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)

	log.Println(token)

}
