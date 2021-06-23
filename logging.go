package hexlink

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/nevzheng/hexlink/shortener"
	t "github.com/nevzheng/hexlink/types"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   shortener.RedirectService
}

func (mw LoggingMiddleware) Find(code string) (output *t.Redirect, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Find",
			"input", code,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)

	}(time.Now())
	output, err = mw.Next.Find(code)
	return
}

func (mw LoggingMiddleware) Store(redirect *t.Redirect) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Store",
			"url", redirect.Url,
			"timeCreated", redirect.TimeCreated,
			"code", redirect.RedirectCode,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)

	}(time.Now())
	output, err = mw.Next.Store(redirect)
	return
}
