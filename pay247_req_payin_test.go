package go_pay247

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestPay247Payin(t *testing.T) {

	vLog := VLog{}
	cli := NewPay247Client(vLog, &Pay247InitParams{
		PayinUrl:        PAY247_PAYIN_URL,
		PayoutUrl:       PAY247_PAYOUT_URL,
		PayinNotifyUrl:  PAY247_PAYIN_NOTIFY_URL,
		PayoutNotifyUrl: PAY247_PAYOUT_NOTIFY_URL,
		ReturnUrl:       RETURN_URL,
	})

	cli.Params.MerchantInfo = MerchantInfo{
		MerchantId:  MERCHANT_ID,
		MerchantKey: MERCHANT_KEY,
	}

	resp, err := cli.Payin(GenPay247PayinRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenPay247PayinRequestDemo() Pay247PayinReq {
	return Pay247PayinReq{
		MchOrderNo: "TRADE-4051890",
		MchUserId:  "U405189",
		Currency:   "USD", // "KHR",
		Amount:     "99.11",
		PayMethod:  "EWALLET",
		PayTheme:   "link",
		ClientIp:   "127.0.0.1",
	}
}
