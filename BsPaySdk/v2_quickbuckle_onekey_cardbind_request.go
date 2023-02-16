/**
 * 一键绑卡
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleOnekeyCardbindRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付Id
    OutCustId string `json:"out_cust_id" structs:"out_cust_id"` // 顾客用户号 
    BankId string `json:"bank_id" structs:"bank_id"` // 银行号
    CardName string `json:"card_name" structs:"card_name"` // 银行卡开户姓名 
    CertId string `json:"cert_id" structs:"cert_id"` // 银行卡绑定身份证
    CertType string `json:"cert_type" structs:"cert_type"` // 银行卡绑定证件类型
    CertEndDate string `json:"cert_end_date" structs:"cert_end_date"` // 证件有效期到期日长期有效不填.格式：yyyyMMdd。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20311111&lt;/font&gt;
    CardMp string `json:"card_mp" structs:"card_mp"` // 银行卡绑定手机号 
    DcType string `json:"dc_type" structs:"dc_type"` // 卡的借贷类型
    AsyncReturnUrl string `json:"async_return_url" structs:"async_return_url"` // 异步请求地址
    TrxDeviceInf string `json:"trx_device_inf" structs:"trx_device_inf"` // 设备信息域

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleOnekeyCardbindRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleOnekeyCardbindRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleOnekeyCardbindRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleOnekeyCardbindRequest(reqParam V2QuickbuckleOnekeyCardbindRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_ONEKEY_CARDBIND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
