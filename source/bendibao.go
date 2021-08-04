package source

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type BenDiBao struct {
	browser *rod.Browser
}

func NewBenDiBao(headless bool) *BenDiBao {
	bendibao := new(BenDiBao)
	la := launcher.New().Headless(headless).MustLaunch()
	bendibao.browser = rod.New().ControlURL(la).MustConnect()
	return bendibao
}
func (b *BenDiBao) Time() string {
	t := b.getPage().MustElement(`p.time`).MustText()
	return t
}

func (b *BenDiBao) HighRisk() []Risk {
	page := b.getPage()
	defer page.MustClose()
	elements := page.MustElement(`.height`).MustElements(`.info-list`)
	if elements.Empty() {
		return nil
	}
	return b.getRisk(elements)
}
func (b *BenDiBao) MiddleRisk() []Risk {
	page := b.getPage()
	defer page.MustClose()
	elements := page.MustElement(`.middle`).MustElements(`.info-list`)
	if elements.Empty() {
		return nil
	}
	return b.getRisk(elements)
}

func (b BenDiBao) Close() error {
	return b.browser.Close()
}

func (b *BenDiBao) getRisk(elements rod.Elements) []Risk {
	risk := []Risk{}
	for _, v := range elements {
		province := v.MustElement(`div.province>span`).MustText()
		infoElements := v.MustElements(`ul.info-detail>li`)
		if infoElements.Empty() {
			continue
		}
		communitys := []string{}
		for _, v := range infoElements {
			info := v.MustElement(`span`).MustText()
			communitys = append(communitys, info)
		}
		r := Risk{
			Province:   province,
			communitys: communitys,
		}
		risk = append(risk, r)
	}
	return risk
}

func (b *BenDiBao) getPage() *rod.Page {
	return b.browser.
		MustPage("http://m.sh.bendibao.com/news/gelizhengce/fengxianmingdan.php").
		MustWaitLoad()
}

func init() {
	Instance.Put("bendibao", NewBenDiBao(false))
}
