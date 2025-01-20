/**
 * 银行大额支付绑卡
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LargeamtBindcardBindRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    CardType string `json:"card_type" structs:"card_type"` // 卡类型
    CardName string `json:"card_name" structs:"card_name"` // 银行账户名
    CardNo string `json:"card_no" structs:"card_no"` // 银行卡号密文
    BankCode string `json:"bank_code" structs:"bank_code"` // 银行编码
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 手机号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LargeamtBindcardBindRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LargeamtBindcardBindRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LargeamtBindcardBindRequest(reqParam)
}

func (bp *BsPay) V2LargeamtBindcardBindRequest(reqParam V2LargeamtBindcardBindRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LARGEAMT_BINDCARD_BIND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
