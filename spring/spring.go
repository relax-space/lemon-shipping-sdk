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
func Create(reqDto *ReqCreateDto, custDto *ReqCustomerDto) (code string, respDto *RespCreateDto, err error) {
	reqDto.Func = "order.submit"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		code = E02
		return
	}
	data, err := toBase64(reqDto)
	if err != nil {
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(custDto.Url,
		[]byte("data="+data),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil || resp == nil {
		code = E01
		return
	}
	if resp.StatusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", resp.StatusCode)
		return
	}
	newBody := strings.Replace(string(body), "\ufeff", "", -1)
	var respCommonDto RespCommonDto
	err = json.Unmarshal([]byte(newBody), &respCommonDto)
	if err != nil {
		code = E04
		return
	}
	if respCommonDto.Result != "true" {
		code = E03
		err = fmt.Errorf("%v-%v", respCommonDto.Remark, respCommonDto.Info)
		return
	}
	var createInfoDto CreateInfoDto
	err = mapstruct.Decode(respCommonDto.Info, &createInfoDto)
	if err != nil {
		code = E04
		return
	}
	respDto = &RespCreateDto{
		RespBase: respCommonDto.RespBase,
		Info:     &createInfoDto,
	}

	code = SUC
	return
}

func Cancel(reqDto *ReqCancelDto, custDto *ReqCustomerDto) (code string, respDto *RespCancelDto, err error) {
	reqDto.Func = "order.cancel"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		code = E02
		return
	}
	data, err := toBase64(reqDto)
	if err != nil {
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(custDto.Url,
		[]byte("data="+data),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil || resp == nil {
		code = E01
		return
	}
	if resp.StatusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", resp.StatusCode)
		return
	}
	newBody := strings.Replace(string(body), "\ufeff", "", -1)
	var respCommonDto RespCommonDto
	err = json.Unmarshal([]byte(newBody), &respCommonDto)
	if err != nil {
		code = E04
		return
	}
	if respCommonDto.Result != "true" {
		code = E03
		err = fmt.Errorf("%v-%v", respCommonDto.Remark, respCommonDto.Info)
		return
	}
	infoObjs, ok := respCommonDto.Info.([]interface{})
	if !ok {
		code = E04
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

	code = SUC
	return
}

func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (code string, respDto *RespQueryDto, err error) {
	reqDto.Func = "order.trace"
	reqDto.Datetime = time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
	signParam := reqDto.Code + reqDto.Password + reqDto.Datetime
	reqDto.VerifyCode, err = sign.GetMD5Hash(signParam, true)
	if err != nil {
		code = E02
		return
	}
	data, err := toBase64(reqDto)
	if err != nil {
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(custDto.Url,
		[]byte("data="+data),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil || resp == nil {
		code = E01
		return
	}
	if resp.StatusCode != http.StatusOK {
		code = E01
		err = fmt.Errorf("http status exp:200,act:%v", resp.StatusCode)
		return
	}
	newBody := strings.Replace(string(body), "\ufeff", "", -1)
	var respCommonDto RespCommonDto
	err = json.Unmarshal([]byte(newBody), &respCommonDto)
	if err != nil {
		code = E04
		return
	}
	if respCommonDto.Result != "true" {
		code = E03
		err = fmt.Errorf("%v-%v", respCommonDto.Remark, respCommonDto.Info)
		return
	}
	infoSlice, ok := respCommonDto.Info.([]interface{})
	if !ok || len(infoSlice) == 0 {
		code = E04
		err = errors.New("spring response format error. ")
		return
	}
	var infoDto []QueryInfoDto
	err = mapstruct.Decode(infoSlice[0], &infoDto)
	if err != nil {
		code = E04
		return
	}
	respDto = &RespQueryDto{
		RespBase: respCommonDto.RespBase,
		Info:     &infoDto,
	}

	code = SUC
	return
}
