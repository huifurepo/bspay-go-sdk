/**
 * 查询批量账单数据
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayOrderBatchDetailRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BillNo string `json:"bill_no" structs:"bill_no"` // 账单编号与原创建批量账单数据请求流水号二选一必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：BN2025091279190693&lt;/font&gt;;
    OriReqSeqId string `json:"ori_req_seq_id" structs:"ori_req_seq_id"` // 原创建批量账单数据请求流水号原创建批量账单数据请求流水号，同一商户号当天唯一；与帐单编号二选一必填
    OriReqDate string `json:"ori_req_date" structs:"ori_req_date"` // 原创建批量账单数据请求日期原创建批量账单数据日期格式：yyyyMMdd，以北京时间为准；与帐单编号二选一必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayOrderBatchDetailRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayOrderBatchDetailRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayOrderBatchDetailRequest(reqParam)
}

func (bp *BsPay) V3BillpayOrderBatchDetailRequest(reqParam V3BillpayOrderBatchDetailRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_ORDER_BATCH_DETAIL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
