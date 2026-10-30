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
    FileUrl string `json:"file_url" structs:"file_url"` // 文件url链接文件url链接与file文件流不能同时上传；文件支持类型：JPG,BMP,PNG；单个图片最大支持2M；除去图片外其他类型文件支持10M&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：[https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/imgs/guide/jinjian/%E8%90%A5%E4%B8%9A%E6%89%A7%E7%85%A71-2.png](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/imgs/guide/jinjian/%E8%90%A5%E4%B8%9A%E6%89%A7%E7%85%A71-2.png)

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
