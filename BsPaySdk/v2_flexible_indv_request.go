/**
 * 灵工个人用户进件
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2FlexibleIndvRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商/商户汇付Id
    BasicInfo string `json:"basic_info" structs:"basic_info"` // 基本信息
    CardInfo string `json:"card_info" structs:"card_info"` // 卡信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2FlexibleIndvRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2FlexibleIndvRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2FlexibleIndvRequest(reqParam)
}

func (bp *BsPay) V2FlexibleIndvRequest(reqParam V2FlexibleIndvRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_FLEXIBLE_INDV
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
