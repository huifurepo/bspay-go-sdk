/**
 * 发票开具 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2InvoiceOpenRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2InvoiceOpenRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2InvoiceOpenRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号huifu_id与ext_mer_id二选一必填，汇付商户号优先；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;
        HuifuId:"6666000107430944",
        // 外部商户号&lt;font color&#x3D;&quot;green&quot;&gt;示例值：&lt;/font&gt;
        ExtMerId:"",
        // 开票方式乐企联用模式下非斗拱交易必填 枚举： TRADE-交易即开票  PAY-支付即开票
        // IvcModeCode:"test",
        // 渠道号汇付商户号为空时，必传；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812124&lt;/font&gt;
        // ChannelId:"test",
        // 发票类型
        IvcType:"1",
        // 开票类型
        OpenType:"0",
        // 购方单位名称
        BuyerName:"张三",
        // 含税合计金额(元)
        OrderAmt:"70.00",
        // 冲红原因open_type&#x3D;1时必填01：开票有误02：销货退回03：服务终止04：销售转让
        // RedApplyReason:"test",
        // 冲红申请来源open_type&#x3D;1时必填01：销方02：购方03：第三方
        // RedApplySource:"test",
        // 原发票代码openType&#x3D;1时必填；参见[发票右上角](https://paas.huifu.com/open/doc/api/#/fp/api_fp_yanglitu.md)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：144032209110&lt;/font&gt;
        OriIvcCode:"90222082",
        // 原发票号码openType&#x3D;1时必填；参见[发票右上角](https://paas.huifu.com/open/doc/api/#/fp/api_fp_yanglitu.md)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20685767&lt;/font&gt;
        OriIvcNumber:"150000020026",
        // 开票商品信息
        GoodsInfos:getE39d35ceD1c946049c4a78abe4385b46(),
        // 不动产销售特殊字段specialFlag为05时，必填；jsonArray格式
        // EstateSales:get3e2cef346688416b829072c10a9ceadd(),
        // 不动产租赁特殊字段specialFlag为16时，必填；jsonArray格式
        // EstateLease:get47f6d2af90b840db984f379ab6303717(),
        // 交易订单信息jsonObject格式，订单信息（乐企联用，AT联用模式必填）
        // OrderInfo:get9f17cfb9260e4d8c9db3Ee08ada0cc13(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2InvoiceOpenRequest(dgReq)
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
    // 开票模式
    // extendInfoMap["issue_mode"] = ""
    // 子商户号
    // extendInfoMap["sub_mer_id"] = ""
    // 税控盘编号
    extendInfoMap["tax_device_id"] = "661919694739"
    // 购方单位识别号
    extendInfoMap["buyer_no"] = ""
    // 购方单位地址
    extendInfoMap["buyer_address"] = ""
    // 购方单位电话
    extendInfoMap["buyer_tel"] = ""
    // 购方开户行名称
    extendInfoMap["buyer_bank_name"] = ""
    // 购方银行账号
    extendInfoMap["buyer_acct_no"] = ""
    // 购方企业类型
    extendInfoMap["buyer_ent_type"] = ""
    // 收票人邮件
    extendInfoMap["rec_ivc_email"] = "test@126.com"
    // 备注
    extendInfoMap["resv"] = "备注"
    // 特殊票种标识
    extendInfoMap["special_flag"] = "00"
    // 红字信息表编号
    extendInfoMap["red_info_number"] = ""
    // 开票人信息
    extendInfoMap["payer_info"] = getCa65c4309ca442c0A892037585b8e873()
    // 开票结果异步通知地址
    extendInfoMap["callback_url"] = "virgo://http://192.168.85.157:30031/sspm/testVirgo"
    // 强制开票标识
    extendInfoMap["buyer_info_confirm"] = ""
    // 购方自然人标识
    // extendInfoMap["buyer_nature_person_type"] = ""
    // 购方自然人证件类型
    // extendInfoMap["buyer_nature_person_card_type"] = ""
    // 购方自然人证件号码
    // extendInfoMap["buyer_nature_person_card_no"] = ""
    // 购方自然人国籍代码
    // extendInfoMap["buyer_nature_person_nation_no"] = ""
    // 减按征税类型代码
    // extendInfoMap["reduced_tax_type"] = ""
    // 二手车信息
    // extendInfoMap["used_car"] = get4095e83b0a9b46eeB3305641800925eb()
    return extendInfoMap
}

func get6e4b3c323c4d4fa6Ae98414ce70f646a() interface{} {
    dto := make(map[string]interface{})
    // 加油站名称
    // dto["station"] = "test"
    // 油枪号
    // dto["gun_number"] = "test"
    // 加油时间
    // dto["oil_time"] = "test"

    return dto;
}

func getE39d35ceD1c946049c4a78abe4385b46() string {
    dto := make(map[string]interface{})
    // 发票行性质
    dto["ivc_nature"] = "0"
    // 商品序号ivc_type&#x3D;1 红票必填，要与开具的蓝票商品一致；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：备注&lt;/font&gt;
    dto["goods_serial_num"] = ""
    // 商品id二选一必填，税收分类编码优先；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：&lt;/font&gt;
    dto["goods_id"] = ""
    // 商品税收分类编码参考[商品编码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/fp/%E7%94%B5%E5%AD%90%E5%8F%91%E7%A5%A8%E5%95%86%E5%93%81%E7%BC%96%E7%A0%81.xlsx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1010101090000000000&lt;/font&gt;
    dto["goods_code"] = "6010000000000000000"
    // 商品名称goods_code不为空时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：&lt;/font&gt;
    dto["goods_name"] = "预付卡"
    // 税率goods_code不为空时必填，最多三位小数；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.130&lt;/font&gt;
    dto["tax_rate"] = "0.03"
    // 含税标识
    dto["is_price_con_tax"] = "1"
    // 金额(元)
    dto["trans_amt"] = "70.00"
    // 成品油信息乐企联用、AT乐企联用模式下成品油票种必传
    // dto["oil_info"] = get6e4b3c323c4d4fa6Ae98414ce70f646a()
    // 规格型号
    dto["goods_model"] = ""
    // 计量单位
    dto["goods_unit"] = ""
    // 优惠政策标识
    dto["preferential_flag"] = "0"
    // 零税率标示
    dto["zero_tax_rate_flag"] = ""
    // 增值税特殊管理
    dto["add_tax_spec_manage"] = ""
    // 商品数量
    dto["goods_count"] = "7"
    // 单价
    dto["goods_price"] = "10"
    // 折扣金额(元)
    dto["sale_amt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getCa65c4309ca442c0A892037585b8e873() string {
    dto := make(map[string]interface{})
    // 收款人
    dto["payee"] = "收款人"
    // 复核人
    dto["reviewer"] = "复核"
    // 开票人
    dto["payer_name"] = "开票人"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3e2cef346688416b829072c10a9ceadd() string {
    dto := make(map[string]interface{})
    // 不动产地址
    // dto["addr"] = "test"
    // 不动产详细地址
    // dto["detail_addr"] = "test"
    // 跨地（市）标志
    // dto["area_flag"] = "test"
    // 土地增值税项目编号
    // dto["tax_item_no"] = "test"
    // 不动产单元代码/网签合同备案编号
    // dto["record_no"] = "test"
    // 核定计税价格
    // dto["total_amt"] = "test"
    // 实际成交含税金额
    // dto["deal_amt"] = "test"
    // 房屋产权证书/不动产产权号
    // dto["estate_no"] = "test"
    // 不动产单位
    // dto["unit"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get47f6d2af90b840db984f379ab6303717() string {
    dto := make(map[string]interface{})
    // 不动产地址
    // dto["addr"] = "test"
    // 不动产详细地址
    // dto["detail_addr"] = "test"
    // 跨地（市）标志
    // dto["area_flag"] = "test"
    // 租赁日期起
    // dto["start_date"] = "test"
    // 租赁日期止
    // dto["end_date"] = "test"
    // 房屋产权证书/不动产产权号
    // dto["estate_no"] = "test"
    // 不动产单位
    // dto["unit"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getAbba899f2acf4386A7b6B14265184add() string {
    dto := make(map[string]interface{})
    // 支付方式
    // dto["pay_way"] = ""
    // 支付金额
    // dto["trade_amt"] = ""
    // 支付通道订单号
    // dto["trade_no"] = ""
    // 支付时间
    // dto["trade_time"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get9f17cfb9260e4d8c9db3Ee08ada0cc13() string {
    dto := make(map[string]interface{})
    // 交易订单号
    // dto["order_id"] = "test"
    // 交易金额示例值：1.00 单位元，is_dg_trans &#x3D; N，必填
    // dto["amt"] = "test"
    // 支付方式jsonArray格式，支付信息明细，is_dg_trans&#x3D;N时必填
    // dto["pay_infos"] = getAbba899f2acf4386A7b6B14265184add()
    // 交易订单日期
    // dto["order_date"] = ""
    // 是否斗拱交易
    // dto["is_dg_trans"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3cec4f3f754c4e91Ae1e5c32d003cd8c() string {
    dto := make(map[string]interface{})
    // 车牌号码
    // dto["plate_number"] = ""
    // 登记证号
    // dto["register_cert_number"] = ""
    // 厂牌型号
    // dto["factory_plate_model"] = ""
    // 车辆类型代码
    // dto["car_type"] = ""
    // 转入地车辆管理所名称
    // dto["into_vehicle_manage_name"] = ""
    // 车架号
    // dto["vehicle_identification_number"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4095e83b0a9b46eeB3305641800925eb() string {
    dto := make(map[string]interface{})
    // 二手车销售统一发票正向开票标识
    // dto["direction_type"] = ""
    // 二手车企业性质
    // dto["ent_prop_type"] = ""
    // 二手车销售统一发票号码
    // dto["unified_sales_inv_number"] = ""
    // 卖方自然人标识
    // dto["car_nature_person_type"] = ""
    // 卖方自然人证件类型
    // dto["car_nature_person_card_type"] = ""
    // 卖方自然人证件号码
    // dto["car_nature_person_card_no"] = ""
    // 卖方自然人国籍代码
    // dto["car_nature_person_nation_no"] = ""
    // 卖方单位名称/自然人姓名
    // dto["car_nature_ent_name"] = ""
    // 卖方税号
    // dto["car_ent_no"] = ""
    // 卖方地址/自然人地址
    // dto["car_nature_ent_address"] = ""
    // 卖方电话/自然人电话
    // dto["car_nature_ent_tell"] = ""
    // 车辆费基本信息集合
    // dto["car_base_list"] = get3cec4f3f754c4e91Ae1e5c32d003cd8c()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

