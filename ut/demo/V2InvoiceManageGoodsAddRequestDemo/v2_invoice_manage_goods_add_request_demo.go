/**
 * 开票商品新增 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2InvoiceManageGoodsAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2InvoiceManageGoodsAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2InvoiceManageGoodsAddRequest{
        // 汇付商户号
        HuifuId:"6666000123123123",
        // 商品名称
        GoodsName:"sint amet minim",
        // 税收分类编码
        // TaxCode:"test",
        // 税率
        TaxRate:"0.130",
        // 是否默认
        IsDefault:"Y",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2InvoiceManageGoodsAddRequest(dgReq)
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
    // 请求日期
    extendInfoMap["req_date"] = tool.GetCurrentDate()
    // 请求流水号
    extendInfoMap["req_seq_id"] = tool.GetReqSeqId()
    // 订单商品明细
    extendInfoMap["order_goods_detail"] = "order_goods_detail"
    // 规格型号
    extendInfoMap["goods_model"] = "id dolor "
    // 计量单位
    extendInfoMap["goods_unit"] = "单位元"
    // 单价
    extendInfoMap["goods_price"] = "1.23"
    // 含税标识
    extendInfoMap["is_price_con_tax"] = "0"
    // 优惠政策标识
    extendInfoMap["preferential_flag"] = "0"
    // 零税率标示
    extendInfoMap["zero_tax_rate_flag"] = ""
    // 增值税特殊管理
    extendInfoMap["add_tax_spec_manage"] = ""
    // 说明
    extendInfoMap["ivc_remark"] = "ivc_remark"
    return extendInfoMap
}

