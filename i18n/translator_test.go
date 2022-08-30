package i18n

import "testing"

func TestLoad(t *testing.T) {
	if err := Load("../locale"); err != nil {
		t.Error(err)
	}
	t.Logf("Files %+v\n", Default.files)
	t.Logf("Translations %+v\n", Default.translations)

	v, ok := Default.translations.languages["en"]
	if !ok {
		t.Fail()
	}

	if len(v.categories) == 0 {
		t.Fail()
	}
}
