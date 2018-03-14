package sf

import "encoding/xml"

//request

type ReqPushDto struct {
	XMLName xml.Name `xml:"Request" json:"Request,omitempty"`
	*ReqBaseDto
	Body *ReqBodyPushDto `xml:"Body,omitempty" json:"Body,omitempty"`
}

type ReqConfirmDto struct {
	XMLName xml.Name `xml:"Request" json:"Request,omitempty"`
	*ReqBaseDto
	Body *ReqBodyConfirmDto `xml:"Body,omitempty" json:"Body,omitempty"`
}
type ReqRouteDto struct {
	XMLName xml.Name `xml:"Request" json:"Request,omitempty"`
	*ReqBaseDto
	Body *ReqBodyRouteDto `xml:"Body,omitempty" json:"Body,omitempty"`
}

type ReqCreateDto struct {
	XMLName xml.Name `xml:"Request" json:"Request,omitempty"`
	*ReqBaseDto
	Body *ReqBodyCreateDto `xml:"Body,omitempty" json:"Body,omitempty"`
}

type ReqCustomerDto struct {
	Url       string `json:"url,omitempty" json:"url,omitempty"`
	Checkword string `json:"checkword,omitempty" json:"checkword,omitempty"`
}

//response

type RespConfirmDto struct {
	XMLName xml.Name            `xml:"Response,omitempty" json:"Response,omitempty"`
	Service string              `xml:"service,attr,omitempty" json:"service,omitempty"`
	Head    string              `xml:"Head,omitempty" json:"Head,omitempty"`
	Body    *RespBodyComfirmDto `xml:"Body,omitempty" json:"Body,omitempty"`
	Error   RespErrorDto        `xml:"ERROR,omitempty" json:"ERROR,omitempty"`
}

type RespRouteDto struct {
	XMLName xml.Name          `xml:"Response,omitempty" json:"Response,omitempty"`
	Service string            `xml:"service,attr,omitempty" json:"service,omitempty"`
	Head    string            `xml:"Head,omitempty" json:"Head,omitempty"`
	Body    *RespBodyRouteDto `xml:"Body,omitempty" json:"Body,omitempty"`
	Error   RespErrorDto      `xml:"ERROR,omitempty" json:"ERROR,omitempty"`
}

type RespCreateDto struct {
	XMLName xml.Name           `xml:"Response,omitempty" json:"Response,omitempty"`
	Service string             `xml:"service,attr,omitempty" json:"service,omitempty"`
	Head    string             `xml:"Head,omitempty" json:"Head,omitempty"`
	Body    *RespBodyCreateDto `xml:"Body,omitempty" json:"Body,omitempty"`
	Error   RespErrorDto       `xml:"ERROR,omitempty" json:"ERROR,omitempty"`
}

//util info

type ReqBodyPushDto struct {
	WaybillRoute *ReqWaybillRouteDto `xml:"WaybillRoute,omitempty" json:"WaybillRoute,omitempty"`
}

type ReqWaybillRouteDto struct {
	Id           int64  `xml:"id,attr,omitempty" json:"id,omitempty"`
	OrderId      string `xml:"orderid,attr,omitempty" json:"orderid,omitempty"`
	MailNo       string `xml:"mailno,attr,omitempty" json:"mailno,omitempty"`
	AcceptTime   string `xml:"acceptTime,attr,omitempty" json:"acceptTime,omitempty"`
	AcceptAddres string `xml:"acceptAddress,attr,omitempty" json:"acceptAddress,omitempty"`

	Remark string `xml:"remark,attr,omitempty" json:"remark,omitempty"`
	OpCode string `xml:"opCode,attr,omitempty" json:"opCode,omitempty"`
}

