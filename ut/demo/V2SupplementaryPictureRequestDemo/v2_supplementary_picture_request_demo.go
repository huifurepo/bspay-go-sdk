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
        // 文件url链接文件url链接与file文件流不能同时上传；文件支持类型：JPG,BMP,PNG；单个图片最大支持2M；除去图片外其他类型文件支持10M&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：[https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/imgs/guide/jinjian/%E8%90%A5%E4%B8%9A%E6%89%A7%E7%85%A71-2.png](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/imgs/guide/jinjian/%E8%90%A5%E4%B8%9A%E6%89%A7%E7%85%A71-2.png)
        FileUrl:"https://example.com/image.jpg",
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
    return extendInfoMap
}

