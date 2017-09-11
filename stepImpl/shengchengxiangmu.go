package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"jichu/scmoxing"
	"jichu/scmain"
)

var _ = gauge.Step("生成moxings", func() {
	scmoxing.Shengchengmoxings()
})

var _ = gauge.Step("生成Main", func() {
	scmain.Shengchengmain()
})
