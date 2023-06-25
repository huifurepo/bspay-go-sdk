/**
 * 图片上传
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeElectronReceiptsPictureUploadRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ThirdChannelType string `json:"third_channel_type" structs:"third_channel_type"` // 三方通道类型
    FileName string `json:"file_name" structs:"file_name"` // 文件名称
    ImageContent string `json:"image_content" structs:"image_content"` // 图片内容

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeElectronReceiptsPictureUploadRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeElectronReceiptsPictureUploadRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeElectronReceiptsPictureUploadRequest(reqParam)
}

func (bp *BsPay) V2TradeElectronReceiptsPictureUploadRequest(reqParam V2TradeElectronReceiptsPictureUploadRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ELECTRON_RECEIPTS_PICTURE_UPLOAD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
