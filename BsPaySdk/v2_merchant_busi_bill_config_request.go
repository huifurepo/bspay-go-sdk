/**
 * 交易结算对账文件配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiBillConfigRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付机构编号
    ReconSendFlag string `json:"recon_send_flag" structs:"recon_send_flag"` // 对账文件生成开关
    FileType string `json:"file_type" structs:"file_type"` // 对账单类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiBillConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiBillConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiBillConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiBillConfigRequest(reqParam V2MerchantBusiBillConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_BILL_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