type ReqBodyConfirmDto struct {
	OrderConfirm *ReqOrderConfirmDto `xml:"OrderConfirm,omitempty" json:"OrderConfirm,omitempty"`
}
type ReqOrderConfirmDto struct {
	OrderId       string `xml:"orderid,attr,omitempty" json:"orderid,omitempty"`
	MailNo        string `xml:"mailno,attr,omitempty" json:"mailno,omitempty"`
	DealType      int64  `xml:"dealtype,attr,omitempty" json:"dealtype,omitempty"`
	CustomsBatchs string `xml:"customs_batchs,attr,omitempty" json:"customs_batchs,omitempty"`
	AgentNo       string `xml:"agent_no,attr,omitempty" json:"agent_no,omitempty"`

	ConsignEmpCode     string                    `xml:"consign_emp_code,attr,omitempty" json:"consign_emp_code,omitempty"`
	SourceZoneCode     string                    `xml:"source_zone_code,attr,omitempty" json:"source_zone_code,omitempty"`
	OrderConfirmOption *ReqOrderConfirmOptionDto `xml:"OrderConfirmOption,omitempty"  json:"OrderConfirmOption,omitempty"`
	Extra              *ExtraDto                 `xml:"Extra,omitempty"  json:"Extra,omitempty"`
}
type ReqOrderConfirmOptionDto struct {
	Weight         float64 `xml:"weight,attr,omitempty"  json:"weight,omitempty"`
	Volumn         string  `xml:"volumn,attr,omitempty"  json:"volumn,omitempty"`
	ReturnTracting string  `xml:"return_tracting,attr,omitempty"  json:"return_tracting,omitempty"`
	ExpressType    string  `xml:"express_type,attr,omitempty"  json:"express_type,omitempty"`
	ChildrenNos    string  `xml:"children_nos,attr,omitempty"  json:"children_nos,omitempty"`

	WaybillSize     int64 `xml:"waybill_size,attr,omitempty"  json:"waybill_size,omitempty"`
	IsGenEletricPic int64 `xml:"is_gen_eletric_pic,attr,omitempty"  json:"is_gen_eletric_pic,omitempty"`
}

type ReqBodyRouteDto struct {
	RouteRequest *ReqRouteRequestDto `xml:"RouteRequest,omitempty"  json:"RouteRequest,omitempty"`
}

type ReqRouteRequestDto struct {
	TrackingType   int64     `xml:"tracking_type,attr,omitempty"  json:"tracking_type,omitempty"`
	TrackingNumber string    `xml:"tracking_number,attr,omitempty"  json:"tracking_number,omitempty"`
	MethodType     int64     `xml:"method_type,attr,omitempty"  json:"method_type,omitempty"`
	Extra          *ExtraDto `xml:"Extra,omitempty"  json:"Extra,omitempty"`
}

type ReqBodyCreateDto struct {
	Order *ReqOrderCreateDto `xml:"Order,omitempty"  json:"Order,omitempty"`
}

type ReqBaseDto struct {
	Service string `xml:"service,attr,omitempty"  json:"service,omitempty"`
	Lang    string `xml:"lang,attr,omitempty"  json:"lang,omitempty"`
	Head    string `xml:"Head,omitempty"  json:"Head,omitempty"`
}

