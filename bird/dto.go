package bird

/*
request or response parameter
*/

type ReqQueryDto struct {
	*ReqBase
	RequestData *ReqQueryDataDto `json:"requestData,omitempty"`
}
type ReqCreateDto struct {
	*ReqBase
	RequestData *ReqCreateDataDto `json:"requestData,omitempty"`
}
type ReqSubscribeDto struct {
	*ReqBase
	RequestData *ReqSubscribeDataDto `json:"requestData,omitempty"`
}
type ReqPushDto struct {
	RequestType string          `json:"requestType,omitempty"`
	DataSign    string          `json:"dataSign,omitempty"`
	RequestData *ReqPushDataDto `json:"requestData,omitempty"`
}
type ReqRecognizeDto struct {
	*ReqBase
	RequestData *ReqRecognizeDataDto `json:"requestData,omitempty"`
}
type ReqECreateDto struct {
	*ReqBase
	RequestData *ReqECreateDataDto `json:"requestData,omitempty"`
}
type ReqECancelDto struct {
	*ReqBase
	RequestData *ReqECancelDataDto `json:"requestData,omitempty"`
}
type ReqEAvailableNumDto struct {
	*ReqBase
	RequestData *ReqEAvailableNumDataDto `json:"requestData,omitempty"`
}

type RespQueryDto struct {
	EBusinessId  string `json:"eBusinessID,omitempty"`
	OrderCode    string `json:"orderCode,omitempty"`
	ShipperCode  string `json:"shipperCode,omitempty"`
	LogisticCode string `json:"logisticCode,omitempty"`
	Success      bool   `json:"success,omitempty"`

	Reason string  `json:"reason,omitempty"`
	State  string  `json:"state,omitempty"`
	Traces []Trace `json:"traces,omitempty"`
}
type RespCreateDto struct {
	EBusinessId          string    `json:"eBusinessID,omitempty"`
	Order                *ReqOrder `json:"order,omitempty"`
	Success              bool      `json:"success,omitempty"`
	ResultCode           string    `json:"resultCode,omitempty"`
	Reason               string    `json:"reason,omitempty"`
	UniquerRequestNumber string    `json:"uniquerRequestNumber,omitempty"`
}
type RespSubscribeDto struct {
	EBusinessId           string `json:"eBusinessID,omitempty"`
	UpdateTime            string `json:"updateTime,omitempty"`
	Success               bool   `json:"success,omitempty"`
	Reason                string `json:"reason,omitempty"`
	EstimatedDeliveryTime string `json:"estimatedDeliveryTime,omitempty"`
}
type RespPushDto struct {
	EBusinessId string `json:"eBusinessID,omitempty"`
	UpdateTime  string `json:"updateTime,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Reason      string `json:"reason,omitempty"`
}
type RespRecognizeDto struct {
	EBusinessId  string    `json:"eBusinessID,omitempty"`
	LogisticCode string    `json:"logisticCode,omitempty"`
	Success      bool      `json:"success,omitempty"`
	Code         int       `json:"code,omitempty"`
	Shippers     []Shipper `json:"shippers,omitempty"`
}
type RespECreateDto struct {
	EBusinessId           string     `json:"eBusinessID,omitempty"`
	Order                 *ReqEOrder `json:"order,omitempty"`
	Success               bool       `json:"success,omitempty"`
	ResultCode            string     `json:"resultCode,omitempty"`
	Reason                string     `json:"reason,omitempty"`
	UniquerRequestNumber  string     `json:"uniquerRequestNumber,omitempty"`
	PrintTemplate         string     `json:"printTemplate,omitempty"`
	EstimatedDeliveryTime string     `json:"estimatedDeliveryTime,omitempty"`
	Callback              string     `json:"callback,omitempty"`
	// SubCount              int        `json:"SubCount,omitempty"`
	// SubOrders             string     `json:"SubOrders,omitempty"`
	// SubPrintTemplates     string     `json:"SubPrintTemplates,omitempty"`
	SignBillPrintTemplate string `json:"signBillPrintTemplate"`
	ReceiverSafePhone     string `json:"receiverSafePhone"`
	SenderSafePhone       string `json:"senderSafePhone"`
	DialPage              string `json:"dialPage"`
}
type RespECancelDto struct {
	EBusinessId string `json:"eBusinessID,omitempty"`
	Success     bool   `json:"success,omitempty"`
	ResultCode  string `json:"resultCode,omitempty"`
	Reason      string `json:"reason,omitempty"`
}
type RespEAvailableNumDto struct {
	EBusinessId   string         `json:"eBusinessID,omitempty"`
	Success       bool           `json:"success,omitempty"`
	ResultCode    string         `json:"resultCode,omitempty"`
	Reason        string         `json:"reason,omitempty"`
	EorderBalance *EorderBalance `json:"eorderBalance,omitempty"`
}

/*
dto
*/
type ReqBase struct {
	EBusinessId string `json:"eBusinessID,omitempty"`
	RequestType string `json:"requestType,omitempty"`
	DataSign    string `json:"dataSign,omitempty"`
	DataType    string `json:"dataType,omitempty"`
}

type ReqOrder struct {
	OrderCode    string `json:"orderCode,omitempty"`
	ShipperCode  string `json:"shipperCode,omitempty"`
	LogisticCode string `json:"logisticCode,omitempty"`
}
type ReqEOrder struct {
	OrderCode       string `json:"orderCode,omitempty"`
	ShipperCode     string `json:"shipperCode,omitempty"`
	LogisticCode    string `json:"logisticCode,omitempty"`
	MarkDestination string `json:"markDestination,omitempty"`
	OriginCode      string `json:"originCode,omitempty"`
	OriginName      string `json:"originName,omitempty"`
	DestinatioCode  string `json:"destinatioCode,omitempty"`
	DestinatioName  string `json:"destinatioName,omitempty"`
	SortingCode     string `json:"sortingCode,omitempty"`
	PackageCode     string `json:"packageCode,omitempty"`
}

type ReqCreateDataDto struct {
	WarehouseId      string `json:"warehouseID,omitempty"`
	WarehouseAddress string `json:"warehouseAddress,omitempty"`
	CallBack         string `json:"callBack,omitempty"`
	MemberId         string `json:"memberID,omitempty"`
	ShipperCode      string `json:"shipperCode,omitempty"`

	LogisticCode string `json:"logisticCode,omitempty"`
	OrderCode    string `json:"orderCode,omitempty"`
	PayType      int    `json:"payType,omitempty"`
	MonthCode    string `json:"monthCode,omitempty"`
	ExpType      int    `json:"expType,omitempty"`

	Cost      float64   `json:"cost,omitempty"`
	OtherCost float64   `json:"otherCost,omitempty"`
	Receiver  *Receiver `json:"receiver,omitempty"`
	Sender    *Sender   `json:"sender,omitempty"`
	StartDate string    `json:"startDate,omitempty"`

	EndDate  string  `json:"endDate,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
	Volume   float64 `json:"volume,omitempty"`
	Remark   string  `json:"remark,omitempty"`

	AddServices []AddService `json:"addService,omitempty"`
	Commoditys  []Commodity  `json:"commodity,omitempty"`
}

