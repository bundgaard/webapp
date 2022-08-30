package i18n

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type I18N struct {
	language     string
	files        []string
	translations *Language
}

var Default = &I18N{
	translations: NewLanguage(),
}

func T(key string) string {
	parts := strings.Split(key, ".")
	category, item := parts[0], parts[1]
	language, ok := Default.translations.GetLanguage(Default.language)
	if !ok {
		return ""
	}
	group, ok := language.GetCategory(category)
	if !ok {
		return ""
	}
	part, ok := group.GetItem(item)
	if !ok {
		return ""
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

	// Category
	for _, file := range Default.files {
		language, category := splitLanguageAndCategory(file)

		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var translations map[string]map[string]string

		if err := yaml.NewDecoder(bytes.NewReader(content)).Decode(&translations); err != nil {
			log.Fatal(err)
		}

		item := Item{items: make(map[string]string)}
		// Category
		for k, v := range translations[language] { // app_title: foo
			item.items[k] = v
		}
		add(language, category, item)

	}

	return nil
}

func add(language string, category string, item Item) {
	_, ok := Default.translations.languages[language].categories[category]
	if !ok {
		Default.translations.languages[language] = Category{categories: map[string]Item{category: item}}
	} else {
		Default.translations.languages[language].categories[category] = item
	}
}

func splitLanguageAndCategory(file string) (string, string) {
	lastSlash := strings.LastIndex(file, "\\")
	file = file[lastSlash+1:]
	parts := strings.Split(file, ".")
	category, language, _ := parts[0], parts[1], parts[2]
	slashIdx := strings.LastIndex(category, "\\")
	category = category[slashIdx+1:]
	return language, category
}
