package getSoftLink
import (
	"github.com/gocolly/colly"
	"strings"
	// "fmt"
)
// 获取redis的最新下载路径
func GetRedisLink() string {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var ans string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if ans == "" && strings.Contains(link,"download.redis.io/releases") {
			ans = link
		}
	})
	c.Visit("https://redis.io/")
	return ans
}