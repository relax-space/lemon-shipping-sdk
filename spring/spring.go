package spring

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/relax-space/go-kit/httpreq"
	"github.com/relax-space/go-kitt/mapstruct"

	"github.com/relax-space/go-kit/sign"
)

/*
1.base64 param
2.request api
*/
func Create(reqDto *ReqCreateDto, custDto *ReqCustomerDto) (resp *http.Response, respDto *RespCreateDto, err error) {
	reqDto.Func = "order.submit"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return
	}
	resp, body, err := httpreq.NewPost(custDto.Url,
		[]byte("data="+data),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil || resp == nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status exp:200,act:%v", resp.StatusCode)
		return
	}
	newBody := strings.Replace(string(body), "\ufeff", "", -1)
	var respCommonDto RespCommonDto
	err = json.Unmarshal([]byte(newBody), &respCommonDto)
	if err != nil {
		return
	}
	if respCommonDto.Result != "true" {
		err = fmt.Errorf("message from spring:%v-%v", respCommonDto.Remark, respCommonDto.Info)
		return
	}
	var createInfoDto CreateInfoDto
	err = mapstruct.Decode(respCommonDto.Info, &createInfoDto)
	if err != nil {
		return
	}
	respDto = &RespCreateDto{
		RespBase: respCommonDto.RespBase,
		Info:     &createInfoDto,
	}

	return
}

func Cancel(reqDto *ReqCancelDto, custDto *ReqCustomerDto) (resp *http.Response, respDto *RespCancelDto, err error) {
	reqDto.Func = "order.cancel"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return
	}
	resp, body, err := httpreq.NewPost(custDto.Url,
		[]byte("data="+data),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil || resp == nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status exp:200,act:%v", resp.StatusCode)
		return
	}
	newBody := strings.Replace(string(body), "\ufeff", "", -1)
	var respCommonDto RespCommonDto
	err = json.Unmarshal([]byte(newBody), &respCommonDto)
	if err != nil {
		return
	}
	if respCommonDto.Result != "true" {
		err = fmt.Errorf("message from spring:%v-%v", respCommonDto.Remark, respCommonDto.Info)
		return
	}
	infoObjs, ok := respCommonDto.Info.([]interface{})
	if !ok {
		err = errors.New("spring response format error. ")
		return
	}
	var infos []string
	for _, v := range infoObjs {
		infos = append(infos, v.(string))
	}
	respDto = &RespCancelDto{
		RespBase: respCommonDto.RespBase,
		Info:     infos,
	}

	return
}

func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (resp *http.Response, respDto *RespQueryDto, err error) {
	reqDto.Func = "order.trace"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		return
	}
	data, err := toBase64(reqDto)
	if err != nil {
		return
	}
	resp, body, err := httpreq.NewPost(custDto.Url,
		[]byte("data="+data),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil || resp == nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status exp:200,act:%v", resp.StatusCode)
		return
	}
	newBody := strings.Replace(string(body), "\ufeff", "", -1)
	var respCommonDto RespCommonDto
	err = json.Unmarshal([]byte(newBody), &respCommonDto)
	if err != nil {
		return
	}
	if respCommonDto.Result != "true" {
		err = fmt.Errorf("message from spring:%v-%v", respCommonDto.Remark, respCommonDto.Info)
		return
	}
	infoSlice, ok := respCommonDto.Info.([]interface{})
	if !ok || len(infoSlice) == 0 {
		err = errors.New("spring response format error. ")
		return
	}
	var infoDto []QueryInfoDto
	err = mapstruct.Decode(infoSlice[0], &infoDto)
	if err != nil {
		return
	}
	respDto = &RespQueryDto{
		RespBase: respCommonDto.RespBase,
		Info:     &infoDto,
	}

	return
}
