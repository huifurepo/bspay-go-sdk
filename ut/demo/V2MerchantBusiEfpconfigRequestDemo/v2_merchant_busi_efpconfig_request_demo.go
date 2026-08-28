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
        OutOrderAcctCard:getD66709b2C1cd4c8f9a1a6e096c4c1d91(),
        // 全域资金开户手续费首次开通时必填 jsonObject格式
        OutOrderAcctOpenFees:get38be8e044bfe44b48ff96f533f80e9e9(),
        // 业务模式acquiringMode:收单模式 switch_state为1时必填
        BusinessModel:"acquiringMode",
        // 银行类型switch_state有值时需填写； ht1-华通银行，xw0-XW银行，ss0-苏商银行,wz1-微众银行&lt;font color&#x3D;&quot;green&quot;&gt;示例值：ht1&lt;/font&gt;
        OutFundsGateId:"xw0",
        // 签约人信息switch_state为1时必填 jsonObject格式
        SignUserInfo:get074b4059011f41f5B6a97432ad51a199(),
        // 入账来源开通全域资金时需填写；01:抖音 02:美团 03:快手 04:拼多多 05:小红书 06:淘宝/天猫/飞猪 07:微信视频号/微信小店 08:京东 09:饿了么 11:得物 12:唯品会 13:携程 14:支付宝直连 15:微信直连 16:滴滴加油 17:团油 18:通联 19:易宝 20:百度 21:顺丰22:希音23:高德 24:贝壳 多个逗号分隔；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01,02,05&lt;/font&gt;；
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
        // 顺丰合作证明材料入账来源包含21:顺丰时必填 文件类型F618；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // SfCooperationProvePic:"test",
        // 希音合作证明材料入账来源包含22:希音时必填 文件类型F619；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XyCooperationProvePic:"test",
        // 高德合作证明材料入账来源包含23:高德时必填 文件类型F615；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // GdCooperationProvePic:"test",
        // 贝壳合作证明材料入账来源包含24:贝壳时必填 文件类型F802；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // BkCooperationProvePic:"test",
        // 抖音平台后台交易数据截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F730；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyPlatformTransDataPic:"test",
        // 抖音平台后台店铺信息截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F731；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyPlatformStoreInfoPic:"test",
        // 抖音平台后台经营主体截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F732；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyPlatformBusiEntPic:"test",
        // 抖音平台后台订单明细截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F733；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyPlatformOrderDetailPic:"test",
        // 抖音平台后台资金流水截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F734；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyPlatformFundFlowPic:"test",
        // 抖音移动端店铺橱窗截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F735；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyMobileStoreShowcasePic:"test",
        // 抖音移动端商家资质截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F736；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyMobileBusiQualificationPic:"test",
        // 抖音移动端店铺详情截屏开通微众银行且入账来源包含01:抖音时必填 文件类型F737；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DyMobileStoreDetailPic:"test",
        // 快手平台后台交易数据截屏开通微众银行且入账来源包含03:快手时必填 文件类型F738；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsPlatformTransDataPic:"test",
        // 快手平台后台店铺信息截屏开通微众银行且入账来源包含03:快手时必填 文件类型F739；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsPlatformStoreInfoPic:"test",
        // 快手平台后台经营主体截屏开通微众银行且入账来源包含03:快手时必填 文件类型F740；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsPlatformBusiEntPic:"test",
        // 快手平台后台订单明细截屏开通微众银行且入账来源包含03:快手时必填 文件类型F741；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsPlatformOrderDetailPic:"test",
        // 快手平台后台资金流水截屏开通微众银行且入账来源包含03:快手时必填 文件类型F742；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsPlatformFundFlowPic:"test",
        // 快手移动端店铺橱窗截屏开通微众银行且入账来源包含03:快手时必填 文件类型F743；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsMobileStoreShowcasePic:"test",
        // 快手移动端商家资质截屏开通微众银行且入账来源包含03:快手时必填 文件类型F744；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsMobileBusiQualificationPic:"test",
        // 快手移动端店铺详情截屏开通微众银行且入账来源包含03:快手时必填 文件类型F745；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // KsMobileStoreDetailPic:"test",
        // 拼多多平台后台交易数据截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F746；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddPlatformTransDataPic:"test",
        // 拼多多平台后台店铺信息截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F747；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddPlatformStoreInfoPic:"test",
        // 拼多多平台后台经营主体截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F748；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddPlatformBusiEntPic:"test",
        // 拼多多平台后台订单明细截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F749；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddPlatformOrderDetailPic:"test",
        // 拼多多平台后台资金流水截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F750；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddPlatformFundFlowPic:"test",
        // 拼多多移动端店铺橱窗截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F751；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddMobileStoreShowcasePic:"test",
        // 拼多多移动端商家资质截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F752；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddMobileBusiQualificationPic:"test",
        // 拼多多移动端店铺详情截屏开通微众银行且入账来源包含04:拼多多时必填 文件类型F753；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // PddMobileStoreDetailPic:"test",
        // 小红书平台后台交易数据截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F754；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsPlatformTransDataPic:"test",
        // 小红书平台后台店铺信息截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F755；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsPlatformStoreInfoPic:"test",
        // 小红书平台后台经营主体截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F756；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsPlatformBusiEntPic:"test",
        // 小红书平台后台订单明细截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F757；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsPlatformOrderDetailPic:"test",
        // 小红书平台后台资金流水截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F758；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsPlatformFundFlowPic:"test",
        // 小红书移动端店铺橱窗截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F759；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsMobileStoreShowcasePic:"test",
        // 小红书移动端商家资质截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F760；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsMobileBusiQualificationPic:"test",
        // 小红书移动端店铺详情截屏开通微众银行且入账来源包含05:小红书时必填 文件类型F761；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // XhsMobileStoreDetailPic:"test",
        // 微信视频号平台后台交易数据截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F762；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphPlatformTransDataPic:"test",
        // 微信视频号平台后台店铺信息截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F763；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphPlatformStoreInfoPic:"test",
        // 微信视频号平台后台经营主体截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F764；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphPlatformBusiEntPic:"test",
        // 微信视频号平台后台订单明细截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F765；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphPlatformOrderDetailPic:"test",
        // 微信视频号平台后台资金流水截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F766；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphPlatformFundFlowPic:"test",
        // 微信视频号移动端店铺橱窗截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F767；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphMobileStoreShowcasePic:"test",
        // 微信视频号移动端商家资质截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F768；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphMobileBusiQualificationPic:"test",
        // 微信视频号移动端店铺详情截屏开通微众银行且入账来源包含07:微信视频号时必填 文件类型F769；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WxsphMobileStoreDetailPic:"test",
        // 京东平台后台交易数据截屏开通微众银行且入账来源包含08:京东时必填 文件类型F770；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdPlatformTransDataPic:"test",
        // 京东平台后台店铺信息截屏开通微众银行且入账来源包含08:京东时必填 文件类型F771；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdPlatformStoreInfoPic:"test",
        // 京东平台后台经营主体截屏开通微众银行且入账来源包含08:京东时必填 文件类型F772；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdPlatformBusiEntPic:"test",
        // 京东平台后台订单明细截屏开通微众银行且入账来源包含08:京东时必填 文件类型F773；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdPlatformOrderDetailPic:"test",
        // 京东平台后台资金流水截屏开通微众银行且入账来源包含08:京东时必填 文件类型F774；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdPlatformFundFlowPic:"test",
        // 京东移动端店铺橱窗截屏开通微众银行且入账来源包含08:京东时必填 文件类型F775；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdMobileStoreShowcasePic:"test",
        // 京东移动端商家资质截屏开通微众银行且入账来源包含08:京东时必填 文件类型F776；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdMobileBusiQualificationPic:"test",
        // 京东移动端店铺详情截屏开通微众银行且入账来源包含08:京东时必填 文件类型F777；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // JdMobileStoreDetailPic:"test",
        // 得物平台后台交易数据截屏开通微众银行且入账来源包含11:得物时必填 文件类型F778；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwPlatformTransDataPic:"test",
        // 得物平台后台店铺信息截屏开通微众银行且入账来源包含11:得物时必填 文件类型F779；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwPlatformStoreInfoPic:"test",
        // 得物平台后台经营主体截屏开通微众银行且入账来源包含11:得物时必填 文件类型F780；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwPlatformBusiEntPic:"test",
        // 得物平台后台订单明细截屏开通微众银行且入账来源包含11:得物时必填 文件类型F781；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwPlatformOrderDetailPic:"test",
        // 得物平台后台资金流水截屏开通微众银行且入账来源包含11:得物时必填 文件类型F782；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwPlatformFundFlowPic:"test",
        // 得物移动端店铺橱窗截屏开通微众银行且入账来源包含11:得物时必填 文件类型F783；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwMobileStoreShowcasePic:"test",
        // 得物移动端商家资质截屏开通微众银行且入账来源包含11:得物时必填 文件类型F784；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwMobileBusiQualificationPic:"test",
        // 得物移动端店铺详情截屏开通微众银行且入账来源包含11:得物时必填 文件类型F785；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // DwMobileStoreDetailPic:"test",
        // 美团平台后台交易数据截屏开通微众银行且入账来源包含02:美团时必填 文件类型F786；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtPlatformTransDataPic:"test",
        // 美团平台后台店铺信息截屏开通微众银行且入账来源包含02:美团时必填 文件类型F787；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtPlatformStoreInfoPic:"test",
        // 美团平台后台经营主体截屏开通微众银行且入账来源包含02:美团时必填 文件类型F788；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtPlatformBusiEntPic:"test",
        // 美团平台后台订单明细截屏开通微众银行且入账来源包含02:美团时必填 文件类型F789；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtPlatformOrderDetailPic:"test",
        // 美团平台后台资金流水截屏开通微众银行且入账来源包含02:美团时必填 文件类型F790；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtPlatformFundFlowPic:"test",
        // 美团移动端店铺橱窗截屏开通微众银行且入账来源包含02:美团时必填 文件类型F791；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtMobileStoreShowcasePic:"test",
        // 美团移动端商家资质截屏开通微众银行且入账来源包含02:美团时必填 文件类型F792；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtMobileBusiQualificationPic:"test",
        // 美团移动端店铺详情截屏开通微众银行且入账来源包含02:美团时必填 文件类型F793；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // MtMobileStoreDetailPic:"test",
        // 唯品会平台后台交易数据截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F794；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphPlatformTransDataPic:"test",
        // 唯品会平台后台店铺信息截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F795；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphPlatformStoreInfoPic:"test",
        // 唯品会平台后台经营主体截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F796；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphPlatformBusiEntPic:"test",
        // 唯品会平台后台订单明细截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F797；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphPlatformOrderDetailPic:"test",
        // 唯品会平台后台资金流水截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F798；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphPlatformFundFlowPic:"test",
        // 唯品会移动端店铺橱窗截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F799；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphMobileStoreShowcasePic:"test",
        // 唯品会移动端商家资质截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F800；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphMobileBusiQualificationPic:"test",
        // 唯品会移动端店铺详情截屏开通微众银行且入账来源包含12:唯品会时必填 文件类型F801；详见[文件类型说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_wjlx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // WphMobileStoreDetailPic:"test",
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
    // extendInfoMap["efp_spb_config"] = get9c68bbd739af493e9027D6d2e4e8f4bb()
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
    // extendInfoMap["efp_encash_fee_config"] = get1990df443ba348bd9768107e2612cc1d()
    // 全域资金付款手续费配置
    // extendInfoMap["efp_payment_fee_config"] = get6030b0527e524d37Ad86Ff411b9a78d1()
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
    // extendInfoMap["efp_spb_fee_config"] = get034ef171B86844e8A7deD5b30093e02a()
    return extendInfoMap
}

func getD66709b2C1cd4c8f9a1a6e096c4c1d91() string {
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

func get38be8e044bfe44b48ff96f533f80e9e9() string {
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

func get074b4059011f41f5B6a97432ad51a199() string {
    dto := make(map[string]interface{})
    // 签约人类型
    dto["type"] = "LEGAL"
    // 签约人手机号
    dto["mobile_no"] = "13777842539"
    // 签约人姓名
    // dto["name"] = ""
    // 签约人身份证
    // dto["cert_no"] = ""
    // 签约人邮箱
    // dto["email"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get76777ca5097c46b385342a0bc037e15a() string {
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

func get9c68bbd739af493e9027D6d2e4e8f4bb() string {
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
    // dto["rule_detail"] = get76777ca5097c46b385342a0bc037e15a()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1990df443ba348bd9768107e2612cc1d() string {
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

func get6030b0527e524d37Ad86Ff411b9a78d1() string {
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

func get034ef171B86844e8A7deD5b30093e02a() string {
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

