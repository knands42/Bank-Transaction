package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	errrosMap := map[bool]error{
		session.IsBlocked: fmt.Errorf("session is blocked"),
		session.Username != refreshPayload.Username: fmt.Errorf("refresh token doesn't match username"),
		session.RefreshToken != req.RefreshToken:    fmt.Errorf("refresh token doesn't match"),
		time.Now().After(session.ExpiresAt):         fmt.Errorf("refresh token is expired"),
	}

	for condition, err := range errrosMap {
		if condition {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
	}

	tokenPayload, accessToken, err := server.tokenMaker.CreateToken(
		req.RefreshToken,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: tokenPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
