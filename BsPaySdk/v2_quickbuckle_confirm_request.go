/**
 * 快捷绑卡确认接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleConfirmRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户Id
    OutCustId string `json:"out_cust_id" structs:"out_cust_id"` // 用户id
    OrderId string `json:"order_id" structs:"order_id"` // 订单号
    OrderDate string `json:"order_date" structs:"order_date"` // 订单日期
    CardId string `json:"card_id" structs:"card_id"` // 银行卡号
    CardName string `json:"card_name" structs:"card_name"` // 银行卡开户姓名
    CertType string `json:"cert_type" structs:"cert_type"` // 银行卡绑定证件类型
    CertId string `json:"cert_id" structs:"cert_id"` // 银行卡绑定身份证
    CardMp string `json:"card_mp" structs:"card_mp"` // 银行卡绑定手机号
    VerifyCode string `json:"verify_code" structs:"verify_code"` // 验证码
    VipCode string `json:"vip_code" structs:"vip_code"` // CVV2信用卡代扣专用需要密文传输，使用商户RSA私钥加密(加密前3位，加密后最长2048位），[参见参考文档](https://paas.huifu.com/partners/guide/#/api_jiami_jiemi)；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Ly+fnExeyPOTzf……rGq0l0NTebfc6XJXZss&#x3D;&lt;/font&gt;
    Expiration string `json:"expiration" structs:"expiration"` // 卡有效期信用卡代扣专用，格式：MMYY，需要密文传输。&lt;br/&gt;使用商户RSA私钥加密(加密前4位，加密后最长2048位），[参见参考文档](https://paas.huifu.com/partners/guide/#/api_jiami_jiemi)；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Ly+fnExeyPO……ebfc6XJXZss&#x3D;&lt;/font&gt;
    TrxDeviceInf string `json:"trx_device_inf" structs:"trx_device_inf"` // 设备信息域 

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleConfirmRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleConfirmRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleConfirmRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleConfirmRequest(reqParam V2QuickbuckleConfirmRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_CONFIRM
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
