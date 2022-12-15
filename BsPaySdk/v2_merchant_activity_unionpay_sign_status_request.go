/**
 * 银联活动报名进度查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantActivityUnionpaySignStatusRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    EnlistId string `json:"enlist_id" structs:"enlist_id"` // 报名编号与serialNumber二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：521724026796785664&lt;/font&gt;
    SerialNumber string `json:"serial_number" structs:"serial_number"` // 报名请求流水号报名时传递的reqSysId；与enlistId二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：ZDTESTrQ202011054108473959671&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantActivityUnionpaySignStatusRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantActivityUnionpaySignStatusRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantActivityUnionpaySignStatusRequest(reqParam)
}

func (bp *BsPay) V2MerchantActivityUnionpaySignStatusRequest(reqParam V2MerchantActivityUnionpaySignStatusRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_ACTIVITY_UNIONPAY_SIGN_STATUS
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
