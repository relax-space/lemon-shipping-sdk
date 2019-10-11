# lemon-shipping-sdk
shipping go-sdk


## Installation
```
go get github.com/relax-space/lemon-shipping-sdk/spring
go get github.com/relax-space/lemon-shipping-sdk/sf
```

## Spring
docs:http://dev.spring56.com:81/doc/?p=api
  >  create  
  >  cancel  
  >  query  

### spring -- create
```go
custDto := &ReqCustomerDto{
		Url: "***",
	}
	reqDto := &ReqCreateDto{
		ReqBase: &ReqBase{
			Code:     "***",
			Password: "***",
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
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
```

###  spring -- cancel
```go
	custDto := &ReqCustomerDto{
		Url: "***",
	}
	reqDto := &ReqCancelDto{
		ReqBase: &ReqBase{
			Code:     "***",
			Password: "***",
		},
		Content: &CancelContentDto{
			BillNo: "7100047050",
		},
	}
	status, respDto, err := Cancel(reqDto, custDto)
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
```

###  spring -- query
```go
	custDto := &ReqCustomerDto{
		Url: "***",
	}
	reqDto := &ReqQueryDto{
		ReqBase: &ReqBase{
			Code:     "***",
			Password: "***",
		},
		Content: &QueryContentDto{
			OrderNo: "1709051572807",
		},
	}
	status, respDto, err := Query(reqDto, custDto)
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
```

###  spring -- get orderno by billno
```go
	custDto := &ReqCustomerDto{
		Url: "***",
	}
	reqDto := &ReqGetOrdernoByBillnoDto{
		ReqBase: &ReqBase{
			Code:     "***",
			Password: "***",
		},
		Content: &QueryGetOrdernoByBillnoDto{
			BillNo: "2003796858",
		},
	}
	status, respDto, err := Query(reqDto, custDto)
	fmt.Println(status, respDto.RespBase, respDto.Info, err)
```