/**
 * 企业用户基本信息修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserBasicdataEntModifyRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    LegalCertNationality string `json:"legal_cert_nationality" structs:"legal_cert_nationality"` // 法人国籍法人的证件类型为外国人居留证时，必填，参见《[国籍编码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/area/%E5%9B%BD%E7%B1%8D.xlsx)》&lt;font color&#x3D;&quot;green&quot;&gt;示例值：CHN&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserBasicdataEntModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserBasicdataEntModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserBasicdataEntModifyRequest(reqParam)
}

func (bp *BsPay) V2UserBasicdataEntModifyRequest(reqParam V2UserBasicdataEntModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_BASICDATA_ENT_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
