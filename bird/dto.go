package bird

/*
request or response parameter
*/

type ReqQueryDto struct {
	*ReqBase
	RequestData *ReqQueryDataDto `json:"RequestData,omitempty"`
}
type ReqCreateDto struct {
	*ReqBase
	RequestData *ReqCreateDataDto `json:"RequestData,omitempty"`
}
type ReqSubscribeDto struct {
	*ReqBase
	RequestData *ReqSubscribeDataDto `json:"RequestData,omitempty"`
}
type ReqPushDto struct {
	RequestType string          `json:"RequestType,omitempty"`
	DataSign    string          `json:"DataSign,omitempty"`
	RequestData *ReqPushDataDto `json:"RequestData,omitempty"`
}
type ReqRecognizeDto struct {
	*ReqBase
	RequestData *ReqRecognizeDataDto `json:"RequestData,omitempty"`
}
type ReqECreateDto struct {
	*ReqBase
	RequestData *ReqECreateDataDto `json:"RequestData,omitempty"`
}
type ReqECancelDto struct {
	*ReqBase
	RequestData *ReqECancelDataDto `json:"RequestData,omitempty"`
}
type ReqEAvailableNumDto struct {
	*ReqBase
	RequestData *ReqEAvailableNumDataDto `json:"RequestData,omitempty"`
}