type ReqOrderCreateDto struct {
	OrderId     string `xml:"orderid,attr,omitempty"  json:"orderid,omitempty"`
	MailNo      string `xml:"mailno,attr,omitempty"  json:"mailno,omitempty"`
	IsGenBillNo int64  `xml:"is_gen_bill_no,attr,omitempty"  json:"is_gen_bill_no,omitempty"`
	JCompany    string `xml:"j_company,attr,omitempty"  json:"j_company,omitempty"`
	JContact    string `xml:"j_contact,attr,omitempty"  json:"j_contact,omitempty"`

	JTel         string `xml:"j_tel,attr,omitempty"  json:"j_tel,omitempty"`
	JMobile      string `xml:"j_mobile,attr,omitempty"  json:"j_mobile,omitempty"`
	JShipperCode string `xml:"j_shippercode,attr,omitempty"  json:"j_shippercode,omitempty"`
	JCountry     string `xml:"j_country,attr,omitempty"  json:"j_country,omitempty"`
	JProvince    string `xml:"j_province,attr,omitempty"  json:"j_province,omitempty"`

	JCity    string `xml:"j_city,attr,omitempty"  json:"j_city,omitempty"`
	JCounty  string `xml:"j_county,attr,omitempty"  json:"j_county,omitempty"`
	JAddress string `xml:"j_address,attr,omitempty"  json:"j_address,omitempty"`
	JPosCode string `xml:"j_pos_code,attr,omitempty"  json:"j_pos_code,omitempty"`
	DCompany string `xml:"d_company,attr,omitempty"  json:"d_company,omitempty"`

	DContact      string `xml:"d_contact,attr,omitempty"  json:"d_contact,omitempty"`
	DTel          string `xml:"d_tel,attr,omitempty"  json:"d_tel,omitempty"`
	DMobile       string `xml:"d_mobile,attr,omitempty"  json:"d_mobile,omitempty"`
	DDeliveryCode string `xml:"d_deliverycode,attr,omitempty"  json:"d_deliverycode,omitempty"`
	DCountry      string `xml:"d_country,attr,omitempty"  json:"d_country,omitempty"`

	DProvince string `xml:"d_province,attr,omitempty"  json:"d_province,omitempty"`
	DCity     string `xml:"d_city,attr,omitempty"  json:"d_city,omitempty"`
	DCounty   string `xml:"d_county,attr,omitempty"  json:"d_county,omitempty"`
	DAddress  string `xml:"d_address,attr,omitempty"  json:"d_address,omitempty"`
	DPostCode string `xml:"d_post_code,attr,omitempty"  json:"d_post_code,omitempty"`

	CustId         string  `xml:"custid,attr,omitempty"  json:"custid,omitempty"`
	PayMethod      int64   `xml:"pay_method,attr,omitempty"  json:"pay_method,omitempty"`
	ExpressType    string  `xml:"express_type,attr,omitempty"  json:"express_type,omitempty"`
	ParcelQuantity int64   `xml:"parcel_quantity,attr,omitempty"  json:"parcel_quantity,omitempty"`
	CargoLength    float64 `xml:"cargo_length,attr,omitempty"  json:"cargo_length,omitempty"`

	CargoWidth       float64 `xml:"cargo_width,attr,omitempty"  json:"cargo_width,omitempty"`
	CargoHeight      float64 `xml:"cargo_height,attr,omitempty"  json:"cargo_height,omitempty"`
	Volumn           float64 `xml:"volumn,attr,omitempty"  json:"volumn,omitempty"`
	CargoTotalWeight float64 `xml:"cargo_total_weight,attr,omitempty"  json:"cargo_total_weight,omitempty"`
	DeclareValue     float64 `xml:"declare_value,attr,omitempty"  json:"declare_value,omitempty"`

	DeclareValueCurrency string `xml:"declare_value_currency,attr,omitempty"  json:"declare_value_currency,omitempty"`
	CustomersBatchs      string `xml:"customer_batchs,attr,omitempty"  json:"customer_batchs,omitempty"`
	SendStartTime        string `xml:"sendstarttime,attr,omitempty"  json:"sendstarttime,omitempty"` //yyyy-MM-dd HH:mm:ss
	IsDocall             int64  `xml:"is_docall,attr,omitempty"  json:"is_docall,omitempty"`
	NeedReturnTrackingNo string `xml:"need_return_tracking_no,attr,omitempty"  json:"need_return_tracking_no,omitempty"`

	ReturnTracking string `xml:"return_tracking,attr,omitempty"  json:"return_tracking,omitempty"`
	DTaxNo         string `xml:"d_tax_no,attr,omitempty"  json:"d_tax_no,omitempty"`
	TaxPayType     string `xml:"tax_pay_type,attr,omitempty"  json:"tax_pay_type,omitempty"`
	TaxSetAccounts string `xml:"tax_set_accounts,attr,omitempty"  json:"tax_set_accounts,omitempty"`
	OriginalNumber string `xml:"original_number,attr,omitempty"  json:"original_number,omitempty"`

	GoodsCode          string `xml:"goods_code,attr,omitempty"  json:"goods_code,omitempty"`
	InProcessWaybillNo string `xml:"in_process_waybill_no,attr,omitempty"  json:"in_process_waybill_no,omitempty"`
	Brand              string `xml:"brand,attr,omitempty"  json:"brand,omitempty"`
	Specifications     string `xml:"specifications,attr,omitempty"  json:"specifications,omitempty"`
	TempRange          int64  `xml:"temp_range,attr,omitempty"  json:"temp_range,omitempty"`

	OrderName     string `xml:"order_name,attr,omitempty"  json:"order_name,omitempty"`
	OrderCertType string `xml:"order_cert_type,attr,omitempty"  json:"order_cert_type,omitempty"`
	OrderCertNo   string `xml:"order_cert_no,attr,omitempty"  json:"order_cert_no,omitempty"`
	OrderSource   string `xml:"order_source,attr,omitempty"  json:"order_source,omitempty"`
	Template      string `xml:"template,attr,omitempty"  json:"template,omitempty"`

	Remark           string `xml:"remark,attr,omitempty"  json:"remark,omitempty"`
	OneselfPickupFlg int64  `xml:"oneself_pickup_flg,attr,omitempty"  json:"oneself_pickup_flg,omitempty"`
	DispatchSys      string `xml:"dispatch_sys,attr,omitempty"  json:"dispatch_sys,omitempty"`
	//ParcelQuantity   string `xml:"parcel_quantity,attr,omitempty"`
	IsGenEletricPic int64 `xml:"is_gen_eletric_pic,attr,omitempty"  json:"is_gen_eletric_pic,omitempty"`

	WaybillSize    float64          `xml:"waybill_size,attr,omitempty"  json:"waybill_size,omitempty"`
	FilterField    string           `xml:"filter_field,attr,omitempty"  json:"filter_field,omitempty"`
	TotalNetWeight float64          `xml:"total_net_weight,attr,omitempty"  json:"total_net_weight,omitempty"`
	Cargo          *CargoDto        `xml:"Cargo,omitempty"  json:"Cargo,omitempty"`
	AddedService   *AddedServiceDto `xml "AddedService,omitempty"  json:"AddedService,omitempty"`
	Extra          *ExtraDto        `xml:"Extra,omitempty"  json:"Extra,omitempty"`
}

