package msgutil

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

// Setter is language setter for message's default catalog.
type Setter func(*catalog.Catalog, language.Tag) error

// SetLanguage set languages to catalog.
func SetLanguage(c *catalog.Catalog, t language.Tag, ls []Setter) error {
	for _, set := range ls {
		err := set(c, t)
		if err != nil {
			return err
		}
	}
	return nil
}

// String returns string setter.
func String(key string, value string) Setter {
	return func(c *catalog.Catalog, t language.Tag) error {
		return c.SetString(t, key, value)
	}
}

// Message returns message setter.
func Message(key string, msgs ...catalog.Message) Setter {
	return func(c *catalog.Catalog, t language.Tag) error {
		return c.Set(t, key, msgs...)
	}
}

// Macro returns macro setter.
func Macro(name string, msgs ...catalog.Message) Setter {
	return func(c *catalog.Catalog, t language.Tag) error {
		return c.SetMacro(t, name, msgs...)
	}
}
