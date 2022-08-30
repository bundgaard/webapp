package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"webapp/html"
	"webapp/i18n"
)

type Foo struct{}

func (f *Foo) Render(w io.Writer) error {
	return nil
}
func Bar(w http.ResponseWriter) {
	one := html.Div(html.WithID("first-child"), html.WithText("Hello, World"))
	two := html.Div(html.WithID("second-child"), html.WithText("Second, Hello World!"))
	three := html.Div(html.WithID("third-child"), html.WithText("Third, Hello World"))
	elm := html.Div(html.WithID("my-id"), html.WithChildren(one, two, three))
	if err := elm.Render(w); err != nil {
		log.Fatal(err)
	}

}
func main() {

	if err := i18n.Load("locale"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/nl", func(w http.ResponseWriter, r *http.Request) {
		i18n.SetLanguage("nl")
		title := i18n.T("general.app_title")
		fmt.Println("Translation of Title:", title)
		Bar(w)
		fmt.Fprintf(w, "Title %s", title)
	})

	http.HandleFunc("/en", func(w http.ResponseWriter, r *http.Request) {
		language := r.Header.Get("Accept-Language")
		if strings.Contains(language, "en") {
			i18n.SetLanguage("en")
		}
		Bar(w)
		title := i18n.T("general.app_title")
		fmt.Println("Translation of Title:", title)
		fmt.Fprintf(w, "Title %s", title)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
