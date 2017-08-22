package stepImpl
import(
	"github.com/getgauge-contrib/gauge-go/gauge"
	"hanfuxin/baseinits"
	"github.com/sclevine/agouti"
	"log"
)

var _ = gauge.Step("just print cuowu",func(){
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()

	page, err := agoutiDriver.NewPage()
	if err != nil{
		log.Println(err)
		panic(err)
	}

	page.Navigate("http://www.baidu.com")
	page.Find("#kw").Fill("love yannuo")
	page.Find("#su").Click()
	page.Destroy()
	agoutiDriver.Stop()

	log.Println(baseinits.Cuowus)
})
