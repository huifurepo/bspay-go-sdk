/**
 * 商户公域授权
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LinkappAuthDoRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    PlatformType string `json:"platform_type" structs:"platform_type"` // 平台类型
    ContractUrl string `json:"contract_url" structs:"contract_url"` // 协议地址
    ContractMerName string `json:"contract_mer_name" structs:"contract_mer_name"` // 签约商户名称
    ContractTime string `json:"contract_time" structs:"contract_time"` // 签约时间
    PhoneNumber string `json:"phone_number" structs:"phone_number"` // 登录用手机号第一次登录有需要手机验证码的情况;（需要授权手机安装一个转发短信的应用）
    MerchantType string `json:"merchant_type" structs:"merchant_type"` // 商户类型商户类型：0个人店 1企业 2个体工商户 3其他(目前固定填3即可)

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LinkappAuthDoRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LinkappAuthDoRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LinkappAuthDoRequest(reqParam)
}

func (bp *BsPay) V2LinkappAuthDoRequest(reqParam V2LinkappAuthDoRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LINKAPP_AUTH_DO
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
