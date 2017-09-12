package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"jichu/scmoxing"
	"jichu/scmain"
	"jichu/scconf"
)

var _ = gauge.Step("生成moxings", func() {
	scmoxing.Shengchengmoxings()
})

var _ = gauge.Step("生成Conf", func() {
	scconf.Shengchengconf()
})




var _ = gauge.Step("生成Main", func() {
	scmain.Shengchengmain()
})
