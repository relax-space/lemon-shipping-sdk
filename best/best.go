package best

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

/*
1.base64 param
2.request api
*/
func Create(reqDto *ReqCreateDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespBase, err error) {
	reqDto.ServiceType = "TMS_CREATE_ORDER"
	bizData, err := xml.Marshal(reqDto.BizData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(bizData) + reqDto.PartnerKey
	reqDto.Sign, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		code = E02
		return
	}
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = reqDto.ServiceType
	reqMap["partnerID"] = reqDto.PartnerID
	reqMap["bizData"] = string(bizData)
	reqMap["sign"] = reqDto.Sign

	data := base.JoinMapString(reqMap)

	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.XmlType
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
	if respDto.Result != true {
		code = E03
		err = fmt.Errorf("%v-%v", respDto.ErrorCode, respDto.ErrorDescription)
		return
	}
	code = SUC
	return
}

func Cancel(reqDto *ReqCancelDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespBase, err error) {
	reqDto.ServiceType = "TMS_CANCEL_ORDER"
	bizData, err := xml.Marshal(reqDto.BizData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(bizData) + reqDto.PartnerKey
	reqDto.Sign, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		code = E02
		return
	}
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = reqDto.ServiceType
	reqMap["partnerID"] = reqDto.PartnerID
	reqMap["bizData"] = string(bizData)
	reqMap["sign"] = reqDto.Sign

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.XmlType
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
	if respDto.Result != true {
		code = E03
		err = fmt.Errorf("%v-%v", respDto.ErrorCode, respDto.ErrorDescription)
		return
	}

	code = SUC
	return
}

func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespQueryDto, err error) {
	reqDto.ServiceType = "TMS_TRACE_QUERY"
	bizData, err := xml.Marshal(reqDto.BizData)
	if err != nil {
		code = E02
		return
	}
	signParam := string(bizData) + reqDto.PartnerKey
	reqDto.Sign, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		code = E02
		return
	}
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = reqDto.ServiceType
	reqMap["partnerID"] = reqDto.PartnerID
	reqMap["bizData"] = string(bizData)
	reqMap["sign"] = reqDto.Sign

	data := base.JoinMapString(reqMap)
	req := httpreq.New(http.MethodPost, custDto.Url, data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		httpReq.RespDataType = httpreq.XmlType
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
	if respDto.Result != true {
		code = E03
		err = fmt.Errorf("%v-%v", respDto.Errors, respDto.OrderInfos)
		return
	}

	code = SUC
	return
}
