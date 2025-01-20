/**
 * 全渠道资金管理配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiEfpconfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 所属渠道商
    SwitchState string `json:"switch_state" structs:"switch_state"` // 开关
    OutOrderAutoAcctFlag string `json:"out_order_auto_acct_flag" structs:"out_order_auto_acct_flag"` // 自动入账开关0:关闭 1:开通；switch_state为1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1&lt;/font&gt;
    OutFeeHuifuid string `json:"out_fee_huifuid" structs:"out_fee_huifuid"` // 支付手续费外扣汇付ID支付手续费外扣标记为1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;
    OutOrderAcctCard string `json:"out_order_acct_card" structs:"out_order_acct_card"` // 全域资金开户使用的银行卡信息首次开通时必填 jsonObject格式
    OutOrderAcctOpenFees string `json:"out_order_acct_open_fees" structs:"out_order_acct_open_fees"` // 全域资金开户手续费首次开通时必填 jsonObject格式
    OtherPaymentInstitutionsPic string `json:"other_payment_institutions_pic" structs:"other_payment_institutions_pic"` // 全渠道资金管理补充材料id涉及文件类型：[F504-全渠道资金管理补充材料](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    XwDigitalCertificatePic string `json:"xw_digital_certificate_pic" structs:"xw_digital_certificate_pic"` // XW银行数字证书及电子签名授权委托书out_funds_gate_id为xw0时必填；涉及文件类型：[F534-银行数字证书及电子签名授权委托书](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    OutFundsGateId string `json:"out_funds_gate_id" structs:"out_funds_gate_id"` // 银行类型
    SignUserInfo string `json:"sign_user_info" structs:"sign_user_info"` // 签约人信息switch_state为1时必填 jsonObject格式
    AcctSource string `json:"acct_source" structs:"acct_source"` // 入账来源

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiEfpconfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiEfpconfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiEfpconfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiEfpconfigRequest(reqParam V2MerchantBusiEfpconfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_EFPCONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
