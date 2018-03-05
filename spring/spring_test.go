package spring

import (
	"fmt"
	"os"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_Create(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url: os.Getenv("SPRING_URL"),
	}
	reqDto := &ReqCreateDto{
		ReqBase: &ReqBase{
			Code:     os.Getenv("SPRING_CODE"),
			Password: os.Getenv("SPRING_PWD"),
		},
		Content: &CreateContentDto{
			Quantity: 1,
			Sender: &SenderDto{
				Code:    "CR00",
				Name:    "xiaoxinmiao",
				Company: "Eland",
				Phone:   "15811016818",
				City:    "北京市,北京市,通州区",
				Address: "恒通商务园B37座",
			},
			Receiver: &ReceiverDto{
				Code:    "CR02",
				Name:    "xiaomiao",
				Company: "Eland",
				Phone:   "15811016818",
				City:    "北京市,北京市,通州区",
				Address: "恒通商务园B51座",
			},
		},
	}
	status, respDto, err := Create(reqDto, custDto)
	test.Ok(t, err)
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
}

func Test_Cancel(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url: os.Getenv("SPRING_URL"),
	}
	reqDto := &ReqCancelDto{
		ReqBase: &ReqBase{
			Code:     os.Getenv("SPRING_CODE"),
			Password: os.Getenv("SPRING_PWD"),
		},
		Content: &CancelContentDto{
			BillNo: "7100046801",
		},
	}
	status, respDto, err := Cancel(reqDto, custDto)
	test.Ok(t, err)
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
}

func Test_Query(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url: os.Getenv("SPRING_URL"),
	}
	reqDto := &ReqQueryDto{
		ReqBase: &ReqBase{
			Code:     os.Getenv("SPRING_CODE"),
			Password: os.Getenv("SPRING_PWD"),
		},
		Content: &QueryContentDto{
			OrderNo: "1709051572807",
		},
	}
	status, respDto, err := Query(reqDto, custDto)
	test.Ok(t, err)
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
}
