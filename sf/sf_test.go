package sf

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_Create(t *testing.T) {
	reqDto := &ReqCreateDto{
		ReqBaseDto: &ReqBaseDto{
			Head:    os.Getenv("SF_Linkcode"),
			Service: "OrderService",
			Lang:    "zh-CN",
		},
		Body: &ReqBodyCreateDto{
			&ReqOrderCreateDto{
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
				Cargo: &CargoDto{
					Name:  "EEBW63751M39165",
					Unit:  "件",
					Count: 0,
				},
			},
		},
	}
	CustDto := &ReqCustomerDto{
		Url:       os.Getenv("SF_URL"),
		Checkword: os.Getenv("SF_CHECKWORD"),
	}
	statusCode, code, respDto, err := Create(reqDto, CustDto)
	fmt.Printf("http status:%+v\n", statusCode)
	fmt.Printf("code:%+v\n", code)
	fmt.Printf("err:%+v\n", err)
	if respDto != nil && respDto.Body != nil {
		fmt.Printf("body:%+v\n", *(respDto.Body))
		if respDto.Body.OrderResponse != nil {
			fmt.Printf("respDto.Body.OrderResponse:%+v\n", *(respDto.Body.OrderResponse))
		}
	}
	test.Ok(t, err)
}

func Test_Route(t *testing.T) {
	reqDto := &ReqRouteDto{
		ReqBaseDto: &ReqBaseDto{
			Head:    os.Getenv("SF_Linkcode"),
			Service: "RouteService",
			Lang:    "zh-CN",
		},
		Body: &ReqBodyRouteDto{
			RouteRequest: &ReqRouteRequestDto{
				TrackingType:   1,
				TrackingNumber: "201712250001",
				MethodType:     1,
			},
		},
	}
	custDto := &ReqCustomerDto{
		Url:       os.Getenv("SF_URL"),
		Checkword: os.Getenv("SF_CHECKWORD"),
	}
	statusCode, code, respDto, err := Route(reqDto, custDto)
	fmt.Printf("http status:%+v\n", statusCode)
	fmt.Printf("code:%+v\n", code)
	fmt.Printf("err:%+v\n", err)
	fmt.Println(statusCode, code, err)
	if respDto != nil && respDto.Body != nil {
		fmt.Printf("body:%+v\n", *(respDto.Body))
		if respDto.Body.RouteResponse != nil {
			fmt.Printf("respDto.Body.RouteResponse:%+v\n", *(respDto.Body.RouteResponse))
			if respDto.Body.RouteResponse.Route != nil {
				fmt.Printf("respDto.Body.RouteResponse.Route:%+v\n", respDto.Body.RouteResponse.Route)
			}
		}
	}
	test.Ok(t, err)
}

func Test_Confirm(t *testing.T) {
	reqDto := &ReqConfirmDto{
		ReqBaseDto: &ReqBaseDto{
			Head:    os.Getenv("SF_Linkcode"),
			Service: "OrderService",
			Lang:    "zh-CN",
		},
		Body: &ReqBodyConfirmDto{
			OrderConfirm: &ReqOrderConfirmDto{
				OrderId:       "TE201500104",
				MailNo:        "444003078326",
				DealType:      1,
				CustomsBatchs: "",
				AgentNo:       "",

				ConsignEmpCode: "",
				SourceZoneCode: "",
				OrderConfirmOption: &ReqOrderConfirmOptionDto{
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
	custDto := &ReqCustomerDto{
		Url:       os.Getenv("SF_URL"),
		Checkword: os.Getenv("SF_CHECKWORD"),
	}
	statusCode, code, respDto, err := Confirm(reqDto, custDto)
	fmt.Println(statusCode, code, err)
	fmt.Printf("%+v", respDto)
	test.Ok(t, err)
}

func Test_Push(t *testing.T) {
	reqDto := &ReqPushDto{
		ReqBaseDto: &ReqBaseDto{
			Service: "RouterPushService",
			Lang:    "zh-CN",
		},
		Body: &ReqBodyPushDto{
			WaybillRoute: &ReqWaybillRouteDto{
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
	fmt.Println(string(b), err)
}
