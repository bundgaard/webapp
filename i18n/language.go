package i18n

type Language struct {
	languages map[string]Category
}

func (l *Language) GetLanguage(language string) (Category, bool) {
	v, ok := l.languages[language]
	return v, ok
}

func NewLanguage() *Language {
	return &Language{languages: map[string]Category{}}
}
