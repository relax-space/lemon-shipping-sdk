package best

import (
	"fmt"
	"os"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_Create(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url: os.Getenv("BEST_URL_DEV"),
	}
	reqDto := &ReqCreateDto{
		ReqBase: &ReqBase{
			PartnerID:  os.Getenv("BEST_PID_DEV"),
			PartnerKey: os.Getenv("BEST_KEY_DEV"),
		},
		BizData: &CreateContentDto{
			CustomerCode:  "ELAND",
			CustomerName:  "衣念",
			ProjectCode:   "ELAND_YLO2O",
			TransportMode: "EXP",
			//ShippingOrderNo:"",
			OrderCode: "best-014",
			Sender: &SenderDto{
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
			Receiver: &ReceiverDto{
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
			Items: &ItemsDto{
				&[]ItemDto{
					ItemDto{
						ItemCode:       "S0000001",
						ItemName:       "款号1",
						PackageCount:   2,
						PackageUomCode: "件",
					},
					ItemDto{
						ItemCode:       "S0000002",
						ItemName:       "款号2",
						PackageCount:   5,
						PackageUomCode: "件",
					},
				},
			},
		},
	}
	_, _, respDto, err := Create(reqDto, custDto)
	fmt.Printf("%+v", respDto)
	fmt.Println("")
	fmt.Printf("%+v", err)
	fmt.Println("")
	test.Ok(t, err)
}

func Test_Cancel(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url: os.Getenv("BEST_URL_DEV"),
	}
	reqDto := &ReqCancelDto{
		ReqBase: &ReqBase{
			PartnerID:  os.Getenv("BEST_PID_DEV"),
			PartnerKey: os.Getenv("BEST_KEY_DEV"),
		},
		BizData: &CancelContentDto{
			CustomerCode: "ELAND",
			OrderCode:    "best-000",
		},
	}
	_, _, respDto, err := Cancel(reqDto, custDto)
	fmt.Printf("%+v", respDto)
	fmt.Println("")
	fmt.Printf("%+v", err)
	fmt.Println("")
	test.Ok(t, err)
}

func Test_Query(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url: os.Getenv("BEST_URL"),
	}
	reqDto := &ReqQueryDto{
		ReqBase: &ReqBase{
			PartnerID:  os.Getenv("BEST_PID"),
			PartnerKey: os.Getenv("BEST_KEY"),
		},
		BizData: &QueryContentDto{
			CustomerCode:   "TMS00319",
			OrderCode:      "58dc6f973c37820001224106",
			CreateTimeFrom: "2017-03-29 19:23:39",
			CreateTimeTo:   "2017-04-01 19:23:39",
			ShipmentCode:   "310016648065",
		},
	}
	_, _, respDto, err := Query(reqDto, custDto)
	fmt.Printf("%+v", respDto)
	fmt.Println("")
	fmt.Printf("%+v", err)
	fmt.Println("")
	test.Ok(t, err)
}
