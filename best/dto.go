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
	Result     bool           `xml:"result" json:"result"`
	Errors     *ErrorsDto     `xml:"errors" json:"errors"`
	OrderInfos *OrderInfosDto `xml:"orderInfos" json:"order_infos"`
}

//dto
type ReqBase struct {
	PartnerID   string `xml:"partnerID" json:"partner_id"`
	PartnerKey  string `xml:"partnerKey" json:"partner_key"`
	ServiceType string `xml:"serviceType" json:"service_type"`
	Sign        string `xml:"sign" json:"sign"`
}
type RespBase struct {
	Result           bool   `xml:"result" json:"result"`
	Note             string `xml:"note" json:"note"`
	ErrorCode        string `xml:"errorCode" json:"error_code"`
	ErrorDescription string `xml:"errorDescription" json:"error_description"`
}

type CreateContentDto struct {
	XMLName         xml.Name     `xml:"request"`
	CustomerCode    string       `xml:"customerCode" json:"customer_code"`
	CustomerName    string       `xml:"customerName" json:"customer_name"`
	ProjectCode     string       `xml:"projectCode" json:"project_code"`
	TransportMode   string       `xml:"transportMode" json:"transport_mode"`
	VehicleModel    string       `xml:"vehicleModel" json:"vehicle_model"`
	LogisticsCode   string       `xml:"logisticsCode" json:"logistics_code"`
	Remark          string       `xml:"remark" json:"remark"`
	ShippingOrderNo string       `xml:"shippingOrderNo" json:"shipping_order_no"`
	OrderCode       string       `xml:"orderCode" json:"order_code"`
	Type            string       `xml:"type" json:"type"`
	Value           string       `xml:"value" json:"value"`
	Sender          *SenderDto   `xml:"sender" json:"sender"`
	Receiver        *ReceiverDto `xml:"receiver" json:"receiver"`
	Services        *ServicesDto `xml:"services" json:"services"`
	Items           *ItemsDto    `xml:"items" json:"items"`
}

type ReceiverDto struct {
	Name              string `xml:"name" json:"name"`
	Province          string `xml:"province" json:"province"`
	City              string `xml:"city" json:"city"`
	District          string `xml:"district" json:"district"`
	Address           string `xml:"address" json:"address"`
	ContactName       string `xml:"contactName" json:"contact_name"`
	Phone             string `xml:"phone" json:"phone"`
	EarlyDeliveryTime string `xml:"earlyDeliveryTime" json:"early_delivery_time"`
	LateDeliveryTime  string `xml:"lateDeliveryTime" json:"late_delivery_time"`
}
type SenderDto struct {
	Name            string `xml:"name" json:"name"`
	Province        string `xml:"province" json:"province"`
	City            string `xml:"city" json:"city"`
	District        string `xml:"district" json:"district"`
	Address         string `xml:"address" json:"address"`
	ContactName     string `xml:"contactName" json:"contact_name"`
	Phone           string `xml:"phone" json:"phone"`
	EarlyPickupTime string `xml:"earlyPickupTime" json:"early_pickup_time"`
	LatePickupTime  string `xml:"latePickupTime" json:"late_pickup_time"`
}
type ServiceDto struct {
	ServiceDefinitionCode string       `xml:"serviceDefinitionCode" json:"service_definition_code"`
	Remark                string       `xml:"remark" json:"remark"`
	ServiceCodes          *ServicesDto `xml:"serviceCodes" json:"service_codes"`
}
type ServicesDto struct {
	Service *[]ServiceCodeDto `xml:"service" json:"service"`
}
type ServiceCodeDto struct {
	ServiceDefinitionCode string `xml:"serviceDefinitionCode" json:"service_definition_code"`
	ActualValue           string `xml:"actualValue" json:"actual_value"`
}

type ItemsDto struct {
	Item *[]ItemDto `xml:"item" json:"item"`
}
type ItemDto struct {
	LineNo         string  `xml:"lineNo" json:"line_no"`
	ItemCode       string  `xml:"itemCode" json:"item_code"`
	ItemName       string  `xml:"itemName" json:"item_name"`
	ItemCount      int     `xml:"itemCount" json:"item_count"`
	PackageCount   int     `xml:"packageCount" json:"package_count"`
	PackageUomCode string  `xml:"packageUomCode" json:"package_uom_code"`
	Weight         float64 `xml:"weight" json:"weight"`
	Volume         float64 `xml:"volume" json:"volume"`
	DeclaredValue  float64 `xml:"declaredValue" json:"declared_value"`
	VolumeWeight   float64 `xml:"volumeWeight" json:"volume_weight"`
	Remark         string  `xml:"remark" json:"remark"`
}

type CancelContentDto struct {
	CustomerCode string `xml:"customerCode" json:"customer_code"`
	OrderCode    string `xml:"orderCode" json:"order_code"`
}

