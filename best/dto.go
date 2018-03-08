package best

import (
	"encoding/xml"
)

//request param
type ReqCreateDto struct {
	*ReqBase
	BizData *CreateContentDto
}
type ReqCancelDto struct {
	*ReqBase
	BizData *CancelContentDto
}
type ReqQueryDto struct {
	*ReqBase
	BizData *QueryContentDto
}

type RespQueryDto struct {
	Result     bool           `xml:"result"`
	Errors     *ErrorsDto     `xml:"errors"`
	OrderInfos *OrderInfosDto `xml:"orderInfos"`
}

//dto
type ReqBase struct {
	PartnerID   string `xml:"partnerID"`
	PartnerKey  string `xml:"partnerKey"`
	ServiceType string `xml:"serviceType"`
	Sign        string `xml:"sign"`
}
type RespBase struct {
	Result           bool   `xml:"result"`
	Note             string `xml:"note"`
	ErrorCode        string `xml:"errorCode"`
	ErrorDescription string `xml:"errorDescription"`
}

type CreateContentDto struct {
	XMLName         xml.Name     `xml:"request"`
	CustomerCode    string       `xml:"customerCode"`
	CustomerName    string       `xml:"customerName"`
	ProjectCode     string       `xml:"projectCode"`
	TransportMode   string       `xml:"transportMode"`
	VehicleModel    string       `xml:"vehicleModel"`
	LogisticsCode   string       `xml:"logisticsCode"`
	Remark          string       `xml:"remark"`
	ShippingOrderNo string       `xml:"shippingOrderNo"`
	OrderCode       string       `xml:"orderCode"`
	Type            string       `xml:"type"`
	Value           string       `xml:"value"`
	Sender          *SenderDto   `xml:"sender"`
	Receiver        *ReceiverDto `xml:"receiver"`
	Services        *ServicesDto `xml:"services"`
	Items           *ItemsDto    `xml:"items"`
}

type ReceiverDto struct {
	Name              string `xml:"name"`
	Province          string `xml:"province"`
	City              string `xml:"city"`
	District          string `xml:"district"`
	Address           string `xml:"address"`
	ContactName       string `xml:"contactName"`
	Phone             string `xml:"phone"`
	EarlyDeliveryTime string `xml:"earlyDeliveryTime"`
	LateDeliveryTime  string `xml:"lateDeliveryTime"`
}
type SenderDto struct {
	Name            string `xml:"name"`
	Province        string `xml:"province"`
	City            string `xml:"city"`
	District        string `xml:"district"`
	Address         string `xml:"address"`
	ContactName     string `xml:"contactName"`
	Phone           string `xml:"phone"`
	EarlyPickupTime string `xml:"earlyPickupTime"`
	LatePickupTime  string `xml:"latePickupTime"`
}
type ServiceDto struct {
	ServiceDefinitionCode string       `xml:"serviceDefinitionCode"`
	Remark                string       `xml:"remark"`
	ServiceCodes          *ServicesDto `xml:"serviceCodes"`
}
type ServicesDto struct {
	Service *[]ServiceCodeDto `xml:"service"`
}
type ServiceCodeDto struct {
	ServiceDefinitionCode string `xml:"serviceDefinitionCode"`
	ActualValue           string `xml:"actualValue"`
}

type ItemsDto struct {
	Item *[]ItemDto `xml:"item"`
}
type ItemDto struct {
	LineNo         string  `xml:"lineNo"`
	ItemCode       string  `xml:"itemCode"`
	ItemName       string  `xml:"itemName"`
	ItemCount      int     `xml:"itemCount"`
	PackageCount   int     `xml:"packageCount"`
	PackageUomCode string  `xml:"packageUomCode"`
	Weight         float64 `xml:"weight"`
	Volume         float64 `xml:"volume"`
	DeclaredValue  float64 `xml:"declaredValue"`
	VolumeWeight   float64 `xml:"volumeWeight"`
	Remark         string  `xml:"remark"`
}

type CancelContentDto struct {
	CustomerCode string `xml:"customerCode"`
	OrderCode    string `xml:"orderCode"`
}

type QueryContentDto struct {
	CustomerCode   string `xml:"customerCode"`
	CreateTimeFrom string `xml:"createTimeFrom"`
	CreateTimeTo   string `xml:"createTimeTo"`
	OrderCode      string `xml:"orderCode"`
	ShipmentCode   string `xml:"shipmentCode"`
}

type ReqCustomerDto struct {
	Url string `xml:"url"`
}

type ErrorsDto struct {
	Error *[]ErrorDto `xml:"error"`
}
type ErrorDto struct {
	ErrorCode        string `xml:"errorCode"`
	ErrorDescription string `xml:"errorDescription"`
}
type OrderInfosDto struct {
	OrderInfo *[]OrderInfoDto `xml:"orderInfo"`
}
type OrderInfoDto struct {
	CustomerCode             string        `xml:"customerCode"`
	CustomerName             string        `xml:"customerName"`
	ProjectCode              string        `xml:"projectCode"`
	TmsCode                  string        `xml:"tmsCode"`
	OrderCode                string        `xml:"orderCode"`
	Status                   string        `xml:"status"`
	CurrentStatusDatetime    string        `xml:"currentStatusDatetime"`
	CurrentStatusLocation    string        `xml:"currentStatusLocation"`
	CurrentStatusDescription string        `xml:"currentStatusDescription"`
	CurrentStatusUpdator     string        `xml:"currentStatusUpdator"`
	PortalUrl                string        `xml:"portalUrl"`
	Shippings                *ShippingsDto `xml:"shippings"`
}
type ShippingsDto struct {
	Shipping *[]ShippingDto `xml:"shipping"`
}
type ShippingDto struct {
	LogisticsCode         string     `xml:"logisticsCode"`
	LogisticsName         string     `xml:"logisticsName"`
	ShipmentCode          string     `xml:"shipmentCode"`
	ReferenceShipCode     string     `xml:"referenceShipCode"`
	TransportMode         string     `xml:"transportMode"`
	SourceLocationAddress string     `xml:"sourceLocationAddress"`
	DestLocationAddress   string     `xml:"destLocationAddress"`
	OperateTime           string     `xml:"operateTime"`
	Operator              string     `xml:"operator"`
	DriverName            string     `xml:"driverName"`
	DriverPhone           string     `xml:"driverPhone"`
	OperateDescription    string     `xml:"operateDescription"`
	ShipmentStatus        string     `xml:"shipmentStatus"`
	Cod                   string     `xml:"cod"`
	CodStatus             string     `xml:"codStatus"`
	CodAmount             float64    `xml:"codAmount"`
	GoodsValue            float64    `xml:"goodsValue"`
	CheapAmount           float64    `xml:"cheapAmount"`
	CodPayType            string     `xml:"codPayType"`
	Items                 *SItemsDto `xml:"items"`
	Traces                *TracesDto `xml:"traces"`
}
type SItemsDto struct {
	Item *[]SItemDto `xml:"item"`
}
type SItemDto struct {
	ItemCode       string  `xml:"itemCode"`
	ItemName       string  `xml:"itemName"`
	PackageUomCode string  `xml:"packageUomCode"`
	Count          int64   `xml:"count"`
	Weight         float64 `xml:"weight"`
	Volume         float64 `xml:"volume"`
}
type TracesDto struct {
	Trace *[]TraceDto `xml:"trace"`
}
type TraceDto struct {
	OpTime     string `xml:"opTime"`
	OpStatus   string `xml:"opStatus"`
	Operator   string `xml:"operator"`
	OpDesc     string `xml:"opDesc"`
	OpLocation string `xml:"opLocation"`
}
