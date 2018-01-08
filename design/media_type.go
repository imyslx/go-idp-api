package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BasicInfoMedia : サーバ基本情報を返却するMT
var BasicInfoMedia = MediaType("application/vnd.bi.result+json", func() {
	Description("MT for BasicInfo.")
	Attributes(func() {
		Attribute("ResponseStatus")
		Attribute("Count", Integer)
		Attribute("BasicInfo", ArrayOf(BasicInfoType))
	})
	View("default", func() {
		Attribute("ResponseStatus")
		Attribute("Count", Integer)
		Attribute("BasicInfo", ArrayOf(BasicInfoType))
	})
})

// SimpleListMedia : サーバ名のArrayを返却するMT
var SimpleListMedia = MediaType("application/vnd.sl.result+json", func() {
	Description("MT for SimpleList.")
	Attributes(func() {
		Attribute("Hostname", ArrayOf(String), func() {
			Example([]string{"commydb01", "commydb02", "vmdb08.local"})
		})
	})
	View("default", func() {
		Attribute("Hostname", ArrayOf(String), func() {
			Example([]string{"commydb01", "commydb02", "vmdb08.local"})
		})
	})
})

// BasicInfoListType : サーバ基本情報をResponseするためのType
var BasicInfoListType = Type("BasicInfoListType", func() {
	Attribute("ResponseStatus", func() {
		Example("success")
	})
	Attribute("Count", Integer, func() {
		Example(1)
	})
	Attribute("BasicInfo", ArrayOf(BasicInfoType))
})

// BasicInfoType : サーバ基本情報を含むType
var BasicInfoType = Type("BasicInfoType", func() {
	Attribute("Hostname", String, func() {
		Example("commydb01")
	})
	Attribute("Status", Boolean, func() {
		Example(true)
	})
	Attribute("Role", ArrayOf(String), func() {
		Example([]string{"Web", "DB", "Cache"})
	})
	Attribute("Type", String, func() {
		Example("VM_Guest")
	})
	Attribute("OperatingSystem", func() {
		Example("CentOS Linux release 7.4.1708 (Core)")
	})
	Attribute("Kernel", func() {
		Example("Linux commydb02.local-i.style 3.10.0-693.el7.x86_64 #1 SMP Tue Aug 22 21:09:27 UTC 2017 x86_64")
	})
	Attribute("Tag", ArrayOf(String), func() {
		Example([]string{"@PC", "@SP", "biche"})
	})
	Attribute("monitoringStatus", Boolean, func() {
		Example(true)
	})
})

// HostsPayload : ホスト一覧用のPayload
var HostsPayload = Type("HostsPayload", func() {
	Member("hostname", String, "ホスト名", func() {
		Example("host-a.local")
	})
	Member("status", Boolean, "ステータス", func() {
		Example(true)
	})
	Member("role", ArrayOf(String), "役割", func() {
		Example([]string{"Web", "DB", "Cache"})
	})
	Member("type", String, "サーバタイプ", func() {
		Example("VM_Guest")
	})
	Member("os", String, "OperatingSystem", func() {
		Example("CentOS Linux release 7.4.1708 (Core)")
	})
	Member("ip", String, "IPセグメント", func() {
		Example("10.1.1.1")
	})
	Member("tag", ArrayOf(String), "サーバに付加されたタグ", func() {
		Example([]string{"@PC", "@SP", "biche"})
	})
})
