package best

import (
	"encoding/xml"
)

//request param
type ReqCreateDto struct {
	*ReqBase
	BizData *CreateContentDto `xml:"bizData" json:"bizData,omitempty"`
}
type ReqCancelDto struct {
	*ReqBase
	BizData *CancelContentDto `xml:"bizData" json:"bizData,omitempty"`
}
type ReqQueryDto struct {
	*ReqBase
	BizData *QueryContentDto `xml:"bizData" json:"bizData,omitempty"`
}

type RespQueryDto struct {
	Result     bool           `xml:"result" json:"result"`
	Errors     []ErrorDto     `xml:"errors>error" json:"errors"`
	OrderInfos []OrderInfoDto `xml:"orderInfos>orderInfo" json:"orderInfos"`
}

type ReqPushDto struct {
	CustomerCode             string        `xml:"customerCode" json:"customerCode"`
	CustomerName             string        `xml:"customerName" json:"customerName"`
	ProjectCode              string        `xml:"projectCode" json:"projectCode"`
	OrderCode                string        `xml:"orderCode" json:"orderCode"`
	Status                   string        `xml:"status" json:"status"`
	CurrentStatusDatetime    string        `xml:"currentStatusDatetime" json:"currentStatusDatetime"`
	CurrentStatusLocation    string        `xml:"currentStatusLocation" json:"currentStatusLocation"`
	CurrentStatusDescription string        `xml:"currentStatusDescription" json:"currentStatusDescription"`
	CurrentStatusUpdator     string        `xml:"currentStatusUpdator" json:"currentStatusUpdator"`
	PortalUrl                string        `xml:"portalUrl" json:"portalUrl"`
	Shippings                []ShippingDto `xml:"shippings>shipping" json:"shippings"`
}

//dto
type ReqBase struct {
	PartnerID   string `xml:"partnerID" json:"partnerID,omitempty"`
	PartnerKey  string `xml:"partnerKey" json:"partnerKey,omitempty"`
	ServiceType string `xml:"serviceType" json:"serviceType,omitempty"`
	Sign        string `xml:"sign" json:"sign,omitempty"`
}
type RespBase struct {
	Result           bool   `xml:"result" json:"result"`
	Note             string `xml:"note" json:"note"`
	ErrorCode        string `xml:"errorCode" json:"errorCode"`
	ErrorDescription string `xml:"errorDescription" json:"errorDescription"`
}

type CreateContentDto struct {
	XMLName         xml.Name     `xml:"request"`
	CustomerCode    string       `xml:"customerCode" json:"customerCode,omitempty"`
	CustomerName    string       `xml:"customerName" json:"customerName,omitempty"`
	ProjectCode     string       `xml:"projectCode" json:"projectCode,omitempty"`
	TransportMode   string       `xml:"transportMode" json:"transportMode,omitempty"`
	VehicleModel    string       `xml:"vehicleModel" json:"vehicleModel,omitempty"`
	LogisticsCode   string       `xml:"logisticsCode" json:"logisticsCode,omitempty"`
	Remark          string       `xml:"remark" json:"remark,omitempty"`
	ShippingOrderNo string       `xml:"shippingOrderNo" json:"shippingOrderNo,omitempty"`
	OrderCode       string       `xml:"orderCode" json:"orderCode,omitempty"`
	Type            string       `xml:"type" json:"type,omitempty"`
	Value           string       `xml:"value" json:"value,omitempty"`
	Sender          *SenderDto   `xml:"sender" json:"sender,omitempty"`
	Receiver        *ReceiverDto `xml:"receiver" json:"receiver,omitempty"`
	Services        []ServiceDto `xml:"services>service" json:"services,omitempty"`
	Items           []ItemDto    `xml:"items>item" json:"items,omitempty"`
}

type ReceiverDto struct {
	Name              string `xml:"name" json:"name,omitempty"`
	Province          string `xml:"province" json:"province,omitempty"`
	City              string `xml:"city" json:"city,omitempty"`
	District          string `xml:"district" json:"district,omitempty"`
	Address           string `xml:"address" json:"address,omitempty"`
	ContactName       string `xml:"contactName" json:"contactName,omitempty"`
	Phone             string `xml:"phone" json:"phone,omitempty"`
	EarlyDeliveryTime string `xml:"earlyDeliveryTime" json:"earlyDeliveryTime,omitempty"`
	LateDeliveryTime  string `xml:"lateDeliveryTime" json:"lateDeliveryTime,omitempty"`
}
type SenderDto struct {
	Name            string `xml:"name" json:"name,omitempty"`
	Province        string `xml:"province" json:"province,omitempty"`
	City            string `xml:"city" json:"city,omitempty"`
	District        string `xml:"district" json:"district,omitempty"`
	Address         string `xml:"address" json:"address,omitempty"`
	ContactName     string `xml:"contactName" json:"contactName,omitempty"`
	Phone           string `xml:"phone" json:"phone,omitempty"`
	EarlyPickupTime string `xml:"earlyPickupTime" json:"earlyPickupTime,omitempty"`
	LatePickupTime  string `xml:"latePickupTime" json:"latePickupTime,omitempty"`
}
type ServiceDto struct {
	ServiceDefinitionCode string           `xml:"serviceDefinitionCode" json:"serviceDefinitionCode,omitempty"`
	Remark                string           `xml:"remark" json:"remark,omitempty"`
	ServiceCodes          []ServiceCodeDto `xml:"serviceCodes>serviceCode" json:"serviceCodes,omitempty"`
}
type ServiceCodeDto struct {
	ServiceDefParameterCode string `xml:"serviceDefParameterCode" json:"serviceDefParameterCode,omitempty"`
	ActualValue             string `xml:"actualValue" json:"actualValue,omitempty"`
}
type ItemDto struct {
	LineNo         string  `xml:"lineNo" json:"lineNo,omitempty"`
	ItemCode       string  `xml:"itemCode" json:"itemCode,omitempty"`
	ItemName       string  `xml:"itemName" json:"itemName,omitempty"`
	ItemCount      int     `xml:"itemCount" json:"itemCount,omitempty"`
	PackageCount   int     `xml:"packageCount" json:"packageCount,omitempty"`
	PackageUomCode string  `xml:"packageUomCode" json:"packageUomCode,omitempty"`
	Weight         float64 `xml:"weight" json:"weight,omitempty"`
	Volume         float64 `xml:"volume" json:"volume,omitempty"`
	DeclaredValue  float64 `xml:"declaredValue" json:"declaredValue,omitempty"`
	VolumeWeight   float64 `xml:"volumeWeight" json:"volumeWeight,omitempty"`
	Remark         string  `xml:"remark" json:"remark,omitempty"`
}

