package authServices

import (
	"fmt"

	authGRPC "github.com/j-ew-s/ms-curso-ordering-api/auth-services/grpc"
	"github.com/valyala/fasthttp"

	sessionModels "github.com/j-ew-s/ms-curso-ordering-api/auth-services/models"
	common "github.com/j-ew-s/ms-curso-ordering-api/common"
)

var GlobalSession *sessionModels.Session

func AuthSessionValidator(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))

		if token != "" {
			GlobalSession = &sessionModels.Session{Token: token}

			err := Authorization()
			if err != nil {
				common.PrepareResponse(ctx, 500, "Error Token Validation")
				return
			}

			if GlobalSession.TokenValidation.Status {
				requestHandler(ctx)
			} else {
				common.PrepareResponse(ctx, 401, "Token não válido")
			}

		} else {
			common.PrepareResponse(ctx, 401, "Não tem acesso ao conteúdo")
		}
	}
}

func Authorization() error {
	fmt.Println(fmt.Sprintf("Token a ser buscado: %s", GlobalSession.Token))

	resp, err := authGRPC.IsTokenValid(GlobalSession.Token)
	if err != nil {
		erro := fmt.Errorf(fmt.Sprintf("gRPC DIAL falhou: %v", err))
		return erro
	}

	GlobalSession.TokenValidation = resp

	return nil
}

func GetUserInfo() error {
	resp, err := authGRPC.GetUserInfo(GlobalSession.Token)
	if err != nil {
		erro := fmt.Errorf(fmt.Sprintf("gRPC DIAL falhou: %v", err))
		return erro
	}

	GlobalSession.Email = resp.Email
	GlobalSession.Id = resp.Id
	GlobalSession.Username = resp.Username
	GlobalSession.Name = resp.Name

	return nil
}