type CargoDto struct {
	Name   string  `xml:"name,attr,omitempty"  json:"name,omitempty"`
	Count  int64   `xml:"count,attr,omitempty"  json:"count,omitempty"`
	Unit   string  `xml:"unit,attr,omitempty"  json:"unit,omitempty"`
	Weight float64 `xml:"weight,attr,omitempty"  json:"weight,omitempty"`
	Amount float64 `xml:"amount,attr,omitempty"  json:"amount,omitempty"`

	Currency        string `xml:"currency,attr,omitempty"  json:"currency,omitempty"`
	SourceArea      string `xml:"source_area,attr,omitempty"  json:"source_area,omitempty"`
	ProductRecordNo string `xml:"product_record_no,attr,omitempty"  json:"product_record_no,omitempty"`
	GoodPrepardNo   string `xml:"good_prepard_no,attr,omitempty"  json:"good_prepard_no,omitempty"`
	TaxNo           string `xml:"tax_no,attr,omitempty"  json:"tax_no,omitempty"`
}

type AddedServiceDto struct {
	Name   string `xml:"name,attr,omitempty"  json:"name,omitempty"`
	Value  string `xml:"value,attr,omitempty"  json:"value,omitempty"`
	Value1 string `xml:"value1,attr,omitempty"  json:"value1,omitempty"`
	Value2 string `xml:"value2,attr,omitempty"  json:"value2,omitempty"`
	Value3 string `xml:"value3,attr,omitempty"  json:"value3,omitempty"`
	Value4 string `xml:"value4,attr,omitempty"  json:"value4,omitempty"`
}

