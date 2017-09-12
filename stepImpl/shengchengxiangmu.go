package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"jichu/scmoxing"
	"jichu/scmain"
	"jichu/scconf"
	"jichu/sccuowus"
	"jichu/scchushihuas"
)

var _ = gauge.Step("生成Moxings", func() {
	scmoxing.Shengchengmoxings()
})

var _ = gauge.Step("生成Conf", func() {
	scconf.Shengchengconf()
})


var _ = gauge.Step("生成Cuowu", func() {
	sccuowus.Shengchengziduancuowu()
})
var _ = gauge.Step("生成Chushihuas-jsongo", func() {
	scchushihuas.Shengchengjsongo()
})
var _ = gauge.Step("生成Chushihuas-lujinghuoqu", func() {
	scchushihuas.Shengchenglujinghuoqu()
})




//需要Chushihuas这个前提
var _ = gauge.Step("生成Chushihuas-duqujson", func() {
	scchushihuas.Shengchengduqujson()
})




var _ = gauge.Step("生成Main", func() {
	scmain.Shengchengmain()
})
