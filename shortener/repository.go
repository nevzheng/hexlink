package shortener

import t "github.com/nevzheng/hexlink/types"

type RedirectRepository interface {
	Find(code string) (*t.Redirect, error)
	Store(redirect *t.Redirect) error
}
