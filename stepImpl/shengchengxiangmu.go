package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"jichu/scmoxing"
)

var _ = gauge.Step("生成表从jichu目录读取mhsyyonghu下的表关联", func() {
	scmoxing.Shengchengmoxings()
})
