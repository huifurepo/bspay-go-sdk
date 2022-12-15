/**
 * 分期证书配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2PcreditCertificateConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    FileList string `json:"file_list" structs:"file_list"` // 证书文件列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2PcreditCertificateConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2PcreditCertificateConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2PcreditCertificateConfigRequest(reqParam)
}

func (bp *BsPay) V2PcreditCertificateConfigRequest(reqParam V2PcreditCertificateConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_PCREDIT_CERTIFICATE_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
