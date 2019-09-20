package getSoftLink
import (
	"github.com/gocolly/colly"
	"strings"
)
// 获取go的最新下载路径（linux-amd64)
func GetGoLink() string {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var ans string
	c.OnHTML("tr > td > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		text := e.Text
		if ans == "" && strings.Contains(text,"linux-amd64") {
			ans = link
		}
	})
	c.Visit("https://golang.org/dl/")
	return ans
}