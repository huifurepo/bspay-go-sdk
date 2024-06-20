/**
 * 用户业务入驻 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2UserBusiOpenRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2UserBusiOpenRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2UserBusiOpenRequest{
        // 汇付ID
        HuifuId:"6666000105765113",
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商/商户汇付Id
        UpperHuifuId:"6666000003084836",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2UserBusiOpenRequest(dgReq)
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
    // 结算信息配置
    extendInfoMap["settle_config"] = getSettleConfig()
    // 结算卡信息
    extendInfoMap["card_info"] = getCardInfo()
    // 取现配置列表
    extendInfoMap["cash_config"] = getCashConfig()
    // 文件列表
    extendInfoMap["file_list"] = getFileList()
    // 延迟入账开关
    // extendInfoMap["delay_flag"] = ""
    // 斗拱e账户功能配置
    // extendInfoMap["elec_acct_config"] = getElecAcctConfig()
    // 灵活用工开关
    // extendInfoMap["open_tax_flag"] = ""
    // 异步请求地址
    extendInfoMap["async_return_url"] = ""
    return extendInfoMap
}

func getSettleConfig() string {
    dto := make(map[string]interface{})
    // 结算周期
    dto["settle_cycle"] = "D1"
    // 结算手续费外扣时的汇付ID外扣手续费承担方的汇付ID。外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    dto["out_settle_huifuid"] = ""
    // 结算手续费外扣时的账户类型外扣手续费账户类型； 01：基本户（为空时默认值）， 05：充值户；外扣时必填；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01&lt;/font&gt;
    dto["out_settle_acct_type"] = ""
    // 结算批次号settle_pattern为P0时必填；[参见结算批次说明](https://paas.huifu.com/partners/api/#/csfl/api_csfl_jspc)
    dto["settle_batch_no"] = "300"
    // 是否优先到账settle_pattern为P0时选填， Y：是 N：否（为空默认取值）；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Y&lt;/font&gt;
    dto["is_priority_receipt"] = ""
    // 自定义结算处理时间settle_pattern为P1时必填，注意：00:00到00:30不能指定；格式：HHmmss；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：103000&lt;/font&gt;
    dto["settle_time"] = ""
    // 节假日结算手续费率settle_cycle为D1时必填。单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；settle_cycle&#x3D;T1时，不生效 ；settle_cycle为D1时，遇节假日按此费率结算 ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;
    dto["fixed_ratio"] = "2"
    // 起结金额
    dto["min_amt"] = "0.01"
    // 留存金额
    dto["remained_amt"] = "10.00"
    // 结算摘要
    dto["settle_abstract"] = "结算摘要"
    // 手续费外扣标记
    dto["out_settle_flag"] = "2"
    // 结算方式
    dto["settle_pattern"] = ""
    // 工作日结算手续费率
    // dto["workday_fixed_ratio"] = ""
    // 工作日结算手续费固定金额
    // dto["workday_constant_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCardInfo() string {
    dto := make(map[string]interface{})
    // 卡类型
    dto["card_type"] = "0"
    // 卡户名
    dto["card_name"] = "交通银行股份有限公司"
    // 卡号
    dto["card_no"] = "6217001210064762121"
    // 银行所在省
    dto["prov_id"] = "310000"
    // 银行所在市
    dto["area_id"] = "310100"
    // 银行号当card_type&#x3D;0时必填，对私可以为空[点击查看](https://paas.huifu.com/partners/api/#/csfl/api_csfl_yhbm)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01040000&lt;/font&gt;
    dto["bank_code"] = "01050000"
    // 支行联行号当card_type&#x3D;0时必填，[点击查看](https://paas.huifu.com/partners/api/#/csfl/api_csfl_yhzhbm)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：103124075619&lt;/font&gt;
    dto["branch_code"] = "105305264815"
    // 持卡人证件类型对私必填；参见《[自然人证件类型](https://paas.huifu.com/partners/api/#/api_ggcsbm?id&#x3D;%e8%87%aa%e7%84%b6%e4%ba%ba%e8%af%81%e4%bb%b6%e7%b1%bb%e5%9e%8b)》说明；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00&lt;/font&gt;
    dto["cert_type"] = "00"
    // 持卡人证件号码对私必填； 如:证件类型为身份证, 则填写身份证号码；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：320926198412032059&lt;/font&gt;
    dto["cert_no"] = "110101197003077513"
    // 持卡人证件有效期类型对私必填；1：长期有效；0：非长期有效；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0&lt;/font&gt;
    dto["cert_validity_type"] = "0"
    // 持卡人证件有效期（起始）对私必填；日期格式：yyyyMMdd，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20110112&lt;/font&gt;
    dto["cert_begin_date"] = "20210806"
    // 持卡人证件有效期（截止）当cert_validity_type&#x3D;0时必须填写；日期格式yyyyMMdd，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20110112&lt;/font&gt;&lt;br/&gt;当cert_validity_type&#x3D;1可不填
    dto["cert_end_date"] = "20410806"
    // 银行卡绑定手机号
    dto["mp"] = "15556622368"
    // 默认结算卡标志
    // dto["is_settle_default"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCashConfig() string {
    dto := make(map[string]interface{})
    // 提现手续费（固定/元）fix_amt与fee_rate至少填写一项， 需保留小数点后两位，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;注：当cash_type&#x3D;D1时为节假日取现手续费
    dto["fix_amt"] = "0.03"
    // 提现手续费率（%）fix_amt与fee_rate至少填写一项，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;注：1、如果fix_amt与fee_rate都填写了则手续费&#x3D;fix_amt+支付金额\*fee_rate2、当cash_type&#x3D;D1时为节假日取现手续费
    dto["fee_rate"] = "2"
    // D1工作日取现手续费固定金额单位元，需保留小数点后两位。不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;cash_type&#x3D;D1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
    // dto["weekday_fix_amt"] = "test"
    // D1工作日取现手续费率单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;cash_type&#x3D;D1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
    // dto["weekday_fee_rate"] = "test"
    // 业务类型
    dto["cash_type"] = "D1"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = ""
    // 手续费承担方
    dto["out_fee_huifu_id"] = ""
    // 交易手续费外扣的账户类型
    dto["out_fee_acct_type"] = ""
    // 是否优先到账
    // dto["is_priority_receipt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F02"
    // 文件jfileID
    dto["file_id"] = "71da066c-5d15-3658-a86d-4e85ee67808a"
    // 文件名称
    dto["file_name"] = "企业营业执照1.jpg"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getElecCardList() string {
    dto := make(map[string]interface{})
    // 银行编码
    // dto["bank_code"] = "test"
    // 支行联行号
    // dto["branch_code"] = "test"
    // 支行名称
    // dto["branch_name"] = "test"
    // 结算账户名
    // dto["card_name"] = "test"
    // 银行卡号
    // dto["card_no"] = "test"
    // 卡类型
    // dto["card_type"] = "test"
    // 银行所在省
    // dto["prov_id"] = ""
    // 银行所在市
    // dto["area_id"] = ""
    // 银行绑定手机号
    // dto["mp"] = ""
    // 默认卡标识
    // dto["default_cash_flag"] = ""
    // 用户授权协议版本号
    // dto["auth_version"] = ""
    // 用户授权协议号
    // dto["auth_no"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getElecAcctConfig() string {
    dto := make(map[string]interface{})
    // 电子账户开关
    // dto["switch_state"] = "test"
    // 账户类型
    // dto["acct_type"] = "test"
    // 电子账户提现手续费承担方
    // dto["cash_fee_party"] = "test"
    // 场景
    // dto["scene"] = "test"
    // 角色类型(角色编号)
    // dto["role_type"] = "test"
    // 银行卡信息
    // dto["elec_card_list"] = getElecCardList()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

