/**
 * 银行大额支付差错申请 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentTransferBankmistakeApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentTransferBankmistakeApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentTransferBankmistakeApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000110468104",
        // 交易金额
        TransAmt:"0.01",
        // 订单类型
        OrderType:"REFUND",
        // 原请求流水号
        OrgReqSeqId:"202308312345678931",
        // 原请求日期
        OrgReqDate:"20230831",
        // 实际打款日期
        RemitDate:"20230615",
        // 实际付款方姓名
        CertificateName:"孙洁",
        // 实际付款方银行卡号
        BankCardNo:"V2olJv4Srh…………78M8A==",
        // 实际付款方银行名称
        BankName:"招商银行",
        // 异步通知地址
        NotifyUrl:"http://www.baidu.com",
        // 商品描述
        // GoodsDesc:"test",
        // 汇款凭证文件id
        // CertificateFileId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentTransferBankmistakeApplyRequest(dgReq)
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
    // 备注
    extendInfoMap["remark"] = "大额支付补入账验证"
    // 银行信息数据
    extendInfoMap["bank_info_data"] = getBankInfoData()
    // 延时标记
    // extendInfoMap["delay_acct_flag"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = getAcctSplitBunch()
    return extendInfoMap
}

func getBankInfoData() string {
    dto := make(map[string]interface{})
    // 省份对公代发必填，[参见省市地区码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/area/%E6%96%97%E6%8B%B1%E4%BB%A3%E5%8F%91%E7%9C%81%E4%BB%BD%E5%9C%B0%E5%8C%BA%E7%BC%96%E7%A0%81.xlsx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0013&lt;/font&gt;
    dto["province"] = "0031"
    // 地区对公代发必填，[参见省市地区码](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/area/%E6%96%97%E6%8B%B1%E4%BB%A3%E5%8F%91%E7%9C%81%E4%BB%BD%E5%9C%B0%E5%8C%BA%E7%BC%96%E7%A0%81.xlsx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1301&lt;/font&gt;
    dto["area"] = "3100"
    // 银行编号
    dto["bank_code"] = "03080000"
    // 联行号选填，参见：[银行支行编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_yhzhbm)； &lt;font color&#x3D;&quot;green&quot;&gt;示例值：102290026507&lt;/font&gt;&lt;br/&gt;对私代发非必填；
    dto["correspondent_code"] = "103290076178"
    // 对公对私标识
    dto["card_acct_type"] = "P"
    // 证件类型
    dto["certificate_type"] = "01"
    // 手机号
    dto["mobile_no"] = "oO6XYz…………Is3nZb/5dFj860Z+nQ=="
    // 证件号
    dto["certify_no"] = "yL09mhS5…………WK04Kdfyg=="
    // 支行名
    dto["subbranch_bank_name"] = "中国农业银行股份有限公司上海联洋支行"
    // 付款方三证合一码
    dto["bank_acct_three_in_one"] = "92650109MA79R8E308"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAcctInfos() interface{} {
    dto := make(map[string]interface{})
    // 支付金额
    // dto["div_amt"] = ""
    // 商户号
    // dto["huifu_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunch() interface{} {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getAcctInfos()

    return dto;
}

