package msgutil

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

// Expr type is language expression that set message to catalog.
type Expr func(*catalog.Catalog, language.Tag) error

// SetLanguage set languages to catalog.
func SetLanguage(c *catalog.Catalog, t language.Tag, ls []Expr) error {
	for _, set := range ls {
		err := set(c, t)
		if err != nil {
			return err
		}
	}
	return nil
}

// String returns string type message expression setter.
func String(key string, value string) Expr {
	return func(c *catalog.Catalog, t language.Tag) error {
		return c.SetString(t, key, value)
	}
}

// Message returns message expression setter.
func Message(key string, msgs ...catalog.Message) Expr {
	return func(c *catalog.Catalog, t language.Tag) error {
		return c.Set(t, key, msgs...)
	}
}

// Macro returns macro expression setter.
func Macro(name string, msgs ...catalog.Message) Expr {
	return func(c *catalog.Catalog, t language.Tag) error {
		return c.SetMacro(t, name, msgs...)
	}
}
