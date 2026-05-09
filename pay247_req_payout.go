package go_pay247

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/listenfengyang/go-pay247/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

func (cli *Pay247Client) Payout(req Pay247PayoutReq) (*Pay247PayoutRsp, error) {
	rawURL := cli.Params.PayoutUrl

	params := make(map[string]string)
	mapstructure.Decode(req, &params)

	params["mch_id"] = cast.ToString(cli.Params.MerchantInfo.MerchantId)
	params["notify_url"] = cast.ToString(cli.Params.PayoutNotifyUrl)
	params["timestamp"] = cast.ToString(time.Now().UnixMilli())
	params["version"] = "v3.0"
	params["uuid"] = fmt.Sprintf("%d", time.Now().UnixNano())

	signStr := utils.MakeSign(params, cli.Params.MerchantInfo.MerchantKey)
	params["sign"] = signStr

	fmt.Printf("params:%+v\n", params)

	var result Pay247PayoutRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	prettyLog, err := json.MarshalIndent(utils.GetRestyLog(resp2), "", "  ")
	if err != nil {
		cli.logger.Infof("PSPResty#pay247#payout marshal log failed: %v", err)
	} else {
		cli.logger.Infof("PSPResty#pay247#payout:\n%s", prettyLog)
	}

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
