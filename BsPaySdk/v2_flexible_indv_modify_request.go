/**
 * 灵工个人用户修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2FlexibleIndvModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商/商户汇付Id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2FlexibleIndvModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2FlexibleIndvModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2FlexibleIndvModifyRequest(reqParam)
}

func (bp *BsPay) V2FlexibleIndvModifyRequest(reqParam V2FlexibleIndvModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_FLEXIBLE_INDV_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