type RespQueryDto struct {
	EBusinessId  string `json:"EBusinessID,omitempty"`
	OrderCode    string `json:"OrderCode,omitempty"`
	ShipperCode  string `json:"ShipperCode,omitempty"`
	LogisticCode string `json:"LogisticCode,omitempty"`
	Success      bool   `json:"Success,omitempty"`

	Reason string  `json:"Reason,omitempty"`
	State  string  `json:"State,omitempty"`
	Traces []Trace `json:"Traces,omitempty"`
}
type RespCreateDto struct {
	EBusinessId          string    `json:"EBusinessID,omitempty"`
	Order                *ReqOrder `json:"Order,omitempty"`
	Success              bool      `json:"Success,omitempty"`
	ResultCode           string    `json:"ResultCode,omitempty"`
	Reason               string    `json:"Reason,omitempty"`
	UniquerRequestNumber string    `json:"UniquerRequestNumber,omitempty"`
}
type RespSubscribeDto struct {
	EBusinessId           string `json:"EBusinessID,omitempty"`
	UpdateTime            string `json:"UpdateTime,omitempty"`
	Success               bool   `json:"Success,omitempty"`
	Reason                string `json:"Reason,omitempty"`
	EstimatedDeliveryTime string `json:"EstimatedDeliveryTime,omitempty"`
}
type RespPushDto struct {
	EBusinessId string `json:"EBusinessID,omitempty"`
	UpdateTime  string `json:"UpdateTime,omitempty"`
	Success     bool   `json:"Success,omitempty"`
	Reason      string `json:"Reason,omitempty"`
}
type RespRecognizeDto struct {
	EBusinessId  string    `json:"EBusinessID,omitempty"`
	LogisticCode string    `json:"LogisticCode,omitempty"`
	Success      bool      `json:"Success,omitempty"`
	Code         int       `json:"Code,omitempty"`
	Shippers     []Shipper `json:"Shippers,omitempty"`
}
type RespECreateDto struct {
	EBusinessId           string     `json:"EBusinessID,omitempty"`
	Order                 *ReqEOrder `json:"Order,omitempty"`
	Success               bool       `json:"Success,omitempty"`
	ResultCode            string     `json:"ResultCode,omitempty"`
	Reason                string     `json:"Reason,omitempty"`
	UniquerRequestNumber  string     `json:"UniquerRequestNumber,omitempty"`
	PrintTemplate         string     `json:"PrintTemplate,omitempty"`
	EstimatedDeliveryTime string     `json:"EstimatedDeliveryTime,omitempty"`
	Callback              string     `json:"Callback,omitempty"`
	// SubCount              int        `json:"SubCount,omitempty"`
	// SubOrders             string     `json:"SubOrders,omitempty"`
	// SubPrintTemplates     string     `json:"SubPrintTemplates,omitempty"`
	SignBillPrintTemplate string `json:"SignBillPrintTemplate"`
	ReceiverSafePhone     string `json:"ReceiverSafePhone"`
	SenderSafePhone       string `json:"SenderSafePhone"`
	DialPage              string `json:"DialPage"`
}
type RespECancelDto struct {
	EBusinessId string `json:"EBusinessID,omitempty"`
	Success     bool   `json:"Success,omitempty"`
	ResultCode  string `json:"ResultCode,omitempty"`
	Reason      string `json:"Reason,omitempty"`
}
type RespEAvailableNumDto struct {
	EBusinessId   string         `json:"EBusinessID,omitempty"`
	Success       bool           `json:"Success,omitempty"`
	ResultCode    string         `json:"ResultCode,omitempty"`
	Reason        string         `json:"Reason,omitempty"`
	EorderBalance *EorderBalance `json:"EorderBalance,omitempty"`
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

type ReqOrder struct {
	OrderCode    string `json:"OrderCode,omitempty"`
	ShipperCode  string `json:"ShipperCode,omitempty"`
	LogisticCode string `json:"LogisticCode,omitempty"`
}
type ReqEOrder struct {
	OrderCode       string `json:"OrderCode,omitempty"`
	ShipperCode     string `json:"ShipperCode,omitempty"`
	LogisticCode    string `json:"LogisticCode,omitempty"`
	MarkDestination string `json:"MarkDestination,omitempty"`
	OriginCode      string `json:"OriginCode,omitempty"`
	OriginName      string `json:"OriginName,omitempty"`
	DestinatioCode  string `json:"DestinatioCode,omitempty"`
	DestinatioName  string `json:"DestinatioName,omitempty"`
	SortingCode     string `json:"SortingCode,omitempty"`
	PackageCode     string `json:"PackageCode,omitempty"`
}

type ReqCreateDataDto struct {
	WarehouseId      string `json:"WarehouseID,omitempty"`
	WarehouseAddress string `json:"WarehouseAddress,omitempty"`
	CallBack         string `json:"CallBack,omitempty"`
	MemberId         string `json:"MemberID,omitempty"`
	ShipperCode      string `json:"ShipperCode,omitempty"`

	LogisticCode string `json:"LogisticCode,omitempty"`
	OrderCode    string `json:"OrderCode,omitempty"`
	PayType      int    `json:"PayType,omitempty"`
	MonthCode    string `json:"MonthCode,omitempty"`
	ExpType      int    `json:"ExpType,omitempty"`

	Cost      float64   `json:"Cost,omitempty"`
	OtherCost float64   `json:"OtherCost,omitempty"`
	Receiver  *Receiver `json:"Receiver,omitempty"`
	Sender    *Sender   `json:"Sender,omitempty"`
	StartDate string    `json:"StartDate,omitempty"`

	EndDate  string  `json:"EndDate,omitempty"`
	Weight   float64 `json:"Weight,omitempty"`
	Quantity int     `json:"Quantity,omitempty"`
	Volume   float64 `json:"Volume,omitempty"`
	Remark   string  `json:"Remark,omitempty"`

	AddServices []AddService `json:"AddService,omitempty"`
	Commoditys  []Commodity  `json:"Commodity,omitempty"`
}

type AddService struct {
	Name       string `json:"Name,omitempty"`
	Value      string `json:"Value,omitempty"`
	CustomerId string `json:"CustomerID,omitempty"`
}

type Commodity struct {
	GoodsName     string  `json:"GoodsName,omitempty"`
	GoodsCode     string  `json:"GoodsCode,omitempty"`
	Goodsquantity int     `json:"Goodsquantity,omitempty"`
	GoodsPrice    float64 `json:"GoodsPrice,omitempty"`
	GoodsWeight   float64 `json:"GoodsWeight,omitempty"`

	GoodsDesc string  `json:"GoodsDesc,omitempty"`
	GoodsVol  float64 `json:"GoodsVol,omitempty"`
}

type Receiver struct {
	Company  string `json:"Company,omitempty"`
	Name     string `json:"Name,omitempty"`
	Tel      string `json:"Tel,omitempty"`
	Mobile   string `json:"Mobile,omitempty"`
	PostCode string `json:"PostCode,omitempty"`

	ProvinceName string `json:"ProvinceName,omitempty"`
	CityName     string `json:"CityName,omitempty"`
	ExpAreaName  string `json:"ExpAreaName,omitempty"`
	Address      string `json:"Address,omitempty"`
}

type Sender struct {
	Company  string `json:"Company,omitempty"`
	Name     string `json:"Name,omitempty"`
	Tel      string `json:"Tel,omitempty"`
	Mobile   string `json:"Mobile,omitempty"`
	PostCode string `json:"PostCode,omitempty"`

	ProvinceName string `json:"ProvinceName,omitempty"`
	CityName     string `json:"CityName,omitempty"`
	ExpAreaName  string `json:"ExpAreaName,omitempty"`
	Address      string `json:"Address,omitempty"`
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
	AcceptTime    string `json:"AcceptTime,omitempty"`
	AcceptStation string `json:"AcceptStation,omitempty"`
	Remark        string `json:"Remark,omitempty"`
}

//订阅RequestData
type ReqSubscribeDataDto struct {
	CallBack     string `json:"CallBack,omitempty"`
	MemberID     string `json:"MemberID,omitempty"`
	WareHouseID  string `json:"WareHouseID,omitempty"`
	CustomerName string `json:"CustomerName,omitempty"`
	CustomerPwd  string `json:"CustomerPwd,omitempty"`

	SendSite     string `json:"SendSite,omitempty"`
	ShipperCode  string `json:"ShipperCode,omitempty"`
	LogisticCode string `json:"LogisticCode,omitempty"`
	OrderCode    string `json:"OrderCode,omitempty"`
	MonthCode    string `json:"MonthCode,omitempty"`

	PayType   int       `json:"PayType,omitempty"`
	ExpType   string    `json:"ExpType,omitempty"`
	Cost      float64   `json:"Cost,omitempty"`
	OtherCost float64   `json:"OtherCost,omitempty"`
	Receiver  *Receiver `json:"Receiver,omitempty"`
	Sender    *Sender   `json:"Sender,omitempty"`

	IsNotice      int     `json:"IsNotice,omitempty"`
	StartDate     string  `json:"StartDate,omitempty"`
	EndDate       string  `json:"EndDate,omitempty"`
	Weight        float64 `json:"Weight,omitempty"`
	Quantity      int     `json:"Quantity,omitempty"`
	Volume        float64 `json:"Volume,omitempty"`
	Remark        string  `json:"Remark,omitempty"`
	IsSendMessage int     `json:"IsSendMessage,omitempty"`

	AddServices []AddService `json:"AddService,omitempty"`
	Commoditys  []Commodity  `json:"Commodity,omitempty"`
}

type ReqPushDataDto struct {
	EBusinessId string        `json:"EBusinessID,omitempty"`
	PushTime    string        `json:"PushTime,omitempty"`
	Count       string        `json:"Count,omitempty"`
	Data        []PushDataDto `json:"Data,omitempty"`
}
type PushDataDto struct {
	EBusinessId           string  `json:"EBusinessID,omitempty"`
	ShipperCode           string  `json:"ShipperCode,omitempty"`
	LogisticCode          string  `json:"LogisticCode,omitempty"`
	Success               bool    `json:"Success,omitempty"`
	Reason                string  `json:"Reason,omitempty"`
	State                 string  `json:"State,omitempty"` //物流状态: 0-无轨迹，1-已揽收，2-在途中，3-签收,4-问题件
	CallBack              string  `json:"CallBack,omitempty"`
	Traces                []Trace `json:"Traces,omitempty"`
	EstimatedDeliveryTime string  `json:"EstimatedDeliveryTime,omitempty"`
}

type ReqRecognizeDataDto struct {
	LogisticCode string `json:"LogisticCode,omitempty"`
}
type Shipper struct {
	ShipperCode string `json:"ShipperCode,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
}

type ReqECreateDataDto struct {
	CallBack     string  `json:"CallBack,omitempty"`
	MemberId     string  `json:"MemberID,omitempty"`
	CustomerName string  `json:"CustomerName,omitempty"`
	CustomerPwd  string  `json:"CustomerPwd,omitempty"`
	SendSite     string  `json:"SendSite,omitempty"`
	ShipperCode  string  `json:"ShipperCode,omitempty"`
	LogisticCode string  `json:"LogisticCode,omitempty"`
	OrderCode    string  `json:"OrderCode,omitempty"`
	ThrOrderCode string  `json:"ThrOrderCode,omitempty"`
	MonthCode    string  `json:"MonthCode,omitempty"`
	PayType      int     `json:"PayType,omitempty"`
	ExpType      int     `json:"ExpType,omitempty"`
	IsNotice     int     `json:"IsNotice,omitempty"`
	Cost         float64 `json:"Cost,omitempty"`
	OtherCost    float64 `json:"OtherCost,omitempty"`

	Receiver *Receiver `json:"Receiver,omitempty"`
	Sender   *Sender   `json:"Sender,omitempty"`

	StartDate string  `json:"StartDate,omitempty"`
	EndDate   string  `json:"EndDate,omitempty"`
	Weight    float64 `json:"Weight,omitempty"`
	Quantity  int     `json:"Quantity,omitempty"`
	Volume    float64 `json:"Volume,omitempty"`
	Remark    string  `json:"Remark,omitempty"`

	AddServices []AddService `json:"AddService,omitempty"`
	Commoditys  []Commodity  `json:"Commodity,omitempty"`

	IsReturnPrintTemplate string `json:"IsReturnPrintTemplate,omitempty"`
	IsSendMessage         int    `json:"IsSendMessage,omitempty"`
	TemplateSize          string `json:"TemplateSize,omitempty"`
}

type ReqECancelDataDto struct {
	ShipperCode  string `json:"ShipperCode,omitempty"`
	OrderCode    string `json:"OrderCode,omitempty"`
	ExpNo        string `json:"ExpNo,omitempty"`
	CustomerName string `json:"CustomerName,omitempty"`
	CustomerPwd  string `json:"CustomerPwd,omitempty"`
}

type ReqEAvailableNumDataDto struct {
	ShipperCode  string `json:"ShipperCode,omitempty"`
	CustomerName string `json:"CustomerName,omitempty"`
	CustomerPwd  string `json:"CustomerPwd,omitempty"`
	StationCode  string `json:"StationCode,omitempty"`
	StationName  string `json:"StationName,omitempty"`
}

type EorderBalance struct {
	TotalNum     int `json:"TotalNum,omitempty"`
	AvailableNum int `json:"AvailableNum,omitempty"`
}