type CancelContentDto struct {
	CustomerCode string `xml:"customerCode" json:"customerCode,omitempty"`
	OrderCode    string `xml:"orderCode" json:"orderCode,omitempty"`
}

type QueryContentDto struct {
	CustomerCode   string `xml:"customerCode" json:"customerCode,omitempty"`
	CreateTimeFrom string `xml:"createTimeFrom" json:"createTimeFrom,omitempty"`
	CreateTimeTo   string `xml:"createTimeTo" json:"createTimeTo,omitempty"`
	OrderCode      string `xml:"orderCode" json:"orderCode,omitempty"`
	ShipmentCode   string `xml:"shipmentCode" json:"shipmentCode,omitempty"`
}

type ReqCustomerDto struct {
	Url string `xml:"url" json:"url,omitempty"`
}

type ErrorDto struct {
	ErrorCode        string `xml:"errorCode" json:"errorCode"`
	ErrorDescription string `xml:"errorDescription" json:"errorDescription"`
}
type OrderInfoDto struct {
	CustomerCode             string        `xml:"customerCode" json:"customerCode"`
	CustomerName             string        `xml:"customerName" json:"customerName"`
	ProjectCode              string        `xml:"projectCode" json:"projectCode"`
	TmsCode                  string        `xml:"tmsCode" json:"tmsCode"`
	OrderCode                string        `xml:"orderCode" json:"orderCode"`
	Status                   string        `xml:"status" json:"status"`
	CurrentStatusDatetime    string        `xml:"currentStatusDatetime" json:"currentStatusDatetime"`
	CurrentStatusLocation    string        `xml:"currentStatusLocation" json:"currentStatusLocation"`
	CurrentStatusDescription string        `xml:"currentStatusDescription" json:"currentStatusDescription"`
	CurrentStatusUpdator     string        `xml:"currentStatusUpdator" json:"currentStatusUpdator"`
	PortalUrl                string        `xml:"portalUrl" json:"portalUrl"`
	Shippings                []ShippingDto `xml:"shippings>shipping" json:"shippings"`
}

type ShippingDto struct {
	LogisticsCode         string         `xml:"logisticsCode" json:"logisticsCode"`
	LogisticsName         string         `xml:"logisticsName" json:"logisticsName"`
	ShipmentCode          string         `xml:"shipmentCode" json:"shipmentCode"`
	ReferenceShipCode     string         `xml:"referenceShipCode" json:"referenceShipCode"`
	TransportMode         string         `xml:"transportMode" json:"transportMode"`
	SourceLocationAddress string         `xml:"sourceLocationAddress" json:"sourceLocationAddress"`
	DestLocationAddress   string         `xml:"destLocationAddress" json:"destLocationAddress"`
	OperateTime           string         `xml:"operateTime" json:"operateTime"`
	Operator              string         `xml:"operator" json:"operator"`
	DriverName            string         `xml:"driverName" json:"driverName"`
	DriverPhone           string         `xml:"driverPhone" json:"driverPhone"`
	OperateDescription    string         `xml:"operateDescription" json:"operateDescription"`
	ShipmentStatus        string         `xml:"shipmentStatus" json:"shipmentStatus"`
	Cod                   string         `xml:"cod" json:"cod"`
	CodStatus             string         `xml:"codStatus" json:"codStatus"`
	CodAmount             float64        `xml:"codAmount" json:"codAmount"`
	GoodsValue            float64        `xml:"goodsValue" json:"goodsValue"`
	CheapAmount           float64        `xml:"cheapAmount" json:"cheapAmount"`
	CodPayType            string         `xml:"codPayType" json:"codPayType"`
	Items                 []QueryItemDto `xml:"items>item" json:"items"`
	Traces                []TraceDto     `xml:"traces>trace" json:"traces"`
}

type QueryItemDto struct {
	ItemCode       string  `xml:"itemCode" json:"itemCode"`
	ItemName       string  `xml:"itemName" json:"itemName"`
	PackageUomCode string  `xml:"packageUomCode" json:"packageUomCode"`
	Count          int64   `xml:"count" json:"count"`
	Weight         float64 `xml:"weight" json:"weight"`
	Volume         float64 `xml:"volume" json:"volume"`
}
type TraceDto struct {
	OpTime     string `xml:"opTime" json:"opTime"`
	OpStatus   string `xml:"opStatus" json:"opStatus"`
	Operator   string `xml:"operator" json:"operator"`
	OpDesc     string `xml:"opDesc" json:"opDesc"`
	OpLocation string `xml:"opLocation" json:"opLocation"`
}
