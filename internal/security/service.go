package security

import (
	"github.com/nmarsollier/commongo/errs"
	"github.com/nmarsollier/commongo/log"
)

type SecurityService interface {
	Validate(token string) (*User, error)
	Invalidate(token string)
}

func NewSecurityService(
	log log.LogRusEntry,
	repo SecurityRepository,
) SecurityService {
	return &securityService{
		log:  log,
		repo: repo,
	}
}

type securityService struct {
	log  log.LogRusEntry
	repo SecurityRepository
}

// Invalidate invalidates a token from the cache
func (s *securityService) Invalidate(token string) {
	if len(token) <= 7 {
		s.log.Info("Token no valido: ", token)
		return
	}

	s.repo.CleanToken(token)
	s.log.Info("Token invalidado: ", token)
}

// Validate checks if the token is valid
func (s *securityService) Validate(token string) (*User, error) {
	// If it is in cache, return the cache
	if user, ok := s.repo.GetToken(token); ok {
		return user, nil
	}

	user, err := s.repo.GetRemoteToken(token)
	if err != nil {
		return nil, errs.Unauthorized
	}

	return user, nil
}
