package shipping_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"lemon-shipping-sdk/sf"
	"net/http"
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func Test_Sf_Create(t *testing.T) {
	var expDto sf.RespCreateDto
	reqDto := sf.ReqCreateDto{
		ReqBaseDto: sf.ReqBaseDto{
			Head:    sfDto.Linkcode,
			Service: "OrderService",
			Lang:    "zh-CN",
		},
		Body: sf.ReqBodyCreateDto{
			sf.ReqOrderCreateDto{
				OrderId:  "1691674864558491516",
				MailNo:   "444834120517",
				JCompany: "ELAND官方旗舰店",
				JContact: "ELAND官方旗舰店",
				JTel:     "13006543727",
				JMobile:  "15005466698",

				JProvince: "山东",
				JCity:     "东营市",
				JCounty:   "",
				JAddress:  "东营 东营市济南路296号银座二楼衣恋专柜",

				DCompany: "武海花",
				DContact: "武海花",
				DTel:     "13979441452",
				DMobile:  "13979441452",

				DProvince: "江西省",
				DCity:     "抚州市",
				DCounty:   "临川区",
				DAddress:  "上顿渡镇江西省抚州市临川区上顿渡临川一中北门吁坊新村",
				CustId:    "9999999999",

				PayMethod:      3,
				IsDocall:       0,
				ExpressType:    "2",
				ParcelQuantity: 0,
				Cargo: sf.CargoDto{
					Name:  "EEBW63751M39165",
					Unit:  "件",
					Count: 0,
				},
			},
		},
	}
	CustDto := sf.ReqCustomerDto{
		Url:       sfDto.Url,
		Checkword: sfDto.CheckWord,
	}
	statusCode, respDto, err := sf.Create(reqDto, CustDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)

}

func Test_Sf_Route(t *testing.T) {
	var expDto sf.RespRouteDto
	reqDto := sf.ReqRouteDto{
		ReqBaseDto: sf.ReqBaseDto{
			Head:    sfDto.Linkcode,
			Service: "RouteService",
			Lang:    "zh-CN",
		},
		Body: sf.ReqBodyRouteDto{
			RouteRequest: sf.ReqRouteRequestDto{
				TrackingType:   1,
				TrackingNumber: "201712250001",
				MethodType:     1,
			},
		},
	}
	custDto := sf.ReqCustomerDto{
		Url:       sfDto.Url,
		Checkword: sfDto.CheckWord,
	}
	statusCode, respDto, err := sf.Route(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)

}

func Test_Sf_Confirm(t *testing.T) {
	var expDto sf.RespConfirmDto
	reqDto := sf.ReqConfirmDto{
		ReqBaseDto: sf.ReqBaseDto{
			Head:    sfDto.Linkcode,
			Service: "OrderService",
			Lang:    "zh-CN",
		},
		Body: sf.ReqBodyConfirmDto{
			OrderConfirm: sf.ReqOrderConfirmDto{
				OrderId:       "TE201500104",
				MailNo:        "444003078326",
				DealType:      1,
				CustomsBatchs: "",
				AgentNo:       "",

				ConsignEmpCode: "",
				SourceZoneCode: "",
				OrderConfirmOption: sf.ReqOrderConfirmOptionDto{
					Weight:         3.56,
					Volumn:         "33,33,33",
					ReturnTracting: "",
					ExpressType:    "",
					ChildrenNos:    "",

					WaybillSize:     0,
					IsGenEletricPic: 0,
				},
			},
		},
	}
	custDto := sf.ReqCustomerDto{
		Url:       sfDto.Url,
		Checkword: sfDto.CheckWord,
	}
	statusCode, respDto, err := sf.Confirm(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)

}

func Test_Sf_Push(t *testing.T) {
	expString := `<Request service="RouterPushService" lang="zh-CN"><Body><WaybillRoute id="10049361064087" orderid="TE201500106" mailno="444003079772" acceptTime="2015-01-04 17:42:00" acceptAddress="深圳" remark="上门收件" opCode="50"></WaybillRoute></Body></Request>`
	reqDto := sf.ReqPushDto{
		ReqBaseDto: sf.ReqBaseDto{
			Service: "RouterPushService",
			Lang:    "zh-CN",
		},
		Body: sf.ReqBodyPushDto{
			WaybillRoute: sf.ReqWaybillRouteDto{
				Id:           10049361064087,
				OrderId:      "TE201500106",
				MailNo:       "444003079772",
				AcceptTime:   "2015-01-04 17:42:00",
				AcceptAddres: "深圳",

				Remark: "上门收件",
				OpCode: "50",
			},
		},
	}
	b, err := xml.Marshal(reqDto)
	test.Ok(t, err)
	test.Equals(t, expString, string(b))
}
