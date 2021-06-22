package shortener

import (
	t "github.com/nevzheng/hexlink/types"
)

type RedirectService interface {
	Find(code string) (*t.Redirect, error)
	Store(redirect *t.Redirect) (string, error)
}
