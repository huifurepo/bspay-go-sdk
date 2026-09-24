/**
 * 发货信息录入
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradeHostingPaymentShippinginfoUploadRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrderKey string `json:"order_key" structs:"order_key"` // 订单信息
    LogisticsType string `json:"logistics_type" structs:"logistics_type"` // 物流模式
    DeliveryMode string `json:"delivery_mode" structs:"delivery_mode"` // 发货模式
    IsAllDelivered string `json:"is_all_delivered" structs:"is_all_delivered"` // 否已全部发货完成分拆发货模式时必填，用于标识分拆发货模式下是否已全部发货完成，只有全部发货完成的情况下才会向用户推送发货完成通知。&lt;br/&gt;枚举值: true，false&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：true&lt;/font&gt;
    UploadTime string `json:"upload_time" structs:"upload_time"` // 上传时间

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradeHostingPaymentShippinginfoUploadRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradeHostingPaymentShippinginfoUploadRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradeHostingPaymentShippinginfoUploadRequest(reqParam)
}

func (bp *BsPay) V3TradeHostingPaymentShippinginfoUploadRequest(reqParam V3TradeHostingPaymentShippinginfoUploadRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_HOSTING_PAYMENT_SHIPPINGINFO_UPLOAD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
