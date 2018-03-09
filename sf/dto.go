package sf

import "encoding/xml"

//request

type ReqCustomerDto struct {
	Url       string `json:"url,omitempty"`
	Checkword string `json:"checkword,omitempty"`
}

type ReqConfirmDto struct {
	XMLName xml.Name `xml:"Request"`
	*ReqBaseDto
	Body *ReqBodyConfirmDto `xml:"Body,omitempty"`
}
type ReqRouteDto struct {
	XMLName xml.Name `xml:"Request"`
	*ReqBaseDto
	Body *ReqBodyRouteDto `xml:"Body,omitempty"`
}

type ReqCreateDto struct {
	XMLName xml.Name `xml:"Request"`
	*ReqBaseDto
	Body *ReqBodyCreateDto `xml:"Body,omitempty"`
}

//response

type RespConfirmDto struct {
	XMLName xml.Name            `xml:"Response,omitempty"`
	Service string              `xml:"service,attr,omitempty"`
	Head    string              `xml:"Head,omitempty"`
	Body    *RespBodyComfirmDto `xml:"Body,omitempty"`
	Error   RespErrorDto        `xml:"ERROR,omitempty"`
}

type RespRouteDto struct {
	XMLName xml.Name          `xml:"Response,omitempty"`
	Service string            `xml:"service,attr,omitempty"`
	Head    string            `xml:"Head,omitempty"`
	Body    *RespBodyRouteDto `xml:"Body,omitempty"`
	Error   RespErrorDto      `xml:"ERROR,omitempty"`
}

type RespCreateDto struct {
	XMLName xml.Name           `xml:"Response,omitempty"`
	Service string             `xml:"service,attr,omitempty"`
	Head    string             `xml:"Head,omitempty"`
	Body    *RespBodyCreateDto `xml:"Body,omitempty"`
	Error   RespErrorDto       `xml:"ERROR,omitempty"`
}

//util info

type ReqBodyConfirmDto struct {
	OrderConfirm *ReqOrderConfirmDto `xml:"OrderConfirm,omitempty"`
}
type ReqOrderConfirmDto struct {
	OrderId       string `xml:"orderid,attr,omitempty"`
	MailNo        string `xml:"mailno,attr,omitempty"`
	DealType      int64  `xml:"dealtype,attr,omitempty"`
	CustomsBatchs string `xml:"customs_batchs,attr,omitempty"`
	AgentNo       string `xml:"agent_no,attr,omitempty"`

	ConsignEmpCode     string                    `xml:"consign_emp_code,attr,omitempty"`
	SourceZoneCode     string                    `xml:"source_zone_code,attr,omitempty"`
	OrderConfirmOption *ReqOrderConfirmOptionDto `xml:OrderConfirmOption,omitempty`
	Extra              *ExtraDto                 `xml:"Extra,omitempty"`
}
type ReqOrderConfirmOptionDto struct {
	Weight         float64 `xml:"weight,attr,omitempty"`
	Volumn         string  `xml:"volumn,attr,omitempty"`
	ReturnTracting string  `xml:"return_tracting,attr,omitempty"`
	ExpressType    string  `xml:"express_type,attr,omitempty"`
	ChildrenNos    string  `xml:"children_nos,attr,omitempty"`

	WaybillSize     int64 `xml:"waybill_size,attr,omitempty"`
	IsGenEletricPic int64 `xml:"is_gen_eletric_pic,attr,omitempty"`
}

type ReqBodyRouteDto struct {
	RouteRequest *ReqRouteRequestDto `xml:"RouteRequest,omitempty"`
}

type ReqRouteRequestDto struct {
	TrackingType   int64     `xml:"tracking_type,attr,omitempty"`
	TrackingNumber string    `xml:"tracking_number,attr,omitempty"`
	MethodType     int64     `xml:"method_type,attr,omitempty"`
	Extra          *ExtraDto `xml:"Extra,omitempty"`
}

type ReqBodyCreateDto struct {
	Order *ReqOrderCreateDto `xml:"Order,omitempty"`
}

type ReqBaseDto struct {
	Service string `xml:"service,attr,omitempty"`
	Lang    string `xml:"lang,attr,omitempty"`
	Head    string `xml:"Head,omitempty"`
}

