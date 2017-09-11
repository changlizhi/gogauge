package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"jichu/scmoxing"
)

var _ = gauge.Step("生成moxings", func() {
	scmoxing.Shengchengmoxings()
})
