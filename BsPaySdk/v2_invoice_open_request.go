/**
 * 发票开具
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceOpenRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号huifu_id与ext_mer_id二选一必填，汇付商户号优先；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;
    ExtMerId string `json:"ext_mer_id" structs:"ext_mer_id"` // 外部商户号&lt;font color&#x3D;&quot;green&quot;&gt;示例值：&lt;/font&gt;
    ChannelId string `json:"channel_id" structs:"channel_id"` // 渠道号汇付商户号为空时，必传；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812124&lt;/font&gt;
    IvcType string `json:"ivc_type" structs:"ivc_type"` // 发票类型
    OpenType string `json:"open_type" structs:"open_type"` // 开票类型
    BuyerName string `json:"buyer_name" structs:"buyer_name"` // 购方单位名称
    OrderAmt string `json:"order_amt" structs:"order_amt"` // 含税合计金额(元)
    RedApplyReason string `json:"red_apply_reason" structs:"red_apply_reason"` // 冲红原因open_type&#x3D;1时必填01：开票有误02：销货退回03：服务终止04：销售转让
    RedApplySource string `json:"red_apply_source" structs:"red_apply_source"` // 冲红申请来源open_type&#x3D;1时必填01：销方02：购方
    OriIvcCode string `json:"ori_ivc_code" structs:"ori_ivc_code"` // 原发票代码openType&#x3D;1时必填；参见[发票右上角](https://paas.huifu.com/open/doc/api/#/fp/api_fp_yanglitu.md)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：144032209110&lt;/font&gt;
    OriIvcNumber string `json:"ori_ivc_number" structs:"ori_ivc_number"` // 原发票号码openType&#x3D;1时必填；参见[发票右上角](https://paas.huifu.com/open/doc/api/#/fp/api_fp_yanglitu.md)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20685767&lt;/font&gt;
    GoodsInfos string `json:"goods_infos" structs:"goods_infos"` // 开票商品信息
    PayerInfo string `json:"payer_info" structs:"payer_info"` // 开票人信息
    EstateSales string `json:"estate_sales" structs:"estate_sales"` // 不动产销售特殊字段specialFlag为05时，必填；jsonArray格式
    EstateLease string `json:"estate_lease" structs:"estate_lease"` // 不动产租赁特殊字段specialFlag为16时，必填；jsonArray格式

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceOpenRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceOpenRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceOpenRequest(reqParam)
}

func (bp *BsPay) V2InvoiceOpenRequest(reqParam V2InvoiceOpenRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_OPEN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
