package entry

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/ps756405678/app-config-sdk/domain"
)

const (
	BaseUrl      = "http://192.168.0.60:8086/api"
	RegisterUrl  = BaseUrl + "/service/register"
	SetConfigUrl = BaseUrl + "/setConfig"
	GetConfigUrl = BaseUrl + "/getConfig"
)

func ServiceRegister(req domain.ServiceRegisterReq) (resp domain.CallResp, err error) {
	return callAppConfigService(http.MethodPost, RegisterUrl, req)
}

func SetConfig(req domain.ConfigSetReq) (resp domain.CallResp, err error) {
	return callAppConfigService(http.MethodPost, SetConfigUrl, req)
}

func GetConfig(req domain.ConfigGetReq) (resp domain.CallResp, err error) {
	formData := url.Values{}
	formData.Set("AppId", req.AppId)
	formData.Set("Type", req.Type)
	formData.Set("Path", req.Path)
	formData.Set("Fields", req.Fields)

	return callAppConfigService(http.MethodGet, GetConfigUrl+"?"+formData.Encode(), req)
}

// 调用SDK sevice
func callAppConfigService(method string, url string, req any) (resp domain.CallResp, err error) {
	// 序列化参数
	bData, _ := json.Marshal(req)

	// 调用sdk service
	// TODO:此处调用的链接为临时解决方案，后续会给出serverless的sdk，使用serverless的sdk来获取此链接
	request, err := http.NewRequest(method, url, bytes.NewReader(bData))
	if err != nil {
		return
	}

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	var buff = make([]byte, httpResp.ContentLength)
	httpResp.Body.Read(buff)

	// 反序列化结果
	err = json.Unmarshal(buff, &resp)
	if err != nil {
		return
	}

	if resp.ErrCode != 0 {
		err = errors.New(resp.Msg)
	}
	return
}
