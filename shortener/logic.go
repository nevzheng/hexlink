package shortener

import (
	"errors"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/teris-io/shortid"

	errs "github.com/pkg/errors"
	validate "gopkg.in/validator.v2"

	t "github.com/nevzheng/hexlink/types"
)

var (
	ErrRedirectNotFound = errors.New("redirect Not Found")
	ErrRedirectInvalid  = errors.New("redirect Invalid")
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
func (r redirectService) Find(code string) (*t.Redirect, error) {
	return r.redirectRepository.Find(code)
}

func (r redirectService) Store(redirect *t.Redirect) (string, error) {
	if err := validate.Validate(redirect); err != nil {
		return "", errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.RedirectCode = t.Code(shortid.MustGenerate()) // TBD: Change up the strategy
	redirect.TimeCreated = time.Now().UTC()
	if err := r.redirectRepository.Store(redirect); err != nil {
		return "", err
	} else {
		return string(redirect.RedirectCode), err
	}
}
