package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Hostを管理するリソース
var _ = Resource("hostdetail", func() {
	BasePath("/host")

	// 基礎情報を含んだHost一覧の取得
	Action("hostdetial", func() {
		Description("Hostの基本情報の取得")
		Routing(
			// Endpoint -> http://localhost:8080/v1/hosts
			GET("/:hostname"),
		)

		Response(OK, DetailInfoMedia)
		Response(BadRequest, ErrorMedia)
	})

})
