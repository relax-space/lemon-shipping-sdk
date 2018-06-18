package bird

import (
	"time"
)

/*
request or response parameter
*/

type ReqQueryDto struct {
	*ReqBase
	RequestData *ReqQueryDataDto `json:"RequestData,omitempty"`
}

type RespQueryDto struct {
	EBusinessId  string `json:"EBusinessID,omitempty"`
	OrderCode    string `json:"OrderCode,omitempty"`
	ShipperCode  string `json:"ShipperCode,omitempty"`
	LogisticCode string `json:"LogisticCode,omitempty"`
	Success      bool   `json:"Success,omitempty"`

	Reason string  `json:"Reason,omitempty"`
	State  string  `json:"State,omitempty"`
	Traces []Trace `json:"Traces>Trace,omitempty"`
}

/*
dto
*/

type ReqBase struct {
	EBusinessId string `json:"EBusinessID,omitempty"`
	RequestType string `json:"RequestType,omitempty"`
	DataSign    string `json:"DataSign,omitempty"`
	DataType    string `json:"DataType,omitempty"`
}

type ReqQueryDataDto struct {
	OrderCode    string `json:"OrderCode,omitempty"`
	ShipperCode  string `json:"ShipperCode,omitempty"`
	LogisticCode string `json:"LogisticCode,omitempty"`
	IsHandleInfo string `json:"IsHandleInfo,omitempty"`
}
type ReqCustomerDto struct {
	Url    string `json:"url,omitempty"`
	ApiKey string `json:"ApiKey,omitempty"`
}
type Trace struct {
	AcceptTime    time.Time `json:"AcceptTime,omitempty"`
	AcceptStation string    `json:"AcceptStation,omitempty"`
	Remark        string    `json:"Remark,omitempty"`
}
