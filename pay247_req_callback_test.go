package go_pay247

import (
	"testing"
)

func TestPay247VerifyPayinCallback(t *testing.T) {
	vLog := VLog{}
	cli := NewPay247Client(vLog, &Pay247InitParams{
		PayinUrl:        PAY247_PAYIN_URL,
		PayoutUrl:       PAY247_PAYOUT_URL,
		PayinNotifyUrl:  PAY247_PAYIN_NOTIFY_URL,
		PayoutNotifyUrl: PAY247_PAYOUT_NOTIFY_URL,
	})

	cli.Params.MerchantInfo = MerchantInfo{
		MerchantId:  MERCHANT_ID,
		MerchantKey: MERCHANT_KEY,
	}

	callbackReq := Pay247PayinCallbackReq{
		MchId:        MERCHANT_ID,
		MchOrderNo:   "TRADE-405189",
		OrderNo:      "PIUSD1037820260508155807319155",
		Currency:     "USD",
		PayMethod:    "EWALLET",
		Status:       "SUCCESS",
		CreatedAt:    1778227087000,
		Sign:         "7cadc3e6dc0097c1330a3038009b402c",
		Amount:       "99.1100",
		Fee:          "0.00",
		ActualAmount: "99.1100",
		PaidAt:       1778227193000,
	}

	isValid := cli.VerifyPayinCallback(callbackReq, func(req Pay247PayinCallbackReq) error {
		return nil
	})
	cli.logger.Infof("Payin callback verification result: %v\n", isValid)
}

func TestPay247VerifyPayoutCallback(t *testing.T) {
	vLog := VLog{}
	cli := NewPay247Client(vLog, &Pay247InitParams{
		PayinUrl:        PAY247_PAYIN_URL,
		PayoutUrl:       PAY247_PAYOUT_URL,
		PayinNotifyUrl:  PAY247_PAYIN_NOTIFY_URL,
		PayoutNotifyUrl: PAY247_PAYOUT_NOTIFY_URL,
	})

	cli.Params.MerchantInfo = MerchantInfo{
		MerchantId:  MERCHANT_ID,
		MerchantKey: MERCHANT_KEY,
	}

	// {\"mch_id\":\"PJ0PF5SO7H10378\",\"order_no\":\"POUSD1037820260513150657534935\",\"mch_order_no\":\"202605131006230205\",\"currency\":\"USD\",
	// \"pay_method\":\"BANK\",\"amount\":\"8025.0000\",\"fee\":\"160.5000\",\"status\":\"SUCCESS\",
	// \"receipt\":\"\",\"paid_at\":1778656383000,\"created_at\":1778656017000,\"error\":\"\",\"sign\":\"c9baaeb487c52ab3dc0f1dbf4d665edb\"}
	callbackReq := Pay247PayoutCallbackReq{
		MchId:      MERCHANT_ID,
		MchOrderNo: "202605131006230205",
		OrderNo:    "POUSD1037820260513150657534935",
		Currency:   "USD",
		Amount:     "8025.0000",
		PayMethod:  "BANK",
		Fee:        "160.5000",
		Status:     "SUCCESS",
		PaidAt:     1778656383000,
		CreatedAt:  1778656017000,
		Error:      "",
		Receipt:    "",
		Sign:       "c9baaeb487c52ab3dc0f1dbf4d665edb",
	}

	isValid := cli.VerifyPayoutCallback(callbackReq, func(req Pay247PayoutCallbackReq) error {
		return nil
	})
	cli.logger.Infof("Payout callback verification result: %v\n", isValid)
}

func GenPay247PayinCallbackDemo() Pay247PayinCallbackReq {
	return Pay247PayinCallbackReq{
		MchId:      MERCHANT_ID,
		MchOrderNo: "TRADE-405189",
		OrderNo:    "MO202308282232593669984082",
		Currency:   "CNY",
		Amount:     "99.11",
		PayMethod:  "GCASH",
		Status:     "SUCCESS",
		CreatedAt:  1693233334134,
		Sign:       "",
	}
}

func GenPay247PayoutCallbackDemo() Pay247PayoutCallbackReq {
	return Pay247PayoutCallbackReq{
		MchId:      MERCHANT_ID,
		MchOrderNo: "R898543254325432",
		OrderNo:    "202007070427475133333",
		Currency:   "PHP",
		Amount:     "99.11",
		PayMethod:  "BANK",
		Fee:        "3.60",
		Status:     "SUCCESS",
		PaidAt:     1594223431233,
		CreatedAt:  1594099906123,
		Sign:       "",
	}
}
