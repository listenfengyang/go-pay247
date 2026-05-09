package go_pay247

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-pay247/utils"
)

type Pay247Client struct {
	Params *Pay247InitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewPay247Client(logger utils.Logger, params *Pay247InitParams) *Pay247Client {
	return &Pay247Client{
		Params: params,

		ryClient:  resty.New(),
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Pay247Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}

func (cli *Pay247Client) SetMerchantInfo(merchant MerchantInfo) {
	cli.Params.MerchantInfo = MerchantInfo{
		MerchantId:  merchant.MerchantId,
		MerchantKey: merchant.MerchantKey,
	}
}
