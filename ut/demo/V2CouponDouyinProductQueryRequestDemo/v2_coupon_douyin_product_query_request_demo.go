/**
 * 抖音套餐映射接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponDouyinProductQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponDouyinProductQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponDouyinProductQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000108412345",
        // 门店绑定流水号
        BindId:"7123fc6e5c564f539e73031c08912345",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponDouyinProductQueryRequest(dgReq)
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
    // 区分商品创建者的查询方式
    extendInfoMap["goods_creator_type"] = "0"
    // 商品名称
    // extendInfoMap["product_name"] = ""
    // 是否查询商品全量关联门店
    // extendInfoMap["query_all_poi"] = ""
    // 筛选在线状态
    extendInfoMap["status"] = "1"
    // 光标
    // extendInfoMap["cursor"] = ""
    // 分页数量
    // extendInfoMap["count"] = ""
    // 门店ID list
    extendInfoMap["poi_ids"] = "[23,45]"
    // 外部门店ID list
    // extendInfoMap["ext_ids"] = ""
    return extendInfoMap
}

