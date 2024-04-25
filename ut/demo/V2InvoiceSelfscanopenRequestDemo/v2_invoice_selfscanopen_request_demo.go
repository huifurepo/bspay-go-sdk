/**
 * 自助扫码开票 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2InvoiceSelfscanopenRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2InvoiceSelfscanopenRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2InvoiceSelfscanopenRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000103675282",
        // 发票类型
        // IvcType:"test",
        // 开票类型
        // OpenType:"test",
        // 含税合计金额（元）
        // OrderAmt:"test",
        // 开票商品信息
        // GoodsInfos:getGoodsInfosRc(),
        // 开票人信息
        // PayerInfo:getPayerInfo(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2InvoiceSelfscanopenRequest(dgReq)
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
    // 税控盘编号
    // extendInfoMap["tax_device_id"] = ""
    // 备注
    // extendInfoMap["resv"] = ""
    // 特殊票种标识
    // extendInfoMap["special_flag"] = ""
    // 开票结果异步通知地址
    // extendInfoMap["callback_url"] = ""
    return extendInfoMap
}

func getGoodsInfosRc() string {
    dto := make(map[string]interface{})
    // 发票行性质
    // dto["ivc_nature"] = "test"
    // 商品名称goods_code不为空时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：电视机&lt;/font&gt;
    // dto["goods_name"] = "test"
    // 税率goods_code不为空时必填，最多三位小数 如：税率13% 则传入&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.13&lt;/font&gt;
    // dto["tax_rate"] = "test"
    // 金额（元）
    // dto["trans_amt"] = "test"
    // 商品id
    // dto["goods_id"] = ""
    // 商品税收分类编码
    // dto["goods_code"] = ""
    // 规格型号
    // dto["goods_model"] = ""
    // 计量单位
    // dto["goods_unit"] = ""
    // 优惠政策标识
    // dto["preferential_flag"] = ""
    // 零税率标示
    // dto["zero_tax_rate_flag"] = ""
    // 增值税特殊管理
    // dto["add_tax_spec_manage"] = ""
    // 含税标识
    // dto["is_price_con_tax"] = ""
    // 数量
    // dto["goods_count"] = ""
    // 单价
    // dto["goods_price"] = ""
    // 折扣金额(元)
    // dto["sale_amt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getPayerInfo() string {
    dto := make(map[string]interface{})
    // 开票人
    // dto["payer_name"] = "test"
    // 收款人
    // dto["payee"] = ""
    // 复核人
    // dto["reviewer"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

