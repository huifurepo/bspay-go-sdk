/*
 * @Author: huaze.fan
 * @LastEditTime: 2022-11-14 16:14:14
 * @LastEditors: huaze.fan
 * @Description: Good Good Study！Day Day Up
 */
package BsPaySdk

type BsPay struct {
	IsProdMode bool   // true 生产，false 测试
	FileDir    string // config 配置文件路径

	Msc *MerchSysConfig // 商户配置
}

func NewBsPay(isProdMode bool, fileDir string) (b *BsPay, e error) {

	msc, err := ReadMerchConfig(fileDir)
	if err != nil {
		BspayPrintln("BsPay init error" + err.Error())
		return nil, err
	}
	BspayPrintln("BsPay init success")
	return &BsPay{
		IsProdMode: isProdMode,
		FileDir:    fileDir,
		Msc:        msc,
	}, nil
}
