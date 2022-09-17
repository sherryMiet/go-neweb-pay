package neweb_pay

import (
	"encoding/json"
	"fmt"
)

type MPGGatewayResult struct {
	//回傳狀態
	Status string
	//回傳訊息
	MerchantID string
	//交易資料 AES 加密
	TradeInfo string
	//交易資料SHA256 加密
	TradeSha string
	//串接程式版本
	Version string
	//加密模式
	EncryptType string
}

type TradeInfo struct {
	//回傳狀態
	Status string
	//回傳訊息
	Message string
	//回傳參數
	Result Result
}
type Result struct {
	// **所有支付方式共同回傳參數**
	//商店代號
	MerchantID string
	//交易金額
	Amt int
	//藍新金流交易序號
	TradeNo string
	//商店訂單編號
	MerchantOrderNo string
	//支付方式
	PaymentType string
	//回傳格式
	RespondType string
	//支付完成時間
	PayTime string
	//交易 IP
	IP string
	//款項保管銀行
	EscrowBank string
	//**信用卡支付回傳參數（一次付清、Google Pay、Samaung Pay、國民旅遊卡、銀聯）**
	//收單金融機構
	AuthBank string
	//金融機構回應碼
	RespondCode string
	//授權碼
	Auth string
	//卡號前六碼
	Card6No string
	//卡號末四碼
	Card4No string
	//分期-期別
	Inst int
	//分期-首期金額
	InstFirst int
	//分期-每期金額
	InstEach int
	//ECI
	//1.3D 回傳值 eci=1,2,5,6，代表為 3D 交易。
	//2.若交易送至收單機構授權時已是失敗狀態，則本欄位的值會以空值回傳。
	ECI string
	//信用卡快速結帳使用狀態
	TokenUseStatus string
	//紅利折抵後實際金額
	RedAmt string
	//交易類別
	PaymentMethod string
	//外幣金額
	DCC_Amt int
	//匯率
	DCC_Rate int
	//風險匯率
	DCC_Markup int
	//幣別
	DCC_Currency string
	//幣別代碼
	DCC_Currency_Code string
	//**WEBATM、ATM 繳費回傳參數**
	//付款人金融機構代碼
	PayBankCode string
	//付款人金融機構帳號末五碼
	PayerAccount5Code string
	//**超商代碼繳費回傳參數**
	//繳費代碼
	CodeNo string
	//繳費門市類別 超商類別名稱
	StoreType string
	//繳費門市代號
	StoreID string
	//**超商條碼繳費回傳參數**
	//第一段條碼
	Barcode_1 string
	//第二段條碼
	Barcode_2 string
	//第三段條碼
	Barcode_3 string
	//付款次數
	RepayTimes string
	//繳費超商
	PayStore string
	//**超商物流回傳參數**
	//超商門市編號
	StoreCode string
	//超商門市名稱
	StoreName string
	//超商門市地址
	StoreAddr string
	//取件交易方式
	TradeType string
	//取貨人
	CVSCOMName string
	//取貨人手機號碼
	CVSCOMPhone string
	//物流寄件單號
	LgsNo string
	//物流型態
	LgsType string
	//**跨境支付回傳參數(包含簡單付電子錢包、簡單付微信支付、簡單付支付寶)**
	//跨境通路類型
	ChannelID string
	//跨境通路 交易序號
	ChannelNo string
	//**玉山 Wallet 回傳參數** **台灣 Pay 回傳參數**
	//實際付款金額
	PayAmt string
	//紅利折抵金額
	RedDisAmt string
}

func NewMPGGatewayResult() *MPGGatewayResult {
	return &MPGGatewayResult{}
}
func (m MPGGatewayResult) DecodeTradeInfo(HashKey string, HashIV string) (*TradeInfo, error) {
	info := new(TradeInfo)
	StrData := DecodeAes256(m.TradeInfo, HashKey, HashIV)
	err := json.Unmarshal([]byte(StrData), info)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	return info, nil
}
func (t *TradeInfo) FillStruct(m map[string]string) error {
	for k, v := range m {
		err := SetField(t, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
