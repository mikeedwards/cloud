package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var AddFirmwarePayload = Type("AddFirmwarePayload", func() {
	Attribute("etag")
	Required("etag")
	Attribute("module")
	Required("module")
	Attribute("profile")
	Required("profile")
	Attribute("url")
	Required("url")
	Attribute("meta")
	Required("meta")
})

var UpdateDeviceFirmwarePayload = Type("UpdateDeviceFirmwarePayload", func() {
	Attribute("deviceId", Integer)
	Required("deviceId")
	Attribute("firmwareId", Integer)
	Required("firmwareId")
})

var FirmwareSummary = MediaType("application/vnd.app.firmware+json", func() {
	TypeName("FirmwareSummary")
	Attributes(func() {
		Attribute("id", Integer)
		Attribute("time", DateTime)
		Attribute("etag", String)
		Attribute("module", String)
		Attribute("profile", String)
		Attribute("url", String)
		Attribute("meta", HashOf(String, Any))
		Attribute("buildTime", Integer)
		Attribute("buildNumber", Integer)
		Required("id")
		Required("time")
		Required("etag")
		Required("module")
		Required("profile")
		Required("url")
		Required("meta")
		Required("buildTime")
		Required("buildNumber")
	})
	View("default", func() {
		Attribute("id")
		Attribute("time")
		Attribute("etag")
		Attribute("module")
		Attribute("profile")
		Attribute("url")
		Attribute("meta")
		Attribute("buildTime")
		Attribute("buildNumber")
	})
})

var Firmwares = MediaType("application/vnd.app.firmwares+json", func() {
	TypeName("Firmwares")
	Attributes(func() {
		Attribute("firmwares", CollectionOf(FirmwareSummary))
		Required("firmwares")
	})
	View("default", func() {
		Attribute("firmwares")
	})
})

var _ = Resource("Firmware", func() {
	Action("download", func() {
		Routing(GET("firmware/:firmwareId/download"))
		Params(func() {
			Param("firmwareId", Integer)
			Required("firmwareId")
		})
		Response(NotFound)
		Response(OK, func() {
			Status(200)
			Headers(func() {
				Header("ETag")
			})
		})
	})

	Action("add", func() {
		Security(JWT, func() {
			Scope("api:admin")
		})
		Routing(PATCH("firmware"))
		Description("Add firmware")
		Params(func() {
		})
		Payload(AddFirmwarePayload)
		Response(BadRequest)
		Response(OK)
	})

	Action("list", func() {
		Routing(GET("firmware"))
		Description("List firmware")
		Params(func() {
			Param("module", String)
			Param("profile", String)
			Param("pageSize", Integer)
			Param("page", Integer)
		})
		Response(OK, func() {
			Media(Firmwares)
		})
	})

	Action("delete", func() {
		Security(JWT, func() {
			Scope("api:admin")
		})
		Routing(DELETE("firmware/:firmwareId"))
		Description("Delete firmware")
		Params(func() {
			Param("firmwareId", Integer)
			Required("firmwareId")
		})
		Response(NotFound)
		Response(OK, func() {
			Status(200)
		})
	})
})
