package go_pay247

type Pay247InitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	PayinUrl        string `json:"payinUrl" mapstructure:"payinUrl" config:"payinUrl" yaml:"payinUrl"`
	PayinNotifyUrl  string `json:"payinNotifyUrl" mapstructure:"payinNotifyUrl" config:"payinNotifyUrl" yaml:"payinNotifyUrl"`
	PayoutUrl       string `json:"payoutUrl" mapstructure:"payoutUrl" config:"payoutUrl" yaml:"payoutUrl"`
	PayoutNotifyUrl string `json:"payoutNotifyUrl" mapstructure:"payoutNotifyUrl" config:"payoutNotifyUrl" yaml:"payoutNotifyUrl"`
	ReturnUrl       string `json:"returnUrl" mapstructure:"returnUrl" config:"returnUrl" yaml:"returnUrl"`
}

type MerchantInfo struct {
	MerchantId   string `json:"mch_id" mapstructure:"mch_id" config:"mch_id" yaml:"mch_id"`
	MerchantCode string `json:"merchantCode" mapstructure:"merchantCode" config:"merchantCode" yaml:"merchantCode"`
	MerchantKey  string `json:"mch_key" mapstructure:"mch_key" config:"mch_key" yaml:"mch_key"`
}

type Pay247PayinReq struct {
	MchId      string `json:"mch_id" mapstructure:"mch_id"`
	MchOrderNo string `json:"mch_order_no" mapstructure:"mch_order_no"`
	MchUserId  string `json:"mch_user_id" mapstructure:"mch_user_id"`
	Currency   string `json:"currency" mapstructure:"currency"`
	Amount     string `json:"amount" mapstructure:"amount"`
	PayMethod  string `json:"pay_method,omitempty" mapstructure:"pay_method"`
	PayTheme   string `json:"pay_theme,omitempty" mapstructure:"pay_theme"`
	ClientIp   string `json:"client_ip" mapstructure:"client_ip"`
	NotifyUrl  string `json:"notify_url" mapstructure:"notify_url"`
	ReturnUrl  string `json:"return_url,omitempty" mapstructure:"return_url"`
	Timestamp  string `json:"timestamp" mapstructure:"timestamp"`
	Version    string `json:"version" mapstructure:"version"`
	Uuid       string `json:"uuid" mapstructure:"uuid"`
	Sign       string `json:"sign" mapstructure:"sign"`
	// PayerId         string `json:"payer_id,omitempty" mapstructure:"payer_id"`
	// PayerName       string `json:"payer_name,omitempty" mapstructure:"payer_name"`
	// PayerPhone      string `json:"payer_phone,omitempty" mapstructure:"payer_phone"`
	// PayerEmail      string `json:"payer_email,omitempty" mapstructure:"payer_email"`
	// PayerAccountNo  string `json:"payer_account_no,omitempty" mapstructure:"payer_account_no"`
	// PayerAddress    string `json:"payer_address,omitempty" mapstructure:"payer_address"`
	// PayerBankCode   string `json:"payer_bank_code,omitempty" mapstructure:"payer_bank_code"`
	// PayerBranchCode string `json:"payer_branch_code,omitempty" mapstructure:"payer_branch_code"`
}

type Pay247PayinRsp struct {
	Code      int             `json:"code" mapstructure:"code"`
	Message   string          `json:"message" mapstructure:"message"`
	Uuid      string          `json:"uuid" mapstructure:"uuid"`
	Timestamp int64           `json:"timestamp" mapstructure:"timestamp"`
	Data      Pay247PayinData `json:"data" mapstructure:"data"`
}

type Pay247PayinData struct {
	MchId      string          `json:"mch_id" mapstructure:"mch_id"`
	MchOrderNo string          `json:"mch_order_no" mapstructure:"mch_order_no"`
	OrderNo    string          `json:"order_no" mapstructure:"order_no"`
	Currency   string          `json:"currency" mapstructure:"currency"`
	Amount     string          `json:"amount" mapstructure:"amount"`
	PayMethod  string          `json:"pay_method" mapstructure:"pay_method"`
	PayTheme   string          `json:"pay_theme" mapstructure:"pay_theme"`
	Status     string          `json:"status" mapstructure:"status"`
	PayUrl     string          `json:"pay_url,omitempty" mapstructure:"pay_url"`
	PayParams  Pay247PayParams `json:"pay_params,omitempty" mapstructure:"pay_params"`
}

type Pay247PayParams struct {
	PayeeName       string `json:"payee_name,omitempty" mapstructure:"payee_name"`
	PayeeAccountNo  string `json:"payee_account_no,omitempty" mapstructure:"payee_account_no"`
	PayeeBankName   string `json:"payee_bank_name,omitempty" mapstructure:"payee_bank_name"`
	PayeeBranchName string `json:"payee_branch_name,omitempty" mapstructure:"payee_branch_name"`
	Qr              string `json:"qr,omitempty" mapstructure:"qr"`
	Utr             string `json:"utr,omitempty" mapstructure:"utr"`
}

