package bird

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pangpanglabs/goutils/httpreq"
	"github.com/relax-space/go-kit/base"
)

func Query(reqDto ReqQueryDto, custDto ReqCustomerDto) (int, RespQueryDto, error) {
	var respDto RespQueryDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil
}

func Create(reqDto ReqCreateDto, custDto ReqCustomerDto) (int, RespCreateDto, error) {
	var respDto RespCreateDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil
}

func Subscribe(reqDto ReqSubscribeDto, custDto ReqCustomerDto) (int, RespSubscribeDto, error) {
	var respDto RespSubscribeDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err

	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err

	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil

}

func Recognize(reqDto ReqRecognizeDto, custDto ReqCustomerDto) (int, RespRecognizeDto, error) {
	var respDto RespRecognizeDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil
}

func ECreate(reqDto ReqECreateDto, custDto ReqCustomerDto) (int, RespECreateDto, error) {
	var respDto RespECreateDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}
func ECancel(reqDto ReqECancelDto, custDto ReqCustomerDto) (int, RespECancelDto, error) {
	var respDto RespECancelDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}
func EAvailableNum(reqDto ReqEAvailableNumDto, custDto ReqCustomerDto) (int, RespEAvailableNumDto, error) {
	var respDto RespEAvailableNumDto
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = SignBird(signParam)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqMap := make(map[string]string, 0)
	reqMap["EBusinessID"] = reqDto.EBusinessId
	reqMap["RequestType"] = reqDto.RequestType
	reqMap["DataSign"] = reqDto.DataSign
	reqMap["DataType"] = reqDto.DataType
	reqMap["RequestData"] = url.QueryEscape(string(b))

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.JsonType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}
