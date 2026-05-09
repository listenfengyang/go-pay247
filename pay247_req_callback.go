package go_pay247

import (
	"fmt"

	"github.com/listenfengyang/go-pay247/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// 充值回调
func (cli *Pay247Client) VerifyPayinCallback(req Pay247PayinCallbackReq, processor func(Pay247PayinCallbackReq) error) error {
	// 手动映射字段，使用JSON tag名称
	params := make(map[string]interface{})
	mapstructure.Decode(req, &params)
	// fmt.Printf("req params: %v\n", params)

	var params2 map[string]string
	params2 = make(map[string]string)
	params2["created_at"] = cast.ToString(params["created_at"].(int64))
	params2["paid_at"] = cast.ToString(params["paid_at"].(int64))
	mapstructure.Decode(params, &params2)

	fmt.Printf("params2: %v\n", params2)

	if !utils.Pay247VerifyCallback(params2, cli.Params.MerchantInfo.MerchantKey) {
		cli.logger.Errorf("pay247 payin callback back verify fail, req: %v", params2)
	}

	return processor(req)
}

// 退款回调
func (cli *Pay247Client) VerifyPayoutCallback(req Pay247PayoutCallbackReq, processor func(Pay247PayoutCallbackReq) error) error {
	// 手动映射字段，使用JSON tag名称
	params := make(map[string]string)
	params["mch_id"] = req.MchId
	params["mch_order_no"] = req.MchOrderNo
	mapstructure.Decode(req, &params)

	//签名校验失败
	if !utils.Pay247VerifyCallback(params, cli.Params.MerchantInfo.MerchantKey) {
		cli.logger.Errorf("pay247 payout callback back verify fail, req: %v", params)
	}

	return processor(req)
}
