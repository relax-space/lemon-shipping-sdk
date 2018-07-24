package bird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pangpanglabs/goutils/httpreq"
	"github.com/relax-space/go-kit/base"
)

func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespQueryDto, err error) {
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = signBird(signParam)
	if err != nil {
		code = E02
		return
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
	statusCode, err = req.Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}
	if respDto.Success != true {
		code = E03
		err = fmt.Errorf("%v", respDto.Reason)
		return
	}

	code = SUC
	return
}

func Create(reqDto *ReqCreateDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespCreateDto, err error) {
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = signBird(signParam)
	if err != nil {
		code = E02
		return
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
	statusCode, err = req.Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}
	if respDto.Success != true {
		code = E03
		err = fmt.Errorf("%v", respDto.Reason)
		return
	}

	code = SUC
	return
}

func Subscribe(reqDto *ReqSubscribeDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespSubscribeDto, err error) {
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = signBird(signParam)
	if err != nil {
		code = E02
		return
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
	statusCode, err = req.Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}
	if respDto.Success != true {
		code = E03
		err = fmt.Errorf("%v", respDto.Reason)
		return
	}

	code = SUC
	return
}

func Recognize(reqDto *ReqRecognizeDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespRecognizeDto, err error) {
	b, err := json.Marshal(reqDto.RequestData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(b) + custDto.ApiKey
	reqDto.DataSign, err = signBird(signParam)
	if err != nil {
		code = E02
		return
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
	statusCode, err = req.Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}
	if respDto.Success != true {
		code = E03
		err = fmt.Errorf("%v", respDto.Code)
		return
	}

	code = SUC
	return
}
