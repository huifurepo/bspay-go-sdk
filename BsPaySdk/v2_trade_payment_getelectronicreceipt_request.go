/**
 * 电子回单查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentGetelectronicreceiptRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ShowFeemat string `json:"show_feemat" structs:"show_feemat"` // 是否展示手续费
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 交易返回的全局流水号交易返回的全局流水号。org_hf_seq_id与（org_req_seq_id、org_req_date、pay_type） 不能同时为空；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：003500TOP2B211021163242P447ac132fd200000&lt;/font&gt;
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期格式:yyyyMMdd；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20221022&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号org_hf_seq_id与（org_req_seq_id、org_req_date、pay_type） 不能同时为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2022012614120615001&lt;/font&gt;
    PayType string `json:"pay_type" structs:"pay_type"` // 支付类型BALANCE_PAY-余额支付，&lt;br/&gt;CASHOUT-取现，&lt;br/&gt;QUICK_PAY-快捷支付，&lt;br/&gt;ONLINE_PAY-网银，&lt;br/&gt;&lt;!--SURROGATE-代发&lt;br/&gt;许士通说暂不支持--&gt;PAY_CONFIRM-交易确认&lt;br/&gt;TRANSFER_ACCT-大额转账(指[银行大额转账](https://paas.huifu.com/partners/api/#/dejy/api_dejy_yhdezz)交易)&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：ONLINE_PAY&lt;/font&gt;&lt;br/&gt;注意：支付类型有值，原交易请求流水号不为空必填； &lt;br/&gt;选择交易确认类型时：请传入交易确认的请求流水号或全局流水号。

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentGetelectronicreceiptRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentGetelectronicreceiptRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentGetelectronicreceiptRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentGetelectronicreceiptRequest(reqParam V2TradePaymentGetelectronicreceiptRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_GETELECTRONICRECEIPT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