type AddService struct {
	Name       string `json:"name,omitempty"`
	Value      string `json:"value,omitempty"`
	CustomerId string `json:"customerID,omitempty"`
}

type Commodity struct {
	GoodsName     string  `json:"goodsName,omitempty"`
	GoodsCode     string  `json:"goodsCode,omitempty"`
	Goodsquantity int     `json:"goodsquantity,omitempty"`
	GoodsPrice    float64 `json:"goodsPrice,omitempty"`
	GoodsWeight   float64 `json:"goodsWeight,omitempty"`

	GoodsDesc string  `json:"goodsDesc,omitempty"`
	GoodsVol  float64 `json:"goodsVol,omitempty"`
}

type Receiver struct {
	Company  string `json:"company,omitempty"`
	Name     string `json:"name,omitempty"`
	Tel      string `json:"tel,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	PostCode string `json:"postCode,omitempty"`

	ProvinceName string `json:"provinceName,omitempty"`
	CityName     string `json:"cityName,omitempty"`
	ExpAreaName  string `json:"expAreaName,omitempty"`
	Address      string `json:"address,omitempty"`
}

type Sender struct {
	Company  string `json:"company,omitempty"`
	Name     string `json:"name,omitempty"`
	Tel      string `json:"tel,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	PostCode string `json:"postCode,omitempty"`

	ProvinceName string `json:"provinceName,omitempty"`
	CityName     string `json:"cityName,omitempty"`
	ExpAreaName  string `json:"expAreaName,omitempty"`
	Address      string `json:"address,omitempty"`
}

