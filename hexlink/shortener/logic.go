package shortener

import (
	"errors"
	"time"

	"github.com/go-kit/kit/log"

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
	logger             log.Logger
}

func NewRedirectService(repository RedirectRepository, logger log.Logger) RedirectService {
	return &redirectService{
		redirectRepository: repository,
		logger:             logger,
	}
}
func (r redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepository.Find(code)
}

func (r redirectService) Store(redirect *Redirect) (string, error) {
	if err := validate.Validate(redirect); err != nil {
		return "", errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate() // TBD: Change up the strategy
	redirect.CreatedAt = time.Now().UTC().Unix()
	if err := r.redirectRepository.Store(redirect); err != nil {
		return "", err
	} else {
		return redirect.Code, err
	}
}
