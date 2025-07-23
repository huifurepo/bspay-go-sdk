/**
 * 全渠道资金管理配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiEfpconfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiEfpconfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiEfpconfigRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付id
        HuifuId:"6666000108139646",
        // 所属渠道商
        UpperHuifuId:"6666000108120249",
        // 支付手续费外扣汇付ID支付手续费外扣标记为1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;
        OutFeeHuifuid:"",
        // 全域资金开户使用的银行卡信息首次开通时必填 jsonObject格式
        OutOrderAcctCard:getE87b0a0314c148faA254100625ae9787(),
        // 全域资金开户手续费首次开通时必填 jsonObject格式
        OutOrderAcctOpenFees:get652184f66c5f43f5Acb6185e66fd53ec(),
        // 银行类型switch_state有值时需填写； ht1-华通银行，xw0-XW银行，ss0-苏商银行；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：ht1&lt;/font&gt;
        OutFundsGateId:"xw0",
        // 签约人信息switch_state为1时必填 jsonObject格式
        SignUserInfo:get99d65aafAbb3475bA282C63b7361c884(),
        // 入账来源开通全域资金时需填写；01:抖音 02:美团 03:快手 04:拼多多 05:小红书 06:淘宝/天猫/飞猪 07:微信视频号/微信小店 08:京东 09:饿了么 11:得物 12:唯品会 13:携程 14:支付宝直连 15:微信直连 16:滴滴加油 17:团油 18:通联 19:易宝 20:百度 多个逗号分隔；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01,02,05&lt;/font&gt;；
        AcctSource:"01",
        // 抖音合作证明材料入账来源包含01:抖音时必填 文件类型F535；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyCooperationProvePic:"test",
        // 美团合作证明材料入账来源包含02:美团时必填 文件类型F536；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtCooperationProvePic:"test",
        // 快手合作证明材料入账来源包含03:快手时必填 文件类型F537；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsCooperationProvePic:"test",
        // 拼多多合作证明材料入账来源包含04:拼多多时必填 文件类型F538；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddCooperationProvePic:"test",
        // 小红书合作证明材料入账来源包含05:小红书时必填 文件类型F539；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsCooperationProvePic:"test",
        // 淘宝天猫飞猪合作证明材料入账来源包含06:淘宝天猫飞猪时必填 文件类型F540；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // ZfbCooperationProvePic:"test",
        // 微信视频号合作证明材料入账来源包含07:微信视频号时必填 文件类型F541；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxCooperationProvePic:"test",
        // 京东合作证明材料入账来源包含08:京东时必填 文件类型F542；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdCooperationProvePic:"test",
        // 饿了么合作证明材料入账来源包含09:饿了么时必填 文件类型F543；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // ElmCooperationProvePic:"test",
        // 得物合作证明材料入账来源包含11:得物时必填 文件类型F591；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwCooperationProvePic:"test",
        // 唯品会合作证明材料入账来源包含12:唯品会时必填 文件类型F592；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphCooperationProvePic:"test",
        // 携程合作证明材料入账来源包含13:携程时必填 文件类型F593；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XcCooperationProvePic:"test",
        // 支付宝直连合作证明材料入账来源包含14:支付宝直连时必填 文件类型F594；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // ZfbzlCooperationProvePic:"test",
        // 微信直连合作证明材料入账来源包含15:微信直连时必填 文件类型F595；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxzlCooperationProvePic:"test",
        // 滴滴加油合作证明材料入账来源包含16:滴滴加油时必填 文件类型F596；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DdjyCooperationProvePic:"test",
        // 团油合作证明材料入账来源包含17:团油时必填 文件类型F597；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // TyCooperationProvePic:"test",
        // 通联合作证明材料入账来源包含18:通联时必填 文件类型F598；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // TlCooperationProvePic:"test",
        // 易宝合作证明材料入账来源包含19:易宝时必填 文件类型F599；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // YbCooperationProvePic:"test",
        // 全渠道资金纸质协议文件协议类型为纸质时必填，文件类型F605；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // EfpPaperAgreementFile:"test",
        // 百度合作证明材料入账来源包含20:百度时必填 文件类型F616；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // BdCooperationProvePic:"test",
        // 主店商户号是否店群为是时必填
        // MainStoreHuifuId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiEfpconfigRequest(dgReq)
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
    // 开关
    extendInfoMap["switch_state"] = "1"
    // 支付手续费百分比
    extendInfoMap["fee_rate"] = "0.04"
    // 支付手续费最小值
    extendInfoMap["fee_min_amt"] = ""
    // 支付手续费外扣账户类型
    extendInfoMap["out_fee_acct_type"] = ""
    // 支付手续费外扣标记
    extendInfoMap["out_fee_flag"] = "2"
    // 业务模式
    extendInfoMap["business_model"] = "acquiringMode"
    // 全渠道资金管理补充材料id
    extendInfoMap["other_payment_institutions_pic"] = "8c4f6254-6c36-3b3c-ae8b-efcf24ca215e"
    // XW银行数字证书及电子签名授权委托书
    // extendInfoMap["xw_digital_certificate_pic"] = ""
    // 异步消息接收地址
    extendInfoMap["async_return_url"] = "http://service.example.com/to/path"
    // 业务开通结果异步消息接收地址
    extendInfoMap["busi_async_return_url"] = "http://service.example.com/to/path"
    // 申请单笔限额
    extendInfoMap["pay_every_deal"] = ""
    // 申请单日限额
    extendInfoMap["pay_every_day"] = ""
    // 全域资金分账规则
    // extendInfoMap["efp_spb_config"] = getAca242e9E64f48f9A44f2606e1ed17a0()
    // 客户ip地址
    // extendInfoMap["ip_address"] = ""
    // 是否线上场景
    // extendInfoMap["online_scene_flag"] = ""
    // 网店网址
    // extendInfoMap["online_store_website"] = ""
    // 网店名称
    // extendInfoMap["online_store_name"] = ""
    // 是否加盟连锁
    // extendInfoMap["franchise_chain_flag"] = ""
    // 交易异步应答地址
    // extendInfoMap["recon_resp_addr"] = ""
    // 协议类型
    // extendInfoMap["agreement_type"] = ""
    // 全域资金取现手续费配置
    // extendInfoMap["efp_encash_fee_config"] = get80fe7a4a0d904b3c95a8F96d865b9d5a()
    // 全域资金付款手续费配置
    // extendInfoMap["efp_payment_fee_config"] = getF4fd618742464a4b8148Ec80d19d7180()
    // 纸质协议开始日期
    // extendInfoMap["agree_begin_date"] = ""
    // 纸质协议结束日期
    // extendInfoMap["agree_end_date"] = ""
    // 是否店群
    // extendInfoMap["store_group_flag"] = ""
    // 商户经营类型
    // extendInfoMap["store_busi_type"] = ""
    // 行业类型
    // extendInfoMap["store_industry_type"] = ""
    // 经营信息材料
    // extendInfoMap["management_file"] = ""
    // 全域资金分账手续费配置
    // extendInfoMap["efp_spb_fee_config"] = getEe78625d1bf14f97A6a9Ced2395345e6()
    return extendInfoMap
}

func getE87b0a0314c148faA254100625ae9787() string {
    dto := make(map[string]interface{})
    // 结算账户名
    dto["card_name"] = "圆务铁白事"
    // 银行卡号
    dto["card_no"] = "4340622119959288"
    // 卡类型
    dto["card_type"] = "0"
    // 银行卡绑定手机号
    dto["mp"] = "13777842539"
    // 开户许可证核准号对公卡必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：J2900123456789&lt;/font&gt;
    dto["open_licence_no"] = "123456789"
    // 银行所在省
    dto["prov_id"] = "310000"
    // 银行所在市
    dto["area_id"] = "310100"
    // 银行编码
    dto["bank_code"] = "01050000"
    // 支行联行号
    dto["branch_code"] = "105290071113"
    // 支行名称
    dto["branch_name"] = "中国建设银行股份有限公司上海市中支行"
    // 持卡人证件有效期（起始）
    dto["cert_begin_date"] = "20240314"
    // 持卡人证件有效期（截止）
    dto["cert_end_date"] = ""
    // 持卡人证件号码
    dto["cert_no"] = "340000199810170714"
    // 持卡人证件类型
    dto["cert_type"] = "00"
    // 持卡人证件有效期类型
    dto["cert_validity_type"] = "1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get652184f66c5f43f5Acb6185e66fd53ec() string {
    dto := make(map[string]interface{})
    // 开户固定手续费(元)
    dto["fee_fix_amt"] = "0"
    // 开户手续费外扣时的账户类型
    dto["out_fee_acct_type"] = ""
    // 开户手续费外扣汇付ID
    dto["out_fee_huifuid"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get99d65aafAbb3475bA282C63b7361c884() string {
    dto := make(map[string]interface{})
    // 签约人类型
    dto["type"] = "LEGAL"
    // 签约人手机号
    dto["mobile_no"] = "13777842539"
    // 签约人姓名签约人类型为OTHER时必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    // dto["name"] = "test"
    // 签约人身份证签约人类型为OTHER时必填 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：321012313213222133&lt;/font&gt;
    // dto["cert_no"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1f83599f94ab4cf9Aebd4de126d4d028() string {
    dto := make(map[string]interface{})
    // 分账接收方汇付ID
    // dto["huifu_id"] = "test"
    // 分账接收方账户类型
    // dto["acct_type"] = "test"
    // 分账比例(百分比)
    // dto["split_bill_rate"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getAca242e9E64f48f9A44f2606e1ed17a0() string {
    dto := make(map[string]interface{})
    // 分账规则来源
    // dto["rule_origin"] = "test"
    // 全域资金分账手续费外扣标记1:外扣 2:内扣 规则来源为01时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1&lt;/font&gt;；
    // dto["out_fee_flag"] = "test"
    // 全域资金分账内扣规则01-商户承担02-分账接收方承担 交易手续费外扣标记为2时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01&lt;/font&gt;；
    // dto["in_fee_rule"] = "test"
    // 全域资金分账手续费外扣汇付ID交易手续费外扣标记为1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;；
    // dto["out_fee_huifuid"] = "test"
    // 全域资金分账手续费外扣账户类型交易手续费外扣标记为1时必填 01-基本户05-充值户 09-营销户；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01&lt;/font&gt;；
    // dto["out_fee_acct_type"] = "test"
    // 分账规则明细规则来源为01时必填 jsonArray格式 最多7条
    // dto["rule_detail"] = get1f83599f94ab4cf9Aebd4de126d4d028()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get80fe7a4a0d904b3c95a8F96d865b9d5a() string {
    dto := make(map[string]interface{})
    // 全域资金取现手续费百分比
    // dto["fee_rate"] = "test"
    // 全域资金取现手续费固定值
    // dto["fee_fix_amt"] = "test"
    // 全域资金取现手续费收取类型
    // dto["fee_charge_type"] = "test"
    // 全域资金取现手续费内外扣标记
    // dto["out_fee_flag"] = "test"
    // 全域资金取现手续费外扣汇付ID全域资金取现手续费内外扣标记为1:外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;
    // dto["out_fee_huifuid"] = "test"
    // 全域资金取现手续费外扣子账户号全域资金取现手续费内外扣标记为1:外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：F00598600&lt;/font&gt;
    // dto["out_fee_acctid"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF4fd618742464a4b8148Ec80d19d7180() string {
    dto := make(map[string]interface{})
    // 全域资金付款手续费百分比
    // dto["fee_rate"] = "test"
    // 全域资金付款手续费固定值
    // dto["fee_fix_amt"] = "test"
    // 全域资金付款手续费收取类型
    // dto["fee_charge_type"] = "test"
    // 全域资金付款手续费内外扣标记
    // dto["out_fee_flag"] = "test"
    // 全域资金付款手续费外扣汇付ID全域资金付款手续费内外扣标记为1:外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812123&lt;/font&gt;
    // dto["out_fee_huifuid"] = "test"
    // 全域资金付款手续费外扣子账户号全域资金付款手续费内外扣标记为1:外扣时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：F00598600&lt;/font&gt;
    // dto["out_fee_acctid"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getEe78625d1bf14f97A6a9Ced2395345e6() string {
    dto := make(map[string]interface{})
    // 全域资金分账配置开关
    // dto["switch_state"] = "test"
    // 全域资金分账手续费百分比全域资金分账配置开关为开时必填，0-100之间的数,&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.21&lt;/font&gt;
    // dto["fee_rate"] = "test"
    // 全域资金分账手续费固定值全域资金分账配置开关为开时必填，整数位最多12位，小数位最多2位；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
    // dto["fee_fix_amt"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