type ReqQueryDataDto struct {
	OrderCode    string `json:"orderCode,omitempty"`
	ShipperCode  string `json:"shipperCode,omitempty"`
	LogisticCode string `json:"logisticCode,omitempty"`
	IsHandleInfo string `json:"isHandleInfo,omitempty"`
}
type ReqCustomerDto struct {
	Url    string `json:"url,omitempty"`
	ApiKey string `json:"apiKey,omitempty"`
}
type Trace struct {
	AcceptTime    string `json:"acceptTime,omitempty"`
	AcceptStation string `json:"acceptStation,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

//订阅RequestData
type ReqSubscribeDataDto struct {
	CallBack     string `json:"callBack,omitempty"`
	MemberID     string `json:"memberID,omitempty"`
	WareHouseID  string `json:"wareHouseID,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	CustomerPwd  string `json:"customerPwd,omitempty"`

	SendSite     string `json:"sendSite,omitempty"`
	ShipperCode  string `json:"shipperCode,omitempty"`
	LogisticCode string `json:"logisticCode,omitempty"`
	OrderCode    string `json:"orderCode,omitempty"`
	MonthCode    string `json:"monthCode,omitempty"`

	PayType   int       `json:"payType,omitempty"`
	ExpType   string    `json:"expType,omitempty"`
	Cost      float64   `json:"cost,omitempty"`
	OtherCost float64   `json:"otherCost,omitempty"`
	Receiver  *Receiver `json:"receiver,omitempty"`
	Sender    *Sender   `json:"sender,omitempty"`

	IsNotice      int     `json:"isNotice,omitempty"`
	StartDate     string  `json:"startDate,omitempty"`
	EndDate       string  `json:"endDate,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	Volume        float64 `json:"volume,omitempty"`
	Remark        string  `json:"remark,omitempty"`
	IsSendMessage int     `json:"isSendMessage,omitempty"`

	AddServices []AddService `json:"addService,omitempty"`
	Commoditys  []Commodity  `json:"commodity,omitempty"`
}

type ReqPushDataDto struct {
	EBusinessId string        `json:"eBusinessID,omitempty"`
	PushTime    string        `json:"pushTime,omitempty"`
	Count       string        `json:"count,omitempty"`
	Data        []PushDataDto `json:"data,omitempty"`
}
type PushDataDto struct {
	EBusinessId           string  `json:"eBusinessID,omitempty"`
	ShipperCode           string  `json:"shipperCode,omitempty"`
	LogisticCode          string  `json:"logisticCode,omitempty"`
	Success               bool    `json:"success,omitempty"`
	Reason                string  `json:"reason,omitempty"`
	State                 string  `json:"state,omitempty"` //物流状态: 0-无轨迹，1-已揽收，2-在途中，3-签收,4-问题件
	CallBack              string  `json:"callBack,omitempty"`
	Traces                []Trace `json:"traces,omitempty"`
	EstimatedDeliveryTime string  `json:"estimatedDeliveryTime,omitempty"`
}

type ReqRecognizeDataDto struct {
	LogisticCode string `json:"logisticCode,omitempty"`
}
type Shipper struct {
	ShipperCode string `json:"shipperCode,omitempty"`
	ShipperName string `json:"shipperName,omitempty"`
}

type ReqECreateDataDto struct {
	CallBack     string  `json:"callBack,omitempty"`
	MemberId     string  `json:"memberID,omitempty"`
	CustomerName string  `json:"customerName,omitempty"`
	CustomerPwd  string  `json:"customerPwd,omitempty"`
	SendSite     string  `json:"sendSite,omitempty"`
	ShipperCode  string  `json:"shipperCode,omitempty"`
	LogisticCode string  `json:"logisticCode,omitempty"`
	OrderCode    string  `json:"orderCode,omitempty"`
	ThrOrderCode string  `json:"thrOrderCode,omitempty"`
	MonthCode    string  `json:"monthCode,omitempty"`
	PayType      int     `json:"payType,omitempty"`
	ExpType      int     `json:"expType,omitempty"`
	IsNotice     int     `json:"isNotice,omitempty"`
	Cost         float64 `json:"cost,omitempty"`
	OtherCost    float64 `json:"otherCost,omitempty"`

	Receiver *Receiver `json:"receiver,omitempty"`
	Sender   *Sender   `json:"sender,omitempty"`

	StartDate string  `json:"startDate,omitempty"`
	EndDate   string  `json:"endDate,omitempty"`
	Weight    float64 `json:"weight,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	Volume    float64 `json:"volume,omitempty"`
	Remark    string  `json:"remark,omitempty"`

	AddServices []AddService `json:"addService,omitempty"`
	Commoditys  []Commodity  `json:"commodity,omitempty"`

	IsReturnPrintTemplate string `json:"isReturnPrintTemplate,omitempty"`
	IsSendMessage         int    `json:"isSendMessage,omitempty"`
	TemplateSize          string `json:"templateSize,omitempty"`
}

type ReqECancelDataDto struct {
	ShipperCode  string `json:"shipperCode,omitempty"`
	OrderCode    string `json:"orderCode,omitempty"`
	ExpNo        string `json:"expNo,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	CustomerPwd  string `json:"customerPwd,omitempty"`
}

type ReqEAvailableNumDataDto struct {
	ShipperCode  string `json:"shipperCode,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	CustomerPwd  string `json:"customerPwd,omitempty"`
	StationCode  string `json:"stationCode,omitempty"`
	StationName  string `json:"stationName,omitempty"`
}

type EorderBalance struct {
	TotalNum     int `json:"totalNum,omitempty"`
	AvailableNum int `json:"availableNum,omitempty"`
}
