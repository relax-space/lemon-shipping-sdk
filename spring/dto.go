package spring

import "github.com/pangpanglabs/goutils/behaviorlog"

/*
request param
*/
type ReqCreateDto struct {
	*ReqBase
	Content *CreateContentDto `json:"content"`
}

type ReqCancelDto struct {
	*ReqBase
	Content *CancelContentDto `json:"content"`
}

type ReqQueryDto struct {
	*ReqBase
	Content *QueryContentDto `json:"content"`
}

type ReqGetOrdernoByBillnoDto struct {
	*ReqBase
	Content *GetOrdernoByBillnoContentDto `json:"content"`
}

type RespCreateDto struct {
	*RespBase
	Info *CreateInfoDto `json:"info"`
}

type RespCancelDto struct {
	*RespBase
	Info []string `json:"info"`
}

type RespQueryDto struct {
	*RespBase
	Info []QueryInfoDto `json:"info"`
}

type RespGetOrdernoByBillnoDto struct {
	*RespBase
	Info *GetOrdernoByBillnoInfoDto `json:"info"`
}

type RespCommonDto struct {
	*RespBase
	Info interface{} `json:"info"`
}

/*
dto
*/

type ReqBase struct {
	Code       string `json:"code"`
	Password   string `json:"password"`
	Datetime   string `json:"datetime"`
	VerifyCode string `json:"verify_code"`
	Func       string `json:"func"`
}
type RespBase struct {
	Result    string `json:"result"`
	ErrorCode string `json:"error_code"`
	Remark    string `json:"remark"`
}

type CreateInfoDto struct {
	OrderNo string `json:"orderno"`
	BillNo  string `json:"billno"`
}

type CreateContentDto struct {
	BillNo   string       `json:"billno"`
	Quantity int64        `json:"quantity"`
	Package  int64        `json:"package"`
	Volume   float64      `json:"volume"`
	Receiver *ReceiverDto `json:"receiver"`
	Sender   *SenderDto   `json:"sender"`
}
type ReceiverDto struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Company string `json:"company"`
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Mobile  string `json:"mobile"`
}
type SenderDto struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Company string `json:"company"`
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Mobile  string `json:"mobile"`
}
type CancelContentDto struct {
	OrderNo string `json:"orderno"`
	BillNo  string `json:"billno"`
}
type QueryContentDto struct {
	OrderNo string `json:"orderno"`
	BillNo  string `json:"billno"`
}
type QueryInfoDto struct {
	Time      string `json:"time"`
	BillNo    string `json:"billno"`
	StateInfo string `json:"state_info"`
	State     string `json:"state"`
}

type GetOrdernoByBillnoContentDto struct {
	BillNo string `json:"billno"`
}
type GetOrdernoByBillnoInfoDto struct {
	OrderNo string `json:"orderno"`
	BillNo  string `json:"billno"`
}

type ReqCustomerDto struct {
	Url        string `json:"url,omitempty"`
	LogContext *behaviorlog.LogContext
}
