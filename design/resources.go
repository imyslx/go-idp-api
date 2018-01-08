package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Hostを管理するリソース
var _ = Resource("hosts", func() {
	BasePath("/hosts")

	// 基礎情報を含んだHost一覧の取得
	Action("list", func() {
		Description("Hostの基本情報の取得")
		Routing(
			// Endpoint -> http://localhost:8080/v1/hosts
			GET(""),
		)

		Payload(HostsPayload)

		Response(OK, BasicInfoMedia)
		Response(BadRequest, ErrorMedia)
	})

	// Host一覧の取得
	Action("simplelist", func() {
		Description("Host名のリストを取得")
		Routing(
			// Endpoint -> http://localhost:8080/v1/hosts
			GET("/simple"),
		)

		// パラメータの定義
		Payload(HostsPayload)

		Response(OK, SimpleListMedia)
		Response(BadRequest, ErrorMedia)
	})
})
