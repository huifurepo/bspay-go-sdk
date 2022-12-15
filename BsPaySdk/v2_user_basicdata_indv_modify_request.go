/**
 * 个人用户基本信息修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2UserBasicdataIndvModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2UserBasicdataIndvModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2UserBasicdataIndvModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2UserBasicdataIndvModifyRequest(reqParam)
}

func (bp *BsPay) V2UserBasicdataIndvModifyRequest(reqParam V2UserBasicdataIndvModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_USER_BASICDATA_INDV_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
