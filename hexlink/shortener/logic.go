package shortener

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"

	errs "github.com/pkg/errors"
	validate "gopkg.in/validator.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepository RedirectRepository
}

func NewRedirectService(repository RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepository: repository,
	}
}
func (r redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepository.Find(code)
}

func (r redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate() // TBD: Change up the strategy
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepository.Store(redirect)
}
