package services

import (
	"theodo.red/creditcompanion/packages/logging"
	"theodo.red/creditcompanion/packages/teatime"
	"theodo.red/creditcompanion/packages/tokens/models"
)

type RefreshingTokenService struct {
	tokenRefreshService          TokenRefreshService
	tokenRepository              models.TokenRepository
	tokenRefreshThresholdSeconds int
	clock                        teatime.Clock
	logger                       logging.Logger
}

func (r *RefreshingTokenService) GetTokenById(id string) (*models.Token, error) {
	token, err := r.tokenRepository.Get(id)
	if err != nil {
		return nil, err
	}

	if r.tokenIsCloseToOrHasExpired(token) {
		refreshedToken, refreshErr := r.tokenRefreshService.RefreshToken(token)
		if refreshErr != nil {
			if r.tokenIsActive(token) {
				r.logger.LogDebug("Token is near to expiry yet refresh failed. Continuing anyway, the request may fail.", token.Id)
			} else {
				return nil, refreshErr
			}
		} else {
			setErr := r.tokenRepository.Set(token.Id, refreshedToken)
			if setErr != nil {
				// TODO: Turns out returning a value and an error is bad practice. Consider changing in the future.
				return refreshedToken, setErr
			}
			token = refreshedToken
		}
	}

	return token, nil
}

func (r *RefreshingTokenService) tokenIsActive(token *models.Token) bool {
	return r.clock.Now().Before(token.ExpiresAfterTime())
}

func (r *RefreshingTokenService) tokenIsCloseToOrHasExpired(token *models.Token) bool {
	if !r.tokenIsActive(token) {
		return true
	}

	return (token.ExpiresAfterTime().Unix() - r.clock.Now().Unix()) < int64(r.tokenRefreshThresholdSeconds)
}
