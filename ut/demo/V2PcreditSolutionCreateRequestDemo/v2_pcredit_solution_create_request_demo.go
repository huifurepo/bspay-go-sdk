/**
 * 创建花呗分期方案 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2PcreditSolutionCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2PcreditSolutionCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2PcreditSolutionCreateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000003084836",
        // 花呗分期商家贴息活动名称
        ActivityName:"花呗分期商家贴息活动名称",
        // 活动开始时间
        StartTime:"2019-07-08 00:00:00",
        // 活动结束时间
        EndTime:"2039-07-10 00:00:00",
        // 免息金额下限(元)
        MinMoneyLimit:"1000",
        // 免息金额上限(元)
        MaxMoneyLimit:"3000",
        // 花呗分期贴息预算金额
        AmountBudget:"60000",
        // 花呗分期数集合
        InstallNumStrList:"3",
        // 预算提醒金额(元)
        BudgetWarningMoney:"58000",
        // 预算提醒邮件列表
        BudgetWarningMailList:"111@alipay.com",
        // 预算提醒手机号列表
        BudgetWarningMobileNoList:"13940001100",
        // 子门店信息集合
        SubShopInfoList:getSubShopInfoList(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2PcreditSolutionCreateRequest(dgReq)
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
    // 开发者的应用ID
    extendInfoMap["app_id"] = ""
    return extendInfoMap
}

func getSubShopInfoList() string {
    dto := make(map[string]interface{})
    // 二级商户号
    dto["sub_mer_id"] = "A4854135335181517376"
    // 二级商户名
    dto["sub_mer_name"] = "预二人"
    // 费率
    dto["fee_type"] = "02"
    // 店铺名称
    dto["mer_name"] = "盈盈超市"
    // 省份
    dto["province"] = "浙江省"
    // 市名
    dto["city"] = "杭州市"
    // 区、县
    dto["county"] = "西湖区"
    // 地址详情
    dto["detail"] = "古荡街道西溪路556号蚂蚁Z空间"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

