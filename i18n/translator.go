package i18n

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Language string
type Type string
type Item string

type I18N struct {
	language     string
	files        []string
	translations map[Language]map[Type]map[Item]string
}

/*
	[language] -> [type]  -> [item]
*/

var Default = &I18N{
	translations: make(map[Language]map[Type]map[Item]string),
}

func Translate(key string) string {
	parts := strings.Split(key, ".")
	category, item := parts[0], parts[1]
	return Default.translations[Language(Default.language)][Type(category)][Item(item)]
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
			Default.translations[Language(language)] = make(map[Type]map[Item]string)
		}

		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var translations map[string]map[string]string

		if err := yaml.NewDecoder(bytes.NewReader(content)).Decode(&translations); err != nil {
			log.Fatal(err)
		}

		Default.translations[Language(language)][Type(category)] = make(map[Item]string)
		for k, v := range translations[language] {
			Default.translations[Language(language)][Type(category)][Item(k)] = v
		}
	}
	return nil
}
