package msgutil

import (
	"testing"

	"reflect"

	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

var japanese = []Setter{
	String("Hello, %s", "こんにちは、%s"),
	String("Good bye, %s", "さようなら、%s"),
}

func TestSetLanguage(t *testing.T) {
	j := catalog.New()
	j.SetString(language.Japanese, "Hello, %s", "こんにちは、%s")
	j.SetString(language.Japanese, "Good bye, %s", "さようなら、%s")

	cases := []struct {
		setters []Setter
		tag     language.Tag
		want    *catalog.Catalog
	}{
		{
			setters: japanese,
			tag:     language.Japanese,
			want:    j,
		},
	}

	for i, tc := range cases {
		got := catalog.New()
		err := SetLanguage(got, tc.tag, tc.setters)
		if err != nil {
			t.Errorf("#%v SetLanguage() occured error %s", i, err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("#%v SetLanguage(\ngot: %#v, want: %#v", i, got, tc.want)
		}
	}
}
