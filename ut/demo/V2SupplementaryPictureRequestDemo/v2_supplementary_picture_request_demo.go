/**
 * 图片上传 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2SupplementaryPictureRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"path/filepath"
)

func V2SupplementaryPictureRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2SupplementaryPictureRequest{
        // 业务请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 业务请求日期
        ReqDate:tool.GetCurrentDate(),
        // 图片类型
        FileType:"F01",
        // 图片名称
        // Picture:"test",
    }
    filePath, _ := filepath.Abs("./ut/IMG_4955.JPG")
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2SupplementaryPictureRequest(dgReq, filePath)
  	if err != nil {
		fmt.Println("请求异常:", err)
	}

	fmt.Println("返回数据:", resp)
}

/**
 * 非必填字段
 */
func getExtendInfos() map[string]interface{} {
    // 设置非必填字段
    extendInfoMap := make(map[string]interface{})
    // 商户号
    extendInfoMap["huifu_id"] = "6666000103413615"
    // 文件url链接
    extendInfoMap["file_url"] = "https://example.com/image.jpg"
    return extendInfoMap
}

