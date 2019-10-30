package spring

import (
	"net/http"
	"time"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/relax-space/go-kit/sign"
)

/*
1.base64 param
2.request api
*/
func Create(reqDto ReqCreateDto, custDto ReqCustomerDto) (int, RespCommonDto, error) {
	var respDto RespCommonDto
	var err error
	reqDto.Func = "order.submit"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	req := httpreq.New(http.MethodPost, custDto.Url, "data="+data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil
}

func Cancel(reqDto ReqCancelDto, custDto ReqCustomerDto) (int, RespCommonDto, error) {
	var respDto RespCommonDto
	var err error
	reqDto.Func = "order.cancel"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	req := httpreq.New(http.MethodPost, custDto.Url, "data="+data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}
	return statusCode, respDto, nil
}

func Query(reqDto ReqQueryDto, custDto ReqCustomerDto) (int, RespCommonDto, error) {
	var respDto RespCommonDto
	var err error
	reqDto.Func = "order.trace"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}

	req := httpreq.New(http.MethodPost, custDto.Url, "data="+data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}

func GetOrdernoByBillno(reqDto ReqGetOrdernoByBillnoDto, custDto ReqCustomerDto) (int, RespCommonDto, error) {
	var respDto RespCommonDto
	var err error
	reqDto.Func = "order.query"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	req := httpreq.New(http.MethodPost, custDto.Url, "data="+data, func(httpReq *httpreq.HttpReq) error {
		httpReq.ReqDataType = httpreq.FormType
		return nil
	})
	statusCode, err := req.Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}
