package shipping_test

import (
	"lemon-shipping-sdk/spring"
	"net/http"
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func Test_Spring_Create(t *testing.T) {
	var expDto = spring.RespCommonDto{
		RespBase: spring.RespBase{
			Result:    "true",
			ErrorCode: "",
			Remark:    "提交成功(Submitted successfully)",
		},
	}
	custDto := spring.ReqCustomerDto{
		Url: springDto.Url,
	}
	reqDto := spring.ReqCreateDto{
		ReqBase: spring.ReqBase{
			Code:     springDto.Code,
			Password: springDto.Pwd,
		},
		Content: spring.CreateContentDto{
			Quantity: 1,
			Sender: spring.SenderDto{
				Code:    "CR00",
				Name:    "xiaoxinmiao",
				Company: "Eland",
				Phone:   "15811016818",
				City:    "北京市,北京市,通州区",
				Address: "恒通商务园B37座",
			},
			Receiver: spring.ReceiverDto{
				Code:    "CR02",
				Name:    "xiaomiao",
				Company: "Eland",
				Phone:   "15811016818",
				City:    "北京市,北京市,通州区",
				Address: "恒通商务园B51座",
			},
		},
	}
	statusCode, respDto, err := spring.Create(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.RespBase, respDto.RespBase)
}

func Test_Spring_Cancel(t *testing.T) {
	var expDto = spring.RespCommonDto{
		RespBase: spring.RespBase{
			Result:    "false",
			ErrorCode: "cancel_02",
			Remark:    "运单号不存在(The waybill does not exist)",
		},
		Info: []interface{}{"7100046801"},
	}
	custDto := spring.ReqCustomerDto{
		Url: springDto.Url,
	}
	reqDto := spring.ReqCancelDto{
		ReqBase: spring.ReqBase{
			Code:     springDto.Code,
			Password: springDto.Pwd,
		},
		Content: spring.CancelContentDto{
			BillNo: "7100046801",
		},
	}
	statusCode, respDto, err := spring.Cancel(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}

func Test_Spring_Query(t *testing.T) {
	var expDto = spring.RespCommonDto{
		RespBase: spring.RespBase{
			Result:    "true",
			ErrorCode: "",
			Remark:    "",
		},
	}
	var expInfo = map[string]interface{}{
		"billno":     "7100091086",
		"state":      "A1",
		"state_info": "下单成功",
		"time":       "",
	}
	custDto := spring.ReqCustomerDto{
		Url: springDto.Url,
	}
	reqDto := spring.ReqQueryDto{
		ReqBase: spring.ReqBase{
			Code:     springDto.Code,
			Password: springDto.Pwd,
		},
		Content: spring.QueryContentDto{
			OrderNo: "1910301777836",
		},
	}
	statusCode, respDto, err := spring.Query(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	info1, ok1 := respDto.Info.([]interface{})
	test.Equals(t, true, ok1)
	info2, ok2 := info1[0].([]interface{})
	test.Equals(t, true, ok2)

	info, ok := info2[0].(map[string]interface{})
	test.Equals(t, true, ok)

	test.Equals(t, expDto.RespBase, respDto.RespBase)
	test.Equals(t, expInfo["billno"], info["billno"])
	test.Equals(t, expInfo["state"], info["state"])
	test.Equals(t, expInfo["state_info"], info["state_info"])
}

func Test_Spring_GetOrdernoByBillno(t *testing.T) {
	var expDto = spring.RespCommonDto{
		RespBase: spring.RespBase{
			Result:    "true",
			ErrorCode: "",
			Remark:    "",
		},
		Info: map[string]interface{}{
			"billno":  "7100091086",
			"orderno": "1910301777836",
		},
	}
	custDto := spring.ReqCustomerDto{
		Url: springDto.Url,
	}
	reqDto := spring.ReqGetOrdernoByBillnoDto{
		ReqBase: spring.ReqBase{
			Code:     springDto.Code,
			Password: springDto.Pwd,
		},
		Content: spring.GetOrdernoByBillnoContentDto{
			BillNo: "7100091086",
		},
	}
	statusCode, respDto, err := spring.GetOrdernoByBillno(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
