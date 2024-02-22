/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package squareup

// Currency : Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://wikipedia.org/wiki/ISO_4217).
type Currency string

// List of Currency
const (
	UNKNOWN_CURRENCY_Currency Currency = "UNKNOWN_CURRENCY"
	AED_Currency              Currency = "AED"
	AFN_Currency              Currency = "AFN"
	ALL_Currency              Currency = "ALL"
	AMD_Currency              Currency = "AMD"
	ANG_Currency              Currency = "ANG"
	AOA_Currency              Currency = "AOA"
	ARS_Currency              Currency = "ARS"
	AUD_Currency              Currency = "AUD"
	AWG_Currency              Currency = "AWG"
	AZN_Currency              Currency = "AZN"
	BAM_Currency              Currency = "BAM"
	BBD_Currency              Currency = "BBD"
	BDT_Currency              Currency = "BDT"
	BGN_Currency              Currency = "BGN"
	BHD_Currency              Currency = "BHD"
	BIF_Currency              Currency = "BIF"
	BMD_Currency              Currency = "BMD"
	BND_Currency              Currency = "BND"
	BOB_Currency              Currency = "BOB"
	BOV_Currency              Currency = "BOV"
	BRL_Currency              Currency = "BRL"
	BSD_Currency              Currency = "BSD"
	BTN_Currency              Currency = "BTN"
	BWP_Currency              Currency = "BWP"
	BYR_Currency              Currency = "BYR"
	BZD_Currency              Currency = "BZD"
	CAD_Currency              Currency = "CAD"
	CDF_Currency              Currency = "CDF"
	CHE_Currency              Currency = "CHE"
	CHF_Currency              Currency = "CHF"
	CHW_Currency              Currency = "CHW"
	CLF_Currency              Currency = "CLF"
	CLP_Currency              Currency = "CLP"
	CNY_Currency              Currency = "CNY"
	COP_Currency              Currency = "COP"
	COU_Currency              Currency = "COU"
	CRC_Currency              Currency = "CRC"
	CUC_Currency              Currency = "CUC"
	CUP_Currency              Currency = "CUP"
	CVE_Currency              Currency = "CVE"
	CZK_Currency              Currency = "CZK"
	DJF_Currency              Currency = "DJF"
	DKK_Currency              Currency = "DKK"
	DOP_Currency              Currency = "DOP"
	DZD_Currency              Currency = "DZD"
	EGP_Currency              Currency = "EGP"
	ERN_Currency              Currency = "ERN"
	ETB_Currency              Currency = "ETB"
	EUR_Currency              Currency = "EUR"
	FJD_Currency              Currency = "FJD"
	FKP_Currency              Currency = "FKP"
	GBP_Currency              Currency = "GBP"
	GEL_Currency              Currency = "GEL"
	GHS_Currency              Currency = "GHS"
	GIP_Currency              Currency = "GIP"
	GMD_Currency              Currency = "GMD"
	GNF_Currency              Currency = "GNF"
	GTQ_Currency              Currency = "GTQ"
	GYD_Currency              Currency = "GYD"
	HKD_Currency              Currency = "HKD"
	HNL_Currency              Currency = "HNL"
	HRK_Currency              Currency = "HRK"
	HTG_Currency              Currency = "HTG"
	HUF_Currency              Currency = "HUF"
	IDR_Currency              Currency = "IDR"
	ILS_Currency              Currency = "ILS"
	INR_Currency              Currency = "INR"
	IQD_Currency              Currency = "IQD"
	IRR_Currency              Currency = "IRR"
	ISK_Currency              Currency = "ISK"
	JMD_Currency              Currency = "JMD"
	JOD_Currency              Currency = "JOD"
	JPY_Currency              Currency = "JPY"
	KES_Currency              Currency = "KES"
	KGS_Currency              Currency = "KGS"
	KHR_Currency              Currency = "KHR"
	KMF_Currency              Currency = "KMF"
	KPW_Currency              Currency = "KPW"
	KRW_Currency              Currency = "KRW"
	KWD_Currency              Currency = "KWD"
	KYD_Currency              Currency = "KYD"
	KZT_Currency              Currency = "KZT"
	LAK_Currency              Currency = "LAK"
	LBP_Currency              Currency = "LBP"
	LKR_Currency              Currency = "LKR"
	LRD_Currency              Currency = "LRD"
	LSL_Currency              Currency = "LSL"
	LTL_Currency              Currency = "LTL"
	LVL_Currency              Currency = "LVL"
	LYD_Currency              Currency = "LYD"
	MAD_Currency              Currency = "MAD"
	MDL_Currency              Currency = "MDL"
	MGA_Currency              Currency = "MGA"
	MKD_Currency              Currency = "MKD"
	MMK_Currency              Currency = "MMK"
	MNT_Currency              Currency = "MNT"
	MOP_Currency              Currency = "MOP"
	MRO_Currency              Currency = "MRO"
	MUR_Currency              Currency = "MUR"
	MVR_Currency              Currency = "MVR"
	MWK_Currency              Currency = "MWK"
	MXN_Currency              Currency = "MXN"
	MXV_Currency              Currency = "MXV"
	MYR_Currency              Currency = "MYR"
	MZN_Currency              Currency = "MZN"
	NAD_Currency              Currency = "NAD"
	NGN_Currency              Currency = "NGN"
	NIO_Currency              Currency = "NIO"
	NOK_Currency              Currency = "NOK"
	NPR_Currency              Currency = "NPR"
	NZD_Currency              Currency = "NZD"
	OMR_Currency              Currency = "OMR"
	PAB_Currency              Currency = "PAB"
	PEN_Currency              Currency = "PEN"
	PGK_Currency              Currency = "PGK"
	PHP_Currency              Currency = "PHP"
	PKR_Currency              Currency = "PKR"
	PLN_Currency              Currency = "PLN"
	PYG_Currency              Currency = "PYG"
	QAR_Currency              Currency = "QAR"
	RON_Currency              Currency = "RON"
	RSD_Currency              Currency = "RSD"
	RUB_Currency              Currency = "RUB"
	RWF_Currency              Currency = "RWF"
	SAR_Currency              Currency = "SAR"
	SBD_Currency              Currency = "SBD"
	SCR_Currency              Currency = "SCR"
	SDG_Currency              Currency = "SDG"
	SEK_Currency              Currency = "SEK"
	SGD_Currency              Currency = "SGD"
	SHP_Currency              Currency = "SHP"
	SLL_Currency              Currency = "SLL"
	SOS_Currency              Currency = "SOS"
	SRD_Currency              Currency = "SRD"
	SSP_Currency              Currency = "SSP"
	STD_Currency              Currency = "STD"
	SVC_Currency              Currency = "SVC"
	SYP_Currency              Currency = "SYP"
	SZL_Currency              Currency = "SZL"
	THB_Currency              Currency = "THB"
	TJS_Currency              Currency = "TJS"
	TMT_Currency              Currency = "TMT"
	TND_Currency              Currency = "TND"
	TOP_Currency              Currency = "TOP"
	TRY_Currency              Currency = "TRY"
	TTD_Currency              Currency = "TTD"
	TWD_Currency              Currency = "TWD"
	TZS_Currency              Currency = "TZS"
	UAH_Currency              Currency = "UAH"
	UGX_Currency              Currency = "UGX"
	USD_Currency              Currency = "USD"
	USN_Currency              Currency = "USN"
	USS_Currency              Currency = "USS"
	UYI_Currency              Currency = "UYI"
	UYU_Currency              Currency = "UYU"
	UZS_Currency              Currency = "UZS"
	VEF_Currency              Currency = "VEF"
	VND_Currency              Currency = "VND"
	VUV_Currency              Currency = "VUV"
	WST_Currency              Currency = "WST"
	XAF_Currency              Currency = "XAF"
	XAG_Currency              Currency = "XAG"
	XAU_Currency              Currency = "XAU"
	XBA_Currency              Currency = "XBA"
	XBB_Currency              Currency = "XBB"
	XBC_Currency              Currency = "XBC"
	XBD_Currency              Currency = "XBD"
	XCD_Currency              Currency = "XCD"
	XDR_Currency              Currency = "XDR"
	XOF_Currency              Currency = "XOF"
	XPD_Currency              Currency = "XPD"
	XPF_Currency              Currency = "XPF"
	XPT_Currency              Currency = "XPT"
	XTS_Currency              Currency = "XTS"
	XXX_Currency              Currency = "XXX"
	YER_Currency              Currency = "YER"
	ZAR_Currency              Currency = "ZAR"
	ZMK_Currency              Currency = "ZMK"
	ZMW_Currency              Currency = "ZMW"
	BTC_Currency              Currency = "BTC"
	XUS_Currency              Currency = "XUS"
)
