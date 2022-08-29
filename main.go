package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/i18n"
)

func main() {

	if err := i18n.Load("locale"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/nl", func(w http.ResponseWriter, r *http.Request) {
		i18n.SetLanguage("nl")
		fmt.Println("Translation of Title:", i18n.Translate("general.app_title"))
		fmt.Fprintf(w, "Title %s", i18n.Translate("general.app_title"))
	})
	http.HandleFunc("/en", func(w http.ResponseWriter, r *http.Request) {
		i18n.SetLanguage("en")
		fmt.Println("Translation of Title:", i18n.Translate("general.app_title"))
		fmt.Fprintf(w, "Title %s", i18n.Translate("general.app_title"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
