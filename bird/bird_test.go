package bird

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/relax-space/go-kit/test"
)

func Test_signBird(t *testing.T) {
	d, err := signBird("{'OrderCode':'','ShipperCode':'SF','LogisticCode':'118954907573'}56da2cf8-c8a2-44b2-b6fa-476cd7d1ba17")
	test.Ok(t, err)
	nd, err := url.QueryUnescape(d)
	fmt.Println(nd)
	test.Ok(t, err)
	test.Equals(t, "OWFhM2I5N2ViM2U2MGRkMjc4YzU2NmVlZWI3ZDk0MmE=", nd)
}

func Test_Query(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url:    os.Getenv("BIRD_URL"),
		ApiKey: os.Getenv("BIRD_APIKEY"),
	}
	reqDto := &ReqQueryDto{
		ReqBase: &ReqBase{
			EBusinessId: os.Getenv("BIRD_EBusinessId"),
			RequestType: "1002",
			DataType:    "2",
		},
		RequestData: &ReqQueryDataDto{
			OrderCode:    "",
			ShipperCode:  "SF",
			LogisticCode: "118461988807",
			IsHandleInfo: "0",
		},
	}
	_, _, respDto, err := Query(reqDto, custDto)
	spew.Dump(respDto)
	test.Ok(t, err)
}
func Test_Create(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url:    os.Getenv("BIRD_URL"),
		ApiKey: os.Getenv("BIRD_APIKEY"),
	}
	reqDto := &ReqCreateDto{
		ReqBase: &ReqBase{
			EBusinessId: os.Getenv("BIRD_EBusinessId"),
			RequestType: "1001",
			DataType:    "2",
		},
		RequestData: &ReqCreateDataDto{
			OrderCode:   "012657018199",
			ShipperCode: "SF",
			PayType:     1,
			MonthCode:   "7553045845",
			ExpType:     1,
			Cost:        1,
			OtherCost:   1,
			Sender: &Sender{
				Company:      "LV",
				Name:         "Taylor",
				Mobile:       "15018442396",
				ProvinceName: "上海",
				CityName:     "上海市",
				ExpAreaName:  "青浦区",
				Address:      "明珠路",
			},
			Receiver: &Receiver{
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
			Commoditys: []Commodity{
				Commodity{
					GoodsName:     "鞋子",
					GoodsCode:     "",
					Goodsquantity: 1,
					GoodsPrice:    0,
					GoodsWeight:   1,

					GoodsDesc: "",
					GoodsVol:  0,
				},
			},
			AddService: []AddService{
				AddService{
					Name:       "COD",
					Value:      "1020",
					CustomerId: "1234",
				},
			},
		},
	}
	spew.Dump(reqDto)
	_, _, respDto, err := Create(reqDto, custDto)
	spew.Dump(respDto)
	test.Ok(t, err)
}

func Test_Subscribe(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url:    os.Getenv("BIRD_URL"),
		ApiKey: os.Getenv("BIRD_APIKEY"),
	}
	reqDto := &ReqSubscribeDto{
		ReqBase: &ReqBase{
			EBusinessId: os.Getenv("BIRD_EBusinessId"),
			RequestType: "1008",
			DataType:    "2",
		},
		RequestData: &ReqSubscribeDataDto{
			ShipperCode:  "SF",
			LogisticCode: "928989812312",
			PayType:      1,
			Sender: &Sender{
				Company:      "LV",
				Name:         "Taylor",
				Mobile:       "15018442396",
				ProvinceName: "上海",
				CityName:     "上海市",
				ExpAreaName:  "青浦区",
				Address:      "明珠路",
			},
			Receiver: &Receiver{
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
			Commoditys: []Commodity{
				Commodity{
					GoodsName:     "鞋子",
					GoodsCode:     "",
					Goodsquantity: 1,
					GoodsPrice:    0,
					GoodsWeight:   1,

					GoodsDesc: "",
					GoodsVol:  0,
				},
			},
			AddService: []AddService{
				AddService{
					Name:       "COD",
					Value:      "1020",
					CustomerId: "1234",
				},
			},
		},
	}
	spew.Dump(reqDto)
	_, _, respDto, err := Subscribe(reqDto, custDto)
	spew.Dump(respDto)
	test.Ok(t, err)
}