type ReqOrderCreateDto struct {
	OrderId     string `xml:"orderid,attr,omitempty"`
	MailNo      string `xml:"mailno,attr,omitempty"`
	IsGenBillNo int64  `xml:"is_gen_bill_no,attr,omitempty"`
	JCompany    string `xml:"j_company,attr,omitempty"`
	JContact    string `xml:"j_contact,attr,omitempty"`

	JTel         string `xml:"j_tel,attr,omitempty"`
	JMobile      string `xml:"j_mobile,attr,omitempty"`
	JShipperCode string `xml:"j_shippercode,attr,omitempty"`
	JCountry     string `xml:"j_country,attr,omitempty"`
	JProvince    string `xml:"j_province,attr,omitempty"`

	JCity    string `xml:"j_city,attr,omitempty"`
	JCounty  string `xml:"j_county,attr,omitempty"`
	JAddress string `xml:"j_address,attr,omitempty"`
	JPosCode string `xml:"j_pos_code,attr,omitempty"`
	DCompany string `xml:"d_company,attr,omitempty"`

	DContact      string `xml:"d_contact,attr,omitempty"`
	DTel          string `xml:"d_tel,attr,omitempty"`
	DMobile       string `xml:"d_mobile,attr,omitempty"`
	DDeliveryCode string `xml:"d_deliverycode,attr,omitempty"`
	DCountry      string `xml:"d_country,attr,omitempty"`

	DProvince string `xml:"d_province,attr,omitempty"`
	DCity     string `xml:"d_city,attr,omitempty"`
	DCounty   string `xml:"d_county,attr,omitempty"`
	DAddress  string `xml:"d_address,attr,omitempty"`
	DPostCode string `xml:"d_post_code,attr,omitempty"`

	CustId         string  `xml:"custid,attr,omitempty"`
	PayMethod      int64   `xml:"pay_method,attr,omitempty"`
	ExpressType    string  `xml:"express_type,attr,omitempty"`
	ParcelQuantity int64   `xml:"parcel_quantity,attr,omitempty"`
	CargoLength    float64 `xml:"cargo_length,attr,omitempty"`

	CargoWidth       float64 `xml:"cargo_width,attr,omitempty"`
	CargoHeight      float64 `xml:"cargo_height,attr,omitempty"`
	Volumn           float64 `xml:"volumn,attr,omitempty"`
	CargoTotalWeight float64 `xml:"cargo_total_weight,attr,omitempty"`
	DeclareValue     float64 `xml:"declare_value,attr,omitempty"`

	DeclareValueCurrency string `xml:"declare_value_currency,attr,omitempty"`
	CustomersBatchs      string `xml:"customer_batchs,attr,omitempty"`
	SendStartTime        string `xml:"sendstarttime,attr,omitempty"` //yyyy-MM-dd HH:mm:ss
	IsDocall             int64  `xml:"is_docall,attr,omitempty"`
	NeedReturnTrackingNo string `xml:"need_return_tracking_no,attr,omitempty"`

	ReturnTracking string `xml:"return_tracking,attr,omitempty"`
	DTaxNo         string `xml:"d_tax_no,attr,omitempty"`
	TaxPayType     string `xml:"tax_pay_type,attr,omitempty"`
	TaxSetAccounts string `xml:"tax_set_accounts,attr,omitempty"`
	OriginalNumber string `xml:"original_number,attr,omitempty"`

	GoodsCode          string `xml:"goods_code,attr,omitempty"`
	InProcessWaybillNo string `xml:"in_process_waybill_no,attr,omitempty"`
	Brand              string `xml:"brand,attr,omitempty"`
	Specifications     string `xml:"specifications,attr,omitempty"`
	TempRange          int64  `xml:"temp_range,attr,omitempty"`

	OrderName     string `xml:"order_name,attr,omitempty"`
	OrderCertType string `xml:"order_cert_type,attr,omitempty"`
	OrderCertNo   string `xml:"order_cert_no,attr,omitempty"`
	OrderSource   string `xml:"order_source,attr,omitempty"`
	Template      string `xml:"template,attr,omitempty"`

	Remark           string `xml:"remark,attr,omitempty"`
	OneselfPickupFlg int64  `xml:"oneself_pickup_flg,attr,omitempty"`
	DispatchSys      string `xml:"dispatch_sys,attr,omitempty"`
	//ParcelQuantity   string `xml:"parcel_quantity,attr,omitempty"`
	IsGenEletricPic int64 `xml:"is_gen_eletric_pic,attr,omitempty"`

	WaybillSize    float64          `xml:"waybill_size,attr,omitempty"`
	FilterField    string           `xml:"filter_field,attr,omitempty"`
	TotalNetWeight float64          `xml:"total_net_weight,attr,omitempty"`
	Cargo          *CargoDto        `xml:"Cargo,omitempty"`
	AddedService   *AddedServiceDto `xml "AddedService,omitempty"`
	Extra          *ExtraDto        `xml:"Extra,omitempty"`
}

