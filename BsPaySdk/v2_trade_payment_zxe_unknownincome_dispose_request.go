/**
 * 不明来账处理
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentZxeUnknownincomeDisposeRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BankSerialNo string `json:"bank_serial_no" structs:"bank_serial_no"` // 银行侧交易流水号参照异步通知里的bank_serial_no；&lt;br/&gt;“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
    PayAcct string `json:"pay_acct" structs:"pay_acct"` // 来账银行账号需要密文传输，使用汇付RSA公钥加密(加密前64位，加密后最长2048位），参见[参考文档](https://paas.huifu.com/open/doc/guide/#/api_jiami_jiemi)；示例值：Ly+fnExeyPOTzfOtgRRur77nJB9TAe4PGgK9M……fc6XJXZss&#x3D;“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
    PayAcctName string `json:"pay_acct_name" structs:"pay_acct_name"` // 来账账户名称“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
    TransDate string `json:"trans_date" structs:"trans_date"` // 交易日期“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
    OperateType string `json:"operate_type" structs:"operate_type"` // 操作类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentZxeUnknownincomeDisposeRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentZxeUnknownincomeDisposeRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentZxeUnknownincomeDisposeRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentZxeUnknownincomeDisposeRequest(reqParam V2TradePaymentZxeUnknownincomeDisposeRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_ZXE_UNKNOWNINCOME_DISPOSE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
