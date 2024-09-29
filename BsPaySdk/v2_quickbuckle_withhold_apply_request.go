/**
 * 代扣绑卡申请
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleWithholdApplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付Id
    ReturnUrl string `json:"return_url" structs:"return_url"` // 返回地址
    OutCustId string `json:"out_cust_id" structs:"out_cust_id"` // 用户id
    OrderId string `json:"order_id" structs:"order_id"` // 绑卡订单号
    OrderDate string `json:"order_date" structs:"order_date"` // 绑卡订单日期
    CardId string `json:"card_id" structs:"card_id"` // 银行卡号
    CardName string `json:"card_name" structs:"card_name"` // 银行卡开户姓名 
    CertType string `json:"cert_type" structs:"cert_type"` // 银行卡绑定证件类型 
    CertId string `json:"cert_id" structs:"cert_id"` // 银行卡绑定身份证
    CardMp string `json:"card_mp" structs:"card_mp"` // 银行卡绑定手机号 
    VipCode string `json:"vip_code" structs:"vip_code"` // CVV2信用卡代扣专用 需要密文传输，需要密文传输，使用汇付RSA公钥加密(加密前64位，加密后最长2048位），参见[参考文档](https://paas.huifu.com/open/doc/guide/#/api_jiami_jiemi)；
    Expiration string `json:"expiration" structs:"expiration"` // 卡有效期 信用卡代扣专用，格式：MMYY 需要密文传输，使用汇付RSA公钥加密(加密前64位，加密后最长2048位），参见[参考文档](https://paas.huifu.com/open/doc/guide/#/api_jiami_jiemi)；
    CertValidityType string `json:"cert_validity_type" structs:"cert_validity_type"` // 个人证件有效期类型
    CertBeginDate string `json:"cert_begin_date" structs:"cert_begin_date"` // 个人证件有效期起始日
    CertEndDate string `json:"cert_end_date" structs:"cert_end_date"` // 个人证件有效期到期日长期有效不填.格式：YYYYMMDD；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20450112&lt;/font&gt;
    DcType string `json:"dc_type" structs:"dc_type"` // 卡的借贷类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleWithholdApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleWithholdApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleWithholdApplyRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleWithholdApplyRequest(reqParam V2QuickbuckleWithholdApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_WITHHOLD_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