type CargoDto struct {
	Name   string  `xml:"name,attr,omitempty"`
	Count  int64   `xml:"count,attr,omitempty"`
	Unit   string  `xml:"unit,attr,omitempty"`
	Weight float64 `xml:"weight,attr,omitempty"`
	Amount float64 `xml:"amount,attr,omitempty"`

	Currency        string `xml:"currency,attr,omitempty"`
	SourceArea      string `xml:"source_area,attr,omitempty"`
	ProductRecordNo string `xml:"product_record_no,attr,omitempty"`
	GoodPrepardNo   string `xml:"good_prepard_no,attr,omitempty"`
	TaxNo           string `xml:"tax_no,attr,omitempty"`
}

type AddedServiceDto struct {
	Name   string `xml:"name,attr,omitempty"`
	Value  string `xml:"value,attr,omitempty"`
	Value1 string `xml:"value1,attr,omitempty"`
	Value2 string `xml:"value2,attr,omitempty"`
	Value3 string `xml:"value3,attr,omitempty"`
	Value4 string `xml:"value4,attr,omitempty"`
}

type ExtraDto struct {
	E1 string `xml:"e1,attr,omitempty"`
	E2 string `xml:"e2,attr,omitempty"`
	E3 string `xml:"e3,attr,omitempty"`
	E4 string `xml:"e4,attr,omitempty"`
	E5 string `xml:"e5,attr,omitempty"`

	E6  string `xml:"e6,attr,omitempty"`
	E7  string `xml:"e7,attr,omitempty"`
	E8  string `xml:"e8,attr,omitempty"`
	E9  string `xml:"e9,attr,omitempty"`
	E10 string `xml:"e10,attr,omitempty"`

	E11 string `xml:"e11,attr,omitempty"`
	E12 string `xml:"e12,attr,omitempty"`
	E13 string `xml:"e13,attr,omitempty"`
	E14 string `xml:"e14,attr,omitempty"`
	E15 string `xml:"e15,attr,omitempty"`

	E16 string `xml:"e16,attr,omitempty"`
	E17 string `xml:"e17,attr,omitempty"`
	E18 string `xml:"e18,attr,omitempty"`
	E19 string `xml:"e19,attr,omitempty"`
	E20 string `xml:"e20,attr,omitempty"`
}

type RespErrorDto struct {
	Code  string `xml:"code,attr,omitempty"`
	Error string `xml:",chardata"`
}
type RespBodyComfirmDto struct {
	OrderConfirmResponse *OrderConfirmResponseDto `xml:"OrderConfirmResponse,omitempty"`
}
type OrderConfirmResponseDto struct {
	OrderId   string `xml:"orderid,attr,omitempty"`
	MailNo    string `xml:"mail_no,attr,omitempty"`
	ResStatus int64  `xml:"res_status,attr,omitempty"`
}

type RespBodyCreateDto struct {
	OrderResponse *OrderResponseDto `xml:"OrderResponse,omitempty"`
}
type OrderResponseDto struct {
	OrderId          string `xml:"orderid,attr,omitempty"`
	MailNo           string `xml:"mail_no,attr,omitempty"`
	ReturnTrackingNo string `xml:"return_tracking_no,attr,omitempty"`
	OriginCode       string `xml:"origincode,attr,omitempty"`
	DestCode         string `xml:"destcode,attr,omitempty"`

	FilterResult int64  `xml:"filter_result,attr,omitempty"`
	Remark       string `xml:"remark,attr,omitempty"`
	AgentMailNo  string `xml:"agent_mail_no,attr,omitempty"`
}

type RespBodyRouteDto struct {
	RouteResponse *RouteResponseDto `xml:"RouteResponse,omitempty"`
}

type RouteResponseDto struct {
	OrderId string      `xml:"orderid,attr,omitempty"`
	MailNo  string      `xml:"mail_no,attr,omitempty"`
	Route   *[]RouteDto `xml:"Route,omitempty"`
}
type RouteDto struct {
	AcceptTime    string `xml:"accept_time,attr,omitempty"`
	AcceptAddress string `xml:"accept_address,attr,omitempty"`
	Remark        string `xml:"remark,attr,omitempty"`
	OpCode        string `xml:"opcode,attr,omitempty"`
}
