/**
 * 快捷绑卡申请接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3QuickbuckleApplyRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    OutCustId string `json:"out_cust_id" structs:"out_cust_id"` // 商户用户id
    CardNo string `json:"card_no" structs:"card_no"` // 银行卡号
    CardName string `json:"card_name" structs:"card_name"` // 银行卡开户姓名
    CertNo string `json:"cert_no" structs:"cert_no"` // 银行卡绑定身份证
    CertValidityType string `json:"cert_validity_type" structs:"cert_validity_type"` // 个人证件有效期类型
    CertBeginDate string `json:"cert_begin_date" structs:"cert_begin_date"` // 个人证件有效期起始日
    CertEndDate string `json:"cert_end_date" structs:"cert_end_date"` // 个人证件有效期到期日长期有效不填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20420905&lt;/font&gt;
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 银行卡绑定手机号
    ProtocolNo string `json:"protocol_no" structs:"protocol_no"` // 挂网协议编号授权信息(招行绑卡需要上送)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：34463343&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3QuickbuckleApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3QuickbuckleApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3QuickbuckleApplyRequest(reqParam)
}

func (bp *BsPay) V3QuickbuckleApplyRequest(reqParam V3QuickbuckleApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_QUICKBUCKLE_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
