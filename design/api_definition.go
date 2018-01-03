package design

import (
	"os"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("imyslx/go-idp-api", func() {
	Title("imyslx/go-idp-api")
	Description("Infraの構成情報などをBackendのCouchbaseを介してやり取りするAPIです。")

	Contact(func() {
		Name("imyslx")
		Email("imyslx29@gmail.com")
		URL("https://github.com/imyslx/go-idp-api/issues")
	})
	License(func() {
		Name("MIT")
		URL("")
	})
	Docs(func() {
		Description("wiki")
		URL("https://github.com/imyslx/go-idp-api/README.md")
	})
	Host(func() string {
		switch os.Getenv("Op") {
		case "develop":
			return "localhost:8080"
		case "staging":
			return "staging.com"
		case "production":
			return "production.com"
		}
		return "localhost:8080"
	}())
	Scheme(func() string {
		switch os.Getenv("Op") {
		case "develop":
			return "http"
		case "staging":
			return "https"
		case "production":
			return "https"
		}
		return "http"
	}())
	BasePath("/v1")
	//　CORSポリシーの定義
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE")
	})

})
