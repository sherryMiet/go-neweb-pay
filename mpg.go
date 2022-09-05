package neweb_pay

import (
	"net/url"
	"strconv"
	"time"
)

type (
	RespondType string
	WEBATM      int
	CREDIT      int
)

const (
	Version                        = "2.0"
	TestMPGGatewayUrl              = "https://ccore.newebpay.com/MPG/mpg_gateway"
	MPGGatewayUrl                  = "https://core.newebpay.com/MPG/mpg_gateway"
	RespondType_JSON   RespondType = "JSON"
	RespondType_STRING RespondType = "STRING"
	TradeLimit                     = 900
	WEBATM_Y           WEBATM      = 1
	WEBATM_N           WEBATM      = 0
	CREDIT_Y           CREDIT      = 1
	CREDIT_N           CREDIT      = 0
)

type Client struct {
	//商店代號
	MerchantID string
	//交易資料 AES 加密
	HashKey string
	//交易資料 SHA256 加密
	HashIV string
}

type MPGGatewayRequestCall struct {
	HashKey           string
	HashIV            string
	MPGGatewayRequest *MPGGatewayRequest
}

type MPGGatewayRequest struct {
	//商店代號
	MerchantID string `json:"MerchantID"`
	//交易資料 AES 加密
	TradeInfo string `json:"TradeInfo"`
	//交易資料 SHA256 加密 將交易資料經過上述 AES 加密過的字串， 透過商店 Key 及 IV 進行 SHA256 加密
	TradeSha string `json:"TradeSha"`
	//串接程式版本
	Version string `json:"Version"`
	//加密模式
	EncryptType int `json:"EncryptType"`
}

type MPGGateWayTradeInfo struct {
	//商店代號
	MerchantID string `json:"MerchantID"`
	//回傳格式 JSON 或是 String
	RespondType RespondType `json:"RespondType"`
	//時間戳記
	TimeStamp string `json:"TimeStamp"`
	//串接程式版本
	Version string `json:"Version"`
	//語系
	LangType string `json:"LangType,omitempty"`
	//商店訂單編號 限英、數字、”_ ”格式  長度限制為 30 字元 編號不可重覆
	MerchantOrderNo string `json:"MerchantOrderNo"`
	//訂單金額
	Amt int `json:"Amt"`
	//商品資訊 限制長度為 50 字元 Utf-8 勿使用斷行符號、單引號等特殊符號 使用特殊符號，系統將自動過濾
	ItemDesc string `json:"ItemDesc"`
	//交易限制秒數
	TradeLimit int `json:"TradeLimit,omitempty"`
	//繳費有效期限
	ExpireDate string `json:"ExpireDate,omitempty"`
	//支付完成 返回商店網址
	ReturnURL string `json:"ReturnURL,omitempty"`
	//支付通知網址
	NotifyURL string `json:"NotifyURL,omitempty"`
	//商店取號網址
	CustomerURL string `json:"CustomerURL,omitempty"`
	//返回商店網址
	ClientBackURL string `json:"ClientBackURL,omitempty"`
	//付款人電子信箱
	Email string `json:"Email,omitempty"`
	//付款人電子信箱是否開放修改
	EmailModify int `json:"EmailModify,omitempty"`
	//藍新金流會員
	LoginType int `json:"LoginType,omitempty"`
	//商店備註
	OrderComment string `json:"OrderComment,omitempty"`
	//信用卡 一次付清啟用
	CREDIT int `json:"CREDIT,omitempty"`
	//Google Pay啟用
	ANDROIDPAY int `json:"ANDROIDPAY,omitempty"`
	//Samsung Pay 啟用
	SAMSUNGPAY int `json:"SAMSUNGPAY,omitempty"`
	//LINE Pay
	LINEPAY int `json:"LINEPAY,omitempty"`
	//LINE Pay產品圖檔連結 網址
	ImageUrl string `json:"ImageUrl,omitempty"`
	//信用卡 分期付款啟用 同時開啟多期別時，將此參數用”，”(半形) 分隔，例如:3,6,12，代表開啟 分 3、6、12 期的功能
	InstFlag string `json:"InstFlag,omitempty"`
	//信用卡紅利啟用
	CreditRed int `json:"CreditRed,omitempty"`
	//信用卡 銀聯卡啟用
	UNIONPAY int `json:"UNIONPAY,omitempty"`
	//WEBATM 啟用
	WEBATM WEBATM `json:"WEBATM,omitempty"`
	//ATM 轉帳啟用
	VACC int `json:"VACC,omitempty"`
	//金融機構
	BankType string `json:"BankType,omitempty"`
	//超商代碼繳費 啟用
	CVS int `json:"CVS,omitempty"`
	//超商條碼繳費 啟用
	BARCODE int `json:"BARCODE,omitempty"`
	//玉山 Wallet
	ESUNWALLET int `json:"ESUNWALLET,omitempty"`
	//台灣 Pay
	TAIWANPAY int `json:"TAIWANPAY,omitempty"`
	//物流啟用
	CVSCOM int `json:"CVSCOM,omitempty"`
	//簡單付電子錢包
	EZPAY int `json:"EZPAY,omitempty"`
	//簡單付微信支付
	EZPWECHAT int `json:"EZPWECHAT,omitempty"`
	//簡單付支付寶
	EZPALIPAY int `json:"EZPALIPAY,omitempty"`
	//物流型態
	LgsType string `json:"LgsType,omitempty"`
	//信用卡 國民旅遊卡交易 註記
	NTCB int `json:"NTCB,omitempty"`
	//旅遊地區代號
	NTCBLocate string `json:"NTCBLocate,omitempty"`
	//國民旅遊卡起始 日期
	NTCBStartDate string `json:"NTCBStartDate,omitempty"`
	//國民旅遊卡結束 日期
	NTCBEndDate string `json:"NTCBEndDate,omitempty"`
	//付款人綁定資料
	TokenTerm string `json:"TokenTerm"`
	//指定付款人信用卡快速結帳必填 欄位
	TokenTermDemand string `json:"TokenTermDemand,omitempty"`
}

