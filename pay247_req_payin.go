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

func (cli *Pay247Client) Payin(req Pay247PayinReq) (*Pay247PayinRsp, error) {
	rawURL := cli.Params.PayinUrl

	params := make(map[string]string)
	mapstructure.Decode(req, &params)

	params["mch_id"] = cast.ToString(cli.Params.MerchantInfo.MerchantId)
	params["notify_url"] = cast.ToString(cli.Params.PayinNotifyUrl)
	params["timestamp"] = cast.ToString(time.Now().UnixMilli())
	params["version"] = "v3.0"
	params["return_url"] = cli.Params.ReturnUrl
	params["uuid"] = fmt.Sprintf("%d", time.Now().UnixNano())

	signStr := utils.MakeSign(params, cli.Params.MerchantInfo.MerchantKey)
	params["sign"] = signStr

	var result Pay247PayinRsp

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
		cli.logger.Infof("PSPResty#pay247#payin marshal log failed: %v", err)
	} else {
		cli.logger.Infof("PSPResty#pay247#payin:\n%s", prettyLog)
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
