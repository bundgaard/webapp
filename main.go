package main

import (
	"fmt"
	"log"
	"webapp/i18n"
)

func main() {

	i18n.SetLanguage("nl")
	if err := i18n.Load("locale"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Translation of Title:", i18n.Translate("general.app_title"))
}