type Pay247PayinCallbackReq struct {
	MchId        string `json:"mch_id" form:"mch_id" mapstructure:"mch_id"`
	MchOrderNo   string `json:"mch_order_no" form:"mch_order_no" mapstructure:"mch_order_no"`
	OrderNo      string `json:"order_no" form:"order_no" mapstructure:"order_no"`
	Currency     string `json:"currency" form:"currency" mapstructure:"currency"`
	Amount       string `json:"amount" form:"amount" mapstructure:"amount"`
	PayMethod    string `json:"pay_method" form:"pay_method" mapstructure:"pay_method"`
	Status       string `json:"status" form:"status" mapstructure:"status"`
	CreatedAt    int64  `json:"created_at" form:"created_at" mapstructure:"created_at"`
	Fee          string `json:"fee" form:"fee" mapstructure:"fee"`
	ActualAmount string `json:"actual_amount" form:"actual_amount" mapstructure:"actual_amount"`
	PaidAt       int64  `json:"paid_at" form:"paid_at" mapstructure:"paid_at"`
	Sign         string `json:"sign" form:"sign" mapstructure:"sign"`
	Error        string `json:"error,omitempty" form:"error,omitempty" mapstructure:"error"`
}

type Pay247PayoutReq struct {
	MchId       string `json:"mch_id" mapstructure:"mch_id"`
	MchOrderNo  string `json:"mch_order_no" mapstructure:"mch_order_no"`
	Currency    string `json:"currency" mapstructure:"currency"`
	Amount      string `json:"amount" mapstructure:"amount"`
	PayMethod   string `json:"pay_method" mapstructure:"pay_method"`
	AccountName string `json:"account_name,omitempty" mapstructure:"account_name"`
	AccountNo   string `json:"account_no" mapstructure:"account_no"`
	BankCode    string `json:"bank_code,omitempty" mapstructure:"bank_code"`
	// BankBranch  string `json:"bank_branch,omitempty" mapstructure:"bank_branch"`
	NotifyUrl string `json:"notify_url,omitempty" mapstructure:"notify_url"`
	Timestamp int64  `json:"timestamp" mapstructure:"timestamp"`
	Version   string `json:"version" mapstructure:"version"`
	Uuid      string `json:"uuid" mapstructure:"uuid"`
	Sign      string `json:"sign" mapstructure:"sign"`
}

type Pay247PayoutRsp struct {
	Code      int              `json:"code" mapstructure:"code"`
	Message   string           `json:"message" mapstructure:"message"`
	Uuid      string           `json:"uuid" mapstructure:"uuid"`
	Timestamp int64            `json:"timestamp" mapstructure:"timestamp"`
	Data      Pay247PayoutData `json:"data" mapstructure:"data"`
}

type Pay247PayoutData struct {
	MchId      string `json:"mch_id" mapstructure:"mch_id"`
	MchOrderNo string `json:"mch_order_no" mapstructure:"mch_order_no"`
	OrderNo    string `json:"order_no" mapstructure:"order_no"`
	Currency   string `json:"currency" mapstructure:"currency"`
	Amount     string `json:"amount" mapstructure:"amount"`
	PayMethod  string `json:"pay_method" mapstructure:"pay_method"`
	Fee        string `json:"fee" mapstructure:"fee"`
	Status     string `json:"status" mapstructure:"status"`
	PaidAt     string `json:"paid_at,omitempty" mapstructure:"paid_at"`
	Error      string `json:"error,omitempty" mapstructure:"error"`
}

type Pay247PayoutCallbackReq struct {
	MchId      string `json:"mch_id" form:"mch_id" mapstructure:"mch_id"`
	MchOrderNo string `json:"mch_order_no" form:"mch_order_no" mapstructure:"mch_order_no"`
	OrderNo    string `json:"order_no" form:"order_no" mapstructure:"order_no"`
	Currency   string `json:"currency" form:"currency" mapstructure:"currency"`
	Amount     string `json:"amount" form:"amount" mapstructure:"amount"`
	PayMethod  string `json:"pay_method" form:"pay_method" mapstructure:"pay_method"`
	Fee        string `json:"fee" form:"fee" mapstructure:"fee"`
	Status     string `json:"status" form:"status" mapstructure:"status"`
	PaidAt     string `json:"paid_at" form:"paid_at" mapstructure:"paid_at"`
	Error      string `json:"error" form:"error" mapstructure:"error"`
	Timestamp  int64  `json:"timestamp" form:"timestamp" mapstructure:"timestamp"`
	Sign       string `json:"sign" form:"sign" mapstructure:"sign"`
}
