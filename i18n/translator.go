package i18n

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Language struct {
	languages map[string]*Category
}

func (l *Language) GetLanguage(language string) (Category, bool) {
	v, ok := l.languages[language]
	return v, ok
}

func (l *Language) AddLanguage(language string) error {
	_, ok := l.languages[language]
	if !ok {
		l.languages[language] = &Category{}
	}
	return nil
}

type Category struct {
	categories map[string]Item
}

func (c *Category) GetCategory(category string) (Item, bool) {
	v, ok := c.categories[category]
	return v, ok
}

func (c *Category) merge(other Category) {
	for k, v := range other.categories {
		c.categories[k] = v
	}
}

type Item struct {
	items map[string]string
}

func (i Item) GetItem(item string) (string, bool) {
	v, ok := i.items[item]
	return v, ok
}

type I18N struct {
	language     string
	files        []string
	translations *Language
}

/*
	[language] -> [type]  -> [item] -> Translated Value
*/

var Default = &I18N{
	translations: &Language{languages: make(map[string]Category)},
}

func T(key string) string {
	parts := strings.Split(key, ".")
	category, item := parts[0], parts[1]
	language, ok := Default.translations.GetLanguage(Default.language)
	if !ok {
	}
	group, ok := language.GetCategory(category)
	if !ok {

	}
	part, ok := group.GetItem(item)
	if !ok {

	}
	return part
}

func SetLanguage(language string) {
	Default.language = language
}

func Load(relFp string) error {

	entries, err := os.ReadDir(relFp)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".yml") || strings.HasSuffix(entry.Name(), ".yaml") {
			Default.files = append(Default.files, filepath.Join(relFp, entry.Name()))
		}
	}

	for _, file := range Default.files {
		parts := strings.Split(file, ".")
		category, language, _ := parts[0], parts[1], parts[2]
		slashIdx := strings.LastIndex(category, "\\")
		category = category[slashIdx+1:]
		_, ok := Default.translations[Language(language)]
		if !ok {
			Default.translations[Language(language)] = make(map[Category]map[Item]string)
		}

		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var translations map[string]map[string]string

		if err := yaml.NewDecoder(bytes.NewReader(content)).Decode(&translations); err != nil {
			log.Fatal(err)
		}

		Default.translations[Language(language)][Category(category)] = make(map[Item]string)
		for k, v := range translations[language] {
			Default.translations[Language(language)][Category(category)][Item(k)] = v
		}
	}
	return nil
}
