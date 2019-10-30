package shipping_test

import (
	"encoding/json"
	"fmt"
	"lemon-shipping-sdk/bird"
	"net/http"
	"net/url"
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func Test_Bird_SignBird(t *testing.T) {
	d, err := bird.SignBird("{'OrderCode':'','ShipperCode':'SF','LogisticCode':'118954907573'}56da2cf8-c8a2-44b2-b6fa-476cd7d1ba17")
	test.Ok(t, err)
	nd, err := url.QueryUnescape(d)
	test.Ok(t, err)
	test.Equals(t, "OWFhM2I5N2ViM2U2MGRkMjc4YzU2NmVlZWI3ZDk0MmE=", nd)
}
func Test_Bird_Query(t *testing.T) {
	var expDto = bird.RespQueryDto{
		EBusinessId:  "test1347907",
		OrderCode:    "",
		ShipperCode:  "SF",
		LogisticCode: "118461988807",
		Success:      true,

		Reason: "",
		State:  "3",
		Traces: []bird.Trace{
			bird.Trace{
				AcceptTime:    "",
				AcceptStation: "快件已经签收，签收人：张启明[武汉市]",
				Remark:        "已经签收",
			},
			bird.Trace{
				AcceptTime:    "",
				AcceptStation: "快件到达武汉市武昌区徐东大街1号网点[武汉市]",
				Remark:        "到达目的城市",
			},
			bird.Trace{
				AcceptTime:    "",
				AcceptStation: "快件在离开深圳集散中心，发往武汉市[深圳市]",
				Remark:        "离开发件城市",
			},
			bird.Trace{
				AcceptTime:    "",
				AcceptStation: "快件已经到达深圳集散中心[深圳市]",
				Remark:        "",
			},
			bird.Trace{
				AcceptTime:    "",
				AcceptStation: "深圳福田保税区网点已揽件[深圳市]",
				Remark:        "已揽件",
			},
		},
	}
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqQueryDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "1002",
			DataType:    "2",
		},
		RequestData: bird.ReqQueryDataDto{
			OrderCode:    "012657018199",
			ShipperCode:  "SF",
			LogisticCode: "118461988807",
			IsHandleInfo: "0",
		},
	}
	statusCode, respDto, err := bird.Query(reqDto, custDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto.EBusinessId, respDto.EBusinessId)
	test.Equals(t, expDto.OrderCode, respDto.OrderCode)
	test.Equals(t, expDto.ShipperCode, respDto.ShipperCode)
	test.Equals(t, expDto.LogisticCode, respDto.LogisticCode)
	test.Equals(t, expDto.Success, respDto.Success)

	test.Equals(t, expDto.Reason, respDto.Reason)
	test.Equals(t, expDto.State, respDto.State)

	for k := range respDto.Traces {
		respDto.Traces[k].AcceptTime = ""
	}
	test.Equals(t, expDto.Traces, respDto.Traces)

}
func Test_Bird_Create(t *testing.T) {
	var expDto = bird.RespCreateDto{
		EBusinessId:          "test1347907",
		Success:              true,
		ResultCode:           "100",
		Reason:               "",
		UniquerRequestNumber: "f456d3f1-a952-44da-887f-f171c74dd025",
		Order: bird.ReqOrder{
			OrderCode:    "012657018199",
			ShipperCode:  "SF",
			LogisticCode: "928989812312",
		},
	}
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqCreateDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "1001",
			DataType:    "2",
		},
		RequestData: bird.ReqCreateDataDto{
			OrderCode:   "012657018199",
			ShipperCode: "SF",
			PayType:     1,
			MonthCode:   "7553045845",
			ExpType:     1,
			Cost:        1,
			OtherCost:   1,
			Sender: bird.Sender{
				Company:      "LV",
				Name:         "Taylor",
				Mobile:       "15018442396",
				ProvinceName: "上海",
				CityName:     "上海市",
				ExpAreaName:  "青浦区",
				Address:      "明珠路",
			},
			Receiver: bird.Receiver{
				Company:      "GCCUI",
				Name:         "Yann",
				Mobile:       "15018442396",
				ProvinceName: "北京",
				CityName:     "北京市",
				ExpAreaName:  "朝阳区",
				Address:      "三里屯街道",
			},
			Weight:   1,
			Quantity: 1,
			Volume:   0,
			Remark:   "小心轻放",
			Commoditys: []bird.Commodity{
				bird.Commodity{
					GoodsName:     "鞋子",
					GoodsCode:     "",
					Goodsquantity: 1,
					GoodsPrice:    0,
					GoodsWeight:   1,

					GoodsDesc: "",
					GoodsVol:  0,
				},
			},
			AddServices: []bird.AddService{
				bird.AddService{
					Name:       "COD",
					Value:      "1020",
					CustomerId: "1234",
				},
			},
		},
	}
	statusCode, respDto, err := bird.Create(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
func Test_Bird_Subscribe(t *testing.T) {
	var expDto bird.RespSubscribeDto
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqSubscribeDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "1008",
			DataType:    "2",
		},
		RequestData: bird.ReqSubscribeDataDto{
			ShipperCode:  "SF",
			LogisticCode: "928989812312",
			PayType:      1,
			Sender: bird.Sender{
				Company:      "LV",
				Name:         "Taylor",
				Mobile:       "15018442396",
				ProvinceName: "上海",
				CityName:     "上海市",
				ExpAreaName:  "青浦区",
				Address:      "明珠路",
			},
			Receiver: bird.Receiver{
				Company:      "GCCUI",
				Name:         "Yann",
				Mobile:       "15018442396",
				ProvinceName: "北京",
				CityName:     "北京市",
				ExpAreaName:  "朝阳区",
				Address:      "三里屯街道",
			},
			Weight:   1,
			Quantity: 1,
			Volume:   0,
			Remark:   "小心轻放",
			Commoditys: []bird.Commodity{
				bird.Commodity{
					GoodsName:     "鞋子",
					GoodsCode:     "",
					Goodsquantity: 1,
					GoodsPrice:    0,
					GoodsWeight:   1,

					GoodsDesc: "",
					GoodsVol:  0,
				},
			},
			AddServices: []bird.AddService{
				bird.AddService{
					Name:       "COD",
					Value:      "1020",
					CustomerId: "1234",
				},
			},
		},
	}
	statusCode, respDto, err := bird.Subscribe(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
func Test_Bird_Recognize(t *testing.T) {
	var expDto bird.RespRecognizeDto
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqRecognizeDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "2002",
			DataType:    "2",
		},
		RequestData: bird.ReqRecognizeDataDto{
			LogisticCode: "51080222000604",
		},
	}

	statusCode, respDto, err := bird.Recognize(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
func Test_Bird_ECreate(t *testing.T) {
	var expDto bird.RespECreateDto
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqECreateDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "1007",
			DataType:    "2",
		},
		RequestData: bird.ReqECreateDataDto{
			OrderCode:   "1234560",
			ShipperCode: "HTKY",
			PayType:     1,
			ExpType:     1,
			Cost:        1,
			OtherCost:   1,
			Sender: bird.Sender{
				Company:      "LV",
				Name:         "Taylor",
				Mobile:       "15018442396",
				ProvinceName: "上海",
				CityName:     "上海市",
				ExpAreaName:  "青浦区",
				Address:      "明珠路",
			},
			Receiver: bird.Receiver{
				Company:      "GCCUI",
				Name:         "Yann",
				Mobile:       "15018442396",
				ProvinceName: "北京",
				CityName:     "北京市",
				ExpAreaName:  "朝阳区",
				Address:      "三里屯街道",
			},
			Weight:                1,
			Quantity:              1,
			Volume:                0,
			Remark:                "小心轻放",
			IsReturnPrintTemplate: "0",
			Commoditys: []bird.Commodity{
				bird.Commodity{
					GoodsName:     "鞋子",
					GoodsCode:     "",
					Goodsquantity: 1,
					GoodsPrice:    0,
					GoodsWeight:   1,

					GoodsDesc: "",
					GoodsVol:  0,
				},
			},
			AddServices: []bird.AddService{
				bird.AddService{
					Name:       "COD",
					Value:      "1020",
					CustomerId: "1234",
				},
			},
		},
	}
	statusCode, respDto, err := bird.ECreate(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
func Test_Bird_ECancel(t *testing.T) {
	var expDto bird.RespECancelDto
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqECancelDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "1147",
			DataType:    "2",
		},
		RequestData: bird.ReqECancelDataDto{
			OrderCode:    "012657700387",
			ShipperCode:  "HTKY",
			ExpNo:        "900008664480",
			CustomerName: "80238728",
			CustomerPwd:  "c0bfe0ba86b66bae5426303c53db0a8b",
		},
	}

	statusCode, respDto, err := bird.ECancel(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
func Test_Bird_EAvailableNum(t *testing.T) {
	var expDto bird.RespEAvailableNumDto
	custDto := bird.ReqCustomerDto{
		Url:    birdDto.Url,
		ApiKey: birdDto.Key,
	}
	reqDto := bird.ReqEAvailableNumDto{
		ReqBase: bird.ReqBase{
			EBusinessId: birdDto.EBusinessId,
			RequestType: "1127",
			DataType:    "2",
		},
		RequestData: bird.ReqEAvailableNumDataDto{
			ShipperCode:  "UC",
			CustomerName: "80238728",
			CustomerPwd:  "c0bfe0ba86b66bae5426303c53db0a81",
			StationCode:  "3001",
			StationName:  "福田网点",
		},
	}

	statusCode, respDto, err := bird.EAvailableNum(reqDto, custDto)
	fmt.Println(statusCode, err)
	res2B, _ := json.Marshal(respDto)
	fmt.Println(string(res2B))
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, statusCode)
	test.Equals(t, expDto, respDto)
}
