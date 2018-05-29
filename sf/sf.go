package sf

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

func Create(reqDto *ReqCreateDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespCreateDto, err error) {
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		code = E02
		return
	}
	xmlData := string(b)
	verifyCode, err := sign.GetMD5Base64(xmlData + custDto.Checkword)
	if err != nil {
		code = E01
		return
	}
	reqParam := map[string]string{
		"xml":        xmlData,
		"verifyCode": verifyCode,
	}
	respDto = &RespCreateDto{}
	statusCode, err = httpreq.New(http.MethodPost, custDto.Url,
		base.JoinMapStringEncode(reqParam), func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}

	if respDto.Head != "OK" {
		code = E03
		err = fmt.Errorf("%v", respDto.Error.Error)
		return
	}

	code = SUC
	return

}

func Route(reqDto *ReqRouteDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespRouteDto, err error) {
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		code = E02
		return
	}
	xmlData := string(b)
	verifyCode, err := sign.GetMD5Base64(xmlData + custDto.Checkword)
	if err != nil {
		code = E01
		return
	}
	reqParam := map[string]string{
		"xml":        xmlData,
		"verifyCode": verifyCode,
	}
	respDto = &RespRouteDto{}

	statusCode, err = httpreq.New(http.MethodPost, custDto.Url,
		base.JoinMapStringEncode(reqParam), func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}

	if respDto.Head != "OK" {
		code = E03
		err = fmt.Errorf("%v", respDto.Error.Error)
		return
	}

	code = SUC
	return

}

func Confirm(reqDto *ReqConfirmDto, custDto *ReqCustomerDto) (statusCode int, code string, respDto *RespConfirmDto, err error) {
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		code = E02
		return
	}
	xmlData := string(b)
	verifyCode, err := sign.GetMD5Base64(xmlData + custDto.Checkword)
	if err != nil {
		code = E01
		return
	}
	reqParam := map[string]string{
		"xml":        xmlData,
		"verifyCode": verifyCode,
	}
	respDto = &RespConfirmDto{}

	statusCode, err = httpreq.New(http.MethodPost, custDto.Url,
		base.JoinMapStringEncode(reqParam), func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		code = E01
		return
	}
	if statusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", statusCode)
		return
	}

	if respDto.Head != "OK" {
		code = E03
		err = fmt.Errorf("%v", respDto.Error.Error)
		return
	}

	code = SUC
	return

}
