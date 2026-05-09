package go_pay247

import (
	"testing"
)

func TestPay247Payout(t *testing.T) {
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

	resp, err := cli.Payout(GenPay247PayoutRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenPay247PayoutRequestDemo() Pay247PayoutReq {
	return Pay247PayoutReq{
		MchOrderNo:  "R898543254325432",
		Currency:    "CNY",
		Amount:      "99.11",
		PayMethod:   "BANK",
		AccountName: "Test Account",
		AccountNo:   "1234567890",
		BankCode:    "BDO",
	}
}
