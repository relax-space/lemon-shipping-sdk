package sf

import (
	"encoding/xml"
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

func Create(reqDto ReqCreateDto, custDto ReqCustomerDto) (int, RespCreateDto, error) {
	var respDto RespCreateDto
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	xmlData := string(b)
	verifyCode, err := sign.GetMD5Base64(xmlData + custDto.Checkword)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqParam := map[string]string{
		"xml":        xmlData,
		"verifyCode": verifyCode,
	}
	statusCode, err := httpreq.New(http.MethodPost, custDto.Url,
		base.JoinMapStringEncode(reqParam), func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil

}

func Route(reqDto ReqRouteDto, custDto ReqCustomerDto) (int, RespRouteDto, error) {
	var respDto RespRouteDto
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	xmlData := string(b)
	verifyCode, err := sign.GetMD5Base64(xmlData + custDto.Checkword)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqParam := map[string]string{
		"xml":        xmlData,
		"verifyCode": verifyCode,
	}

	statusCode, err := httpreq.New(http.MethodPost, custDto.Url,
		base.JoinMapStringEncode(reqParam), func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil

}

func Confirm(reqDto ReqConfirmDto, custDto ReqCustomerDto) (int, RespConfirmDto, error) {
	var respDto RespConfirmDto
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	xmlData := string(b)
	verifyCode, err := sign.GetMD5Base64(xmlData + custDto.Checkword)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	reqParam := map[string]string{
		"xml":        xmlData,
		"verifyCode": verifyCode,
	}

	statusCode, err := httpreq.New(http.MethodPost, custDto.Url,
		base.JoinMapStringEncode(reqParam), func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.FormType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil

}
