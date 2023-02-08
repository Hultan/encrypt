package encrypt

import (
	"errors"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type GtkBuilder struct {
	Builder *gtk.Builder
}

// GetObject : Gets a gtk object by name
func (g *GtkBuilder) GetObject(name string) glib.IObject {
	if g.Builder == nil {
		panic(errors.New("need to call CreateBuilder first"))
	}
	obj, err := g.Builder.GetObject(name)
	if err != nil {
		panic(err)
	}

	return obj
}
