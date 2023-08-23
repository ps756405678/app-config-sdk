package domain

type DeleteKeyForm struct {
	AppId string `json:"AppId" form:"AppId" validate:"required" message:"AppId 必须"`
	Type  string `json:"Type" form:"Type" validate:"required|enum:APP_CURRENT_OPT,APP_SERVICE,APP_CURRENT_CONFIG,APP_INFRA" message:"required:Type required|enum:Type invalid"`
	Path  string `json:"Path" form:"Path"`
}

type ConfigGetReq struct {
	AppId  string `json:"AppId" form:"AppId" validate:"required" message:"AppId 必须"`
	Type   string `json:"Type" form:"Type"`
	Fields string `json:"Fields" form:"Fields"`
	Path   string `json:"Path" form:"Path"`
}

type ServiceRegisterReq struct {
	AppId       string                 `json:"AppId" form:"AppId" validate:"required" message:"AppId required"`
	ServiceName string                 `json:"ServiceName" form:"ServiceName" validate:"required" message:"ServiceName required"`
	Data        map[string]interface{} `json:"Data" form:"Data" validate:"required" message:"Data required"`
}

type ConfigSetReq struct {
	AppId string      `json:"AppId" form:"AppId" validate:"required" message:"AppId required"`
	Data  interface{} `json:"Data" form:"Data" validate:"required" message:"Data required"`
	Type  string      `json:"Type" form:"Type" validate:"required|enum:APP_CURRENT_OPT,APP_SERVICE,APP_CURRENT_CONFIG,APP_INFRA" message:"required:Type required|enum:Type invalid"`
}

type CallResp struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}
