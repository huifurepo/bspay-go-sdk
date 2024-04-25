/**
 * 一键绑卡-工行卡号查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleBindCardinfoQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付Id
    ProductId string `json:"product_id" structs:"product_id"` // 产品Id
    CardName string `json:"card_name" structs:"card_name"` // 银行卡开户姓名
    CertType string `json:"cert_type" structs:"cert_type"` // 身份证类型
    CertNo string `json:"cert_no" structs:"cert_no"` // 银行卡绑定身份证
    CardMobile string `json:"card_mobile" structs:"card_mobile"` // 银行卡绑定手机号
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 回调地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleBindCardinfoQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleBindCardinfoQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleBindCardinfoQueryRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleBindCardinfoQueryRequest(reqParam V2QuickbuckleBindCardinfoQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_BIND_CARDINFO_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