type QueryContentDto struct {
	CustomerCode   string `xml:"customerCode" json:"customer_code"`
	CreateTimeFrom string `xml:"createTimeFrom" json:"create_time_from"`
	CreateTimeTo   string `xml:"createTimeTo" json:"create_time_to"`
	OrderCode      string `xml:"orderCode" json:"order_code"`
	ShipmentCode   string `xml:"shipmentCode" json:"shipment_code"`
}

type ReqCustomerDto struct {
	Url string `xml:"url" json:"url"`
}

type ErrorsDto struct {
	Error *[]ErrorDto `xml:"error" json:"error"`
}
type ErrorDto struct {
	ErrorCode        string `xml:"errorCode" json:"error_code"`
	ErrorDescription string `xml:"errorDescription" json:"error_description"`
}
type OrderInfosDto struct {
	OrderInfo *[]OrderInfoDto `xml:"orderInfo" json:"order_info"`
}
type OrderInfoDto struct {
	CustomerCode             string        `xml:"customerCode" json:"customer_code"`
	CustomerName             string        `xml:"customerName" json:"customer_name"`
	ProjectCode              string        `xml:"projectCode" json:"project_code"`
	TmsCode                  string        `xml:"tmsCode" json:"tms_code"`
	OrderCode                string        `xml:"orderCode" json:"order_code"`
	Status                   string        `xml:"status" json:"status"`
	CurrentStatusDatetime    string        `xml:"currentStatusDatetime" json:"current_status_datetime"`
	CurrentStatusLocation    string        `xml:"currentStatusLocation" json:"current_status_location"`
	CurrentStatusDescription string        `xml:"currentStatusDescription" json:"current_status_description"`
	CurrentStatusUpdator     string        `xml:"currentStatusUpdator" json:"current_status_updator"`
	PortalUrl                string        `xml:"portalUrl" json:"portal_url"`
	Shippings                *ShippingsDto `xml:"shippings" json:"shippings"`
}
type ShippingsDto struct {
	Shipping *[]ShippingDto `xml:"shipping" json:"shipping"`
}
type ShippingDto struct {
	LogisticsCode         string     `xml:"logisticsCode" json:"logistics_code"`
	LogisticsName         string     `xml:"logisticsName" json:"logistics_name"`
	ShipmentCode          string     `xml:"shipmentCode" json:"shipment_code"`
	ReferenceShipCode     string     `xml:"referenceShipCode" json:"reference_ship_code"`
	TransportMode         string     `xml:"transportMode" json:"transport_mode"`
	SourceLocationAddress string     `xml:"sourceLocationAddress" json:"source_location_address"`
	DestLocationAddress   string     `xml:"destLocationAddress" json:"dest_location_address"`
	OperateTime           string     `xml:"operateTime" json:"operate_time"`
	Operator              string     `xml:"operator" json:"operator"`
	DriverName            string     `xml:"driverName" json:"driver_name"`
	DriverPhone           string     `xml:"driverPhone" json:"driver_phone"`
	OperateDescription    string     `xml:"operateDescription" json:"operate_description"`
	ShipmentStatus        string     `xml:"shipmentStatus" json:"shipment_status"`
	Cod                   string     `xml:"cod" json:"cod"`
	CodStatus             string     `xml:"codStatus" json:"cod_status"`
	CodAmount             float64    `xml:"codAmount" json:"cod_amount"`
	GoodsValue            float64    `xml:"goodsValue" json:"goods_value"`
	CheapAmount           float64    `xml:"cheapAmount" json:"cheap_amount"`
	CodPayType            string     `xml:"codPayType" json:"cod_pay_type"`
	Items                 *SItemsDto `xml:"items" json:"items"`
	Traces                *TracesDto `xml:"traces" json:"traces"`
}
type SItemsDto struct {
	Item *[]SItemDto `xml:"item" json:"item"`
}
type SItemDto struct {
	ItemCode       string  `xml:"itemCode" json:"item_code"`
	ItemName       string  `xml:"itemName" json:"item_name"`
	PackageUomCode string  `xml:"packageUomCode" json:"package_uom_code"`
	Count          int64   `xml:"count" json:"count"`
	Weight         float64 `xml:"weight" json:"weight"`
	Volume         float64 `xml:"volume" json:"volume"`
}
type TracesDto struct {
	Trace *[]TraceDto `xml:"trace" json:"trace"`
}
type TraceDto struct {
	OpTime     string `xml:"opTime" json:"op_time"`
	OpStatus   string `xml:"opStatus" json:"op_status"`
	Operator   string `xml:"operator" json:"operator"`
	OpDesc     string `xml:"opDesc" json:"op_desc"`
	OpLocation string `xml:"opLocation" json:"op_location"`
}
