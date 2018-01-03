package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Hostを管理するリソース
var _ = Resource("hosts", func() {
	BasePath("/hosts")

	// Host一覧の取得
	Action("list", func() {
		Description("Host一覧に必要な情報の取得")
		Routing(
			// Endpoint -> http://localhost:8080/v1/hosts
			GET(""),
		)

		// パラメータの定義
		Payload(func() {
			Member("hostname", String, "ホスト名")
			Member("status", Boolean, "ステータス")
			Member("role", ArrayOf(String), "役割")
			Member("type", ArrayOf(String), "サーバタイプ")
			Member("os", String, "OperatingSystem")
			Member("ip", String, "IPセグメント")
			Member("tag", ArrayOf(String), "サーバに付加されたタグ")
		})

		Response(OK, ArrayOf(BasicInfoMedia))
		Response(BadRequest, ErrorMedia)
	})
})