type OptionValue struct {
	//回傳格式 JSON 或是 String
	RespondType RespondType `json:"RespondType"`
	//時間戳記
	TimeStamp string `json:"TimeStamp"`
	//串接程式版本
	Version string `json:"Version"`
	//語系
	LangType string `json:"LangType,omitempty"`
	//商店訂單編號 限英、數字、”_ ”格式  長度限制為 30 字元 編號不可重覆
	MerchantOrderNo string `json:"MerchantOrderNo"`
	//交易限制秒數
	TradeLimit int `json:"TradeLimit,omitempty"`
	//支付完成 返回商店網址
	ReturnURL string `json:"ReturnURL,omitempty"`
	//支付通知網址
	NotifyURL string `json:"NotifyURL,omitempty"`
	//商店取號網址
	CustomerURL string `json:"CustomerURL,omitempty"`
	//返回商店網址
	ClientBackURL string `json:"ClientBackURL,omitempty"`
	//付款人電子信箱
	Email string `json:"Email,omitempty"`
	//付款人電子信箱是否開放修改
	EmailModify int `json:"EmailModify,omitempty"`
	//藍新金流會員
	LoginType int `json:"LoginType,omitempty"`
	//商店備註
	OrderComment string `json:"OrderComment,omitempty"`
	//付款人綁定資料
	TokenTerm string `json:"TokenTerm"`
	//指定付款人信用卡快速結帳必填 欄位
	TokenTermDemand string `json:"TokenTermDemand,omitempty"`
	//金融機構
	BankType string `json:"BankType,omitempty"`
}

type AtmPayment struct {
	//繳費有效期限
	ExpireDate string `json:"ExpireDate,omitempty"`
	//WEBATM 啟用
	WEBATM int `json:"WEBATM,omitempty"`
}

func NewClient(MerchantID string, HashKey string, HashIV string) *Client {
	return &Client{
		MerchantID: MerchantID,
		HashKey:    HashKey,
		HashIV:     HashIV,
	}
}

func NewMPGGateWayTradeInfo() *MPGGateWayTradeInfo {
	return &MPGGateWayTradeInfo{}
}

func (m *MPGGateWayTradeInfo) WithOptional(OptionValue OptionValue) *MPGGateWayTradeInfo {
	if OptionValue.RespondType == "" {
		m.RespondType = RespondType_JSON
	} else {
		m.RespondType = OptionValue.RespondType
	}
	if OptionValue.TimeStamp == "" {
		m.TimeStamp = strconv.Itoa(int(time.Now().Unix()))
	} else {
		m.TimeStamp = OptionValue.TimeStamp
	}
	if OptionValue.Version == "" {
		m.Version = Version
	} else {
		m.Version = OptionValue.Version
	}
	if OptionValue.TradeLimit == 0 {
		m.TradeLimit = TradeLimit
	} else {
		m.TradeLimit = OptionValue.TradeLimit
	}
	m.ReturnURL = url.QueryEscape(OptionValue.ReturnURL)
	m.NotifyURL = url.QueryEscape(OptionValue.NotifyURL)
	m.CustomerURL = url.QueryEscape(OptionValue.CustomerURL)
	m.ClientBackURL = url.QueryEscape(OptionValue.ClientBackURL)
	m.Email = url.QueryEscape(OptionValue.Email)
	m.EmailModify = OptionValue.EmailModify
	m.LoginType = OptionValue.LoginType
	m.OrderComment = OptionValue.OrderComment
	if OptionValue.TokenTerm == "" {
		m.TokenTerm = GenSonyflake()
	} else {
		m.TokenTerm = OptionValue.TokenTerm
	}
	m.TokenTermDemand = OptionValue.TokenTermDemand
	m.BankType = OptionValue.BankType

	return m
}

