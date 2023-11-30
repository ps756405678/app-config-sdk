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

type SortValue struct {
	Sort  string `json:"sort" form:"sort"`
	Field string `json:"field" from:"field"`
}
type ConfigListReq struct {
	AppId    string         `json:"AppId" form:"AppId" validate:"required" message:"AppId 必须"`
	Fileter  map[string]any `json:"Fileter" form:"Fileter"`
	Sort     []SortValue    `json:"Sort" form:"Sort"`
	Page     int64          `json:"Page" form:"Page" default:"1"`
	PageSize int64          `json:"PageSize" form:"PageSize" default:"20"`
}

type ServiceRegisterReq struct {
	AppId       string                 `json:"AppId" form:"AppId" validate:"required" message:"AppId required"`
	ServiceName string                 `json:"ServiceName" form:"ServiceName" validate:"required" message:"ServiceName required"`
	Data        map[string]interface{} `json:"Data" form:"Data" validate:"required" message:"Data required"`
}

type ConfigSetReq struct {
	AppId string      `json:"AppId" form:"AppId" validate:"required" message:"AppId required"`
	Data  interface{} `json:"Data" form:"Data" validate:"required" message:"Data required"`
	Path  string      `json:"Path" form:"Path"`
	Type  string      `json:"Type" form:"Type" validate:"required|enum:APP_CURRENT_OPT,APP_SERVICE,APP_CURRENT_CONFIG,APP_INFRA" message:"required:Type required|enum:Type invalid"`
}

type CallResp[T any] struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
	Data    T      `json:"data"`
}

type GetListResp[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}
