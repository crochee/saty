package v

import "net/textproto"

var (
	HeaderAccountID     = textproto.CanonicalMIMEHeaderKey("X-Account-ID")
	HeaderUserID        = textproto.CanonicalMIMEHeaderKey("X-User-ID")
	HeaderRealAccountID = textproto.CanonicalMIMEHeaderKey("X-Real-Account-ID")
	HeaderRealUserID    = textproto.CanonicalMIMEHeaderKey("X-Real-User-ID")
	HeaderSource        = textproto.CanonicalMIMEHeaderKey("X-Source")
	HeaderAccountName   = textproto.CanonicalMIMEHeaderKey("X-Account-Name")
	HeaderRealIP        = textproto.CanonicalMIMEHeaderKey("X-Real-IP")

	HeaderTraceID = textproto.CanonicalMIMEHeaderKey("X-Trace-ID")

	HeaderCacheControl = textproto.CanonicalMIMEHeaderKey("Cache-Control")

	// 网关解析的请求头
	HeaderGWAccountID = textproto.CanonicalMIMEHeaderKey("Accountid")
	HeaderGWUserID    = textproto.CanonicalMIMEHeaderKey("Userid")
	HeaderGWToken     = textproto.CanonicalMIMEHeaderKey("Token")
)

const (
	HeaderAdminID      = "X-Admin-ID"
	HeaderOperatorID   = "operatorId"
	HeaderOperatorName = "innerLoginUserName"
	HeaderOperatorType = "operateUserType"

	HeaderStaff       = "staff"
	HeaderForwardType = "forward-type"

	HeaderAuthorization = "Authorization"
	HeaderAPIFrom       = "X-API-From"
	HeaderAccessToken   = "ACCESS_TOKEN"

	// Header from caller
	HeaderCallerID    = "X-Caller-ID"
	HeaderCallerCode  = "X-Caller-Code"
	HeaderCallerUser  = "X-Caller-User"
	HeaderCallerExtra = "X-Caller-Extra"
)

const (
	// X-Source请求头允许传值的列表
	HeaderXSourceValueMsp     = "msp"
	HeaderXSourceValueCsk     = "csk"
	HeaderXSourceValueDsk     = "dsk"
	HeaderXSourceValueNas     = "nas"
	HeaderXSourceValueCeen    = "ceen"
	HeaderXSourceValuePaas    = "paas"
	HeaderXSourceValueConsole = "console"
)

// Authorization Basic
const (
	AuthorizationBasic = "Basic"
)
