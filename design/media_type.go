package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BasicInfoMedia : サーバ基本情報を返却するMT
var BasicInfoMedia = MediaType("application/vnd.result+json", func() {
	Description("MT for BasicInfo.")
	// どのような値があるか（複数定義出来る）
	Attributes(func() {
		Attribute("status")
		Attribute("count", Integer)
		Attribute("basicinfo", BasicInfoType)
	})
	// 返すレスポンスのフォーマット（別記事で紹介予定）
	View("default", func() {
		Attribute("basicinfo", BasicInfoType)
	})
})

// BasicInfoType : サーバ基本情報を含むType
var BasicInfoType = Type("BasicInfoType", func() {
	Attribute("hostname", String)
	Attribute("status", String)
	Attribute("role", ArrayOf(String))
	Attribute("hosttype", String)
	Attribute("os")
	Attribute("kernel")
	Attribute("tag", ArrayOf(String))
	Attribute("monitorstatus", Boolean)
})
