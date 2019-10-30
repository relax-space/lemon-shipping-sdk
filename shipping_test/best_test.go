package shipping_test

import (
	"lemon-shipping-sdk/best"
	"net/http"
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func Test_Best_Create(t *testing.T) {
	expDto := best.RespBase{
		Result:           false,
		Note:             "",
		ErrorCode:        "TRANSPORT_ORDER_EXISTED",
		ErrorDescription: "运输单（编码[best-016]）已存在，不能重复创建",
	}
	custDto := best.ReqCustomerDto{
		Url: bestDto.Url,
	}
	reqDto := best.ReqCreateDto{
		ReqBase: best.ReqBase{
			PartnerID:  bestDto.Pid,
			PartnerKey: bestDto.Key,
		},
		BizData: best.CreateContentDto{
			CustomerCode:  "ELAND",
			CustomerName:  "衣念",
			ProjectCode:   "ELAND_YLO2O",
			TransportMode: "EXP",
			//ShippingOrderNo:"",
			OrderCode: "best-016",
			Sender: best.SenderDto{
				Name:            "shipping",
				Province:        "北京市",
				City:            "北京市",
				District:        "朝阳区",
				Address:         "恒通商务园B37座",
				ContactName:     "chengcuicui",
				Phone:           "13051224588",
				EarlyPickupTime: "2018-03-06 14:17",
				LatePickupTime:  "2019-03-06 14:17",
			},
			Receiver: best.ReceiverDto{
				Name:              "CR01",
				Province:          "北京市",
				City:              "北京市",
				District:          "朝阳区",
				Address:           "恒通商务园B37座",
				ContactName:       "guoguo",
				Phone:             "13051224566",
				EarlyDeliveryTime: "2018-04-06 14:17",
				LateDeliveryTime:  "2019-04-06 14:17",
			},
			Items: []best.ItemDto{
				best.ItemDto{
					ItemCode:       "S0000001",
					ItemName:       "款号1",
					PackageCount:   2,
					PackageUomCode: "件",
				},
				best.ItemDto{
					ItemCode:       "S0000002",
					ItemName:       "款号2",
					PackageCount:   5,
					PackageUomCode: "件",
				},
			},
		},
	}

	statusCode, respDto, err := best.Create(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}

func Test_Best_Cancel(t *testing.T) {
	expDto := best.RespBase{
		Result:           false,
		Note:             "",
		ErrorCode:        "",
		ErrorDescription: "运输单（编码[best-000]）不存在",
	}
	custDto := best.ReqCustomerDto{
		Url: bestDto.Url,
	}
	reqDto := best.ReqCancelDto{
		ReqBase: best.ReqBase{
			PartnerID:  bestDto.Pid,
			PartnerKey: bestDto.Key,
		},
		BizData: best.CancelContentDto{
			CustomerCode: "ELAND",
			OrderCode:    "best-000",
		},
	}

	statusCode, respDto, err := best.Cancel(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}

func Test_Best_Query(t *testing.T) {
	expDto := best.RespBase{
		Result:           false,
		Note:             "",
		ErrorCode:        "",
		ErrorDescription: "",
	}
	custDto := best.ReqCustomerDto{
		Url: bestDto.Url,
	}
	reqDto := best.ReqQueryDto{
		ReqBase: best.ReqBase{
			PartnerID:  bestDto.Pid,
			PartnerKey: bestDto.Key,
		},
		BizData: best.QueryContentDto{
			CustomerCode:   "TMS00319",
			OrderCode:      "58dc6f973c37820001224106",
			CreateTimeFrom: "2017-03-29 19:23:39",
			CreateTimeTo:   "2017-04-01 19:23:39",
			ShipmentCode:   "310016648065",
		},
	}
	statusCode, respDto, err := best.Query(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
