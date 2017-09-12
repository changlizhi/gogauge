package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"jichu/scmoxing"
	"jichu/scmain"
	"jichu/scconf"
	"jichu/sccuowus"
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




var _ = gauge.Step("生成Main", func() {
	scmain.Shengchengmain()
})