type ExtraDto struct {
	E1 string `xml:"e1,attr,omitempty" json:"e1,omitempty"`
	E2 string `xml:"e2,attr,omitempty" json:"e2,omitempty"`
	E3 string `xml:"e3,attr,omitempty" json:"e3,omitempty"`
	E4 string `xml:"e4,attr,omitempty" json:"e4,omitempty"`
	E5 string `xml:"e5,attr,omitempty" json:"e5,omitempty"`

	E6  string `xml:"e6,attr,omitempty" json:"e6,omitempty"`
	E7  string `xml:"e7,attr,omitempty" json:"e7,omitempty"`
	E8  string `xml:"e8,attr,omitempty" json:"e8,omitempty"`
	E9  string `xml:"e9,attr,omitempty" json:"e9,omitempty"`
	E10 string `xml:"e10,attr,omitempty" json:"e10,omitempty"`

	E11 string `xml:"e11,attr,omitempty" json:"e11,omitempty"`
	E12 string `xml:"e12,attr,omitempty" json:"e12,omitempty"`
	E13 string `xml:"e13,attr,omitempty" json:"e13,omitempty"`
	E14 string `xml:"e14,attr,omitempty" json:"e14,omitempty"`
	E15 string `xml:"e15,attr,omitempty" json:"e15,omitempty"`

	E16 string `xml:"e16,attr,omitempty" json:"e16,omitempty"`
	E17 string `xml:"e17,attr,omitempty" json:"e17,omitempty"`
	E18 string `xml:"e18,attr,omitempty" json:"e18,omitempty"`
	E19 string `xml:"e19,attr,omitempty" json:"e19,omitempty"`
	E20 string `xml:"e20,attr,omitempty" json:"e20,omitempty"`
}

type RespErrorDto struct {
	Code  string `xml:"code,attr,omitempty" json:"code,omitempty"`
	Error string `xml:",chardata" json:"Error"`
}
type RespBodyComfirmDto struct {
	OrderConfirmResponse *OrderConfirmResponseDto `xml:"OrderConfirmResponse,omitempty" json:"OrderConfirmResponse"`
}
type OrderConfirmResponseDto struct {
	OrderId   string `xml:"orderid,attr,omitempty" json:"orderid"`
	MailNo    string `xml:"mail_no,attr,omitempty" json:"mail_no"`
	ResStatus int64  `xml:"res_status,attr,omitempty" json:"res_status"`
}

type RespBodyCreateDto struct {
	OrderResponse *OrderResponseDto `xml:"OrderResponse,omitempty" json:"OrderResponse"`
}
type OrderResponseDto struct {
	OrderId          string `xml:"orderid,attr,omitempty" json:"orderid"`
	MailNo           string `xml:"mail_no,attr,omitempty" json:"mail_no"`
	ReturnTrackingNo string `xml:"return_tracking_no,attr,omitempty" json:"return_tracking_no"`
	OriginCode       string `xml:"origincode,attr,omitempty" json:"origincode"`
	DestCode         string `xml:"destcode,attr,omitempty" json:"destcode"`

	FilterResult int64  `xml:"filter_result,attr,omitempty" json:"filter_result"`
	Remark       string `xml:"remark,attr,omitempty" json:"remark"`
	AgentMailNo  string `xml:"agent_mail_no,attr,omitempty" json:"agent_mail_no"`
}

type RespBodyRouteDto struct {
	RouteResponse *RouteResponseDto `xml:"RouteResponse,omitempty" json:"RouteResponse"`
}

type RouteResponseDto struct {
	OrderId string      `xml:"orderid,attr,omitempty" json:"orderid"`
	MailNo  string      `xml:"mail_no,attr,omitempty" json:"mail_no"`
	Route   *[]RouteDto `xml:"Route,omitempty" json:"Route"`
}
type RouteDto struct {
	AcceptTime    string `xml:"accept_time,attr,omitempty" json:"accept_time"`
	AcceptAddress string `xml:"accept_address,attr,omitempty" json:"accept_address"`
	Remark        string `xml:"remark,attr,omitempty" json:"remark"`
	OpCode        string `xml:"opcode,attr,omitempty" json:"opcode"`
}