func (m *MPGGateWayTradeInfo) SetExpireDate(ExpireDate string) *MPGGateWayTradeInfo {
	m.ExpireDate = ExpireDate
	return m
}

func (m *MPGGateWayTradeInfo) SetCREDIT(InstFlag string, CreditRed int, UNIONPAY int) *MPGGateWayTradeInfo {
	m.CREDIT = 1
	m.InstFlag = InstFlag
	m.CreditRed = CreditRed
	m.UNIONPAY = UNIONPAY
	return m
}

func (m *MPGGateWayTradeInfo) SetANDROIDPAY() *MPGGateWayTradeInfo {
	m.ANDROIDPAY = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetSAMSUNGPAY() *MPGGateWayTradeInfo {
	m.SAMSUNGPAY = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetLINEPAY(ImageUrl string) *MPGGateWayTradeInfo {
	m.LINEPAY = 1
	m.ImageUrl = ImageUrl
	return m
}

func (m *MPGGateWayTradeInfo) SetVACC() *MPGGateWayTradeInfo {
	m.VACC = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetCVS() *MPGGateWayTradeInfo {
	m.CVS = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetBARCODE() *MPGGateWayTradeInfo {
	m.BARCODE = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetTAIWANPAY() *MPGGateWayTradeInfo {
	m.TAIWANPAY = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetCVSCOM(LgsType string) *MPGGateWayTradeInfo {
	m.CVSCOM = 1
	m.LgsType = LgsType
	return m
}

func (m *MPGGateWayTradeInfo) SetEZPAY() *MPGGateWayTradeInfo {
	m.EZPAY = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetEZPWECHAT() *MPGGateWayTradeInfo {
	m.EZPWECHAT = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetESUNWALLET() *MPGGateWayTradeInfo {
	m.ESUNWALLET = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetEZPALIPAY() *MPGGateWayTradeInfo {
	m.EZPALIPAY = 1
	return m
}

func (m *MPGGateWayTradeInfo) SetWebATM() *MPGGateWayTradeInfo {
	m.WEBATM = WEBATM_Y
	return m
}

func (m *MPGGateWayTradeInfo) SetNTCB(NTCBLocate string, NTCBStartDate string, NTCBEndDate string) *MPGGateWayTradeInfo {
	m.NTCB = 1
	//旅遊地區代號
	m.NTCBLocate = NTCBLocate
	//國民旅遊卡起始 日期
	m.NTCBStartDate = NTCBStartDate
	//國民旅遊卡結束 日期
	m.NTCBEndDate = NTCBEndDate

	return m
}

func (m *MPGGateWayTradeInfo) CreateOrder(MerchantOrderNo string, Amt int, ItemDesc string) *MPGGateWayTradeInfo {
	m.MerchantOrderNo = MerchantOrderNo
	m.Amt = Amt
	m.ItemDesc = ItemDesc
	return m
}

func (c *Client) MPGGateway(Data *MPGGateWayTradeInfo) *MPGGatewayRequestCall {
	Data.MerchantID = c.MerchantID
	params := StructToParamsMap(Data)
	paramStr := NewValuesFromMap(params).Encode()
	req := new(MPGGatewayRequest)
	req.MerchantID = c.MerchantID
	req.TradeInfo = Aes256(paramStr, c.HashKey, c.HashIV)
	req.TradeSha = SHA256("HashKey=" + c.HashKey + "&" + req.TradeInfo + "&HashIV=" + c.HashIV)
	req.Version = Version
	req.EncryptType = 0
	return &MPGGatewayRequestCall{
		HashIV:            c.HashIV,
		HashKey:           c.HashKey,
		MPGGatewayRequest: req,
	}
}

func (m *MPGGatewayRequestCall) Do() string {
	params := StructToParamsMap(m.MPGGatewayRequest)
	html := GenerateAutoSubmitHtmlForm(params, MPGGatewayUrl)
	return html
}

func (m *MPGGatewayRequestCall) DoTest() string {
	params := StructToParamsMap(m.MPGGatewayRequest)
	html := GenerateAutoSubmitHtmlForm(params, TestMPGGatewayUrl)
	return html
}
