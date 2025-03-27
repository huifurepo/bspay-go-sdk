/**
 * 图片上传
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2SupplementaryPictureRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 业务请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 业务请求日期
    FileType string `json:"file_type" structs:"file_type"` // 图片类型
    Picture string `json:"picture" structs:"picture"` // 图片名称

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2SupplementaryPictureRequest(reqStr string, filepath string) (map[string]interface{}, error) {
	reqParam := V2SupplementaryPictureRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2SupplementaryPictureRequest(reqParam, filepath)
}

func (bp *BsPay) V2SupplementaryPictureRequest(reqParam V2SupplementaryPictureRequest, filepath string) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_SUPPLEMENTARY_PICTURE
    return UploadBsPay(reqUrl, ToMap(reqParam), filepath, bp.Msc)
}
