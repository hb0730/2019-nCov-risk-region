package source

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type BenDiBao struct {
	Headless bool
	browser  *rod.Browser
}

func NewBenDiBao(headless bool) *BenDiBao {
	bendibao := new(BenDiBao)
	bendibao.Headless = headless
	return bendibao
}
func (b *BenDiBao) Time() string {
	if b.browser == nil {
		b.openBrowser()
	}
	t := b.getPage().MustElement(`p.time`).MustText()
	return t
}

func (b *BenDiBao) HighRisk() []Risk {
	if b.browser == nil {
		b.openBrowser()
	}
	page := b.getPage()
	defer page.MustClose()
	elements := page.MustElement(`.height`).MustElements(`.info-list`)
	if elements.Empty() {
		return nil
	}
	return b.getRisk(elements)
}
func (b *BenDiBao) MiddleRisk() []Risk {
	if b.browser == nil {
		b.openBrowser()
	}
	page := b.getPage()
	defer page.MustClose()
	elements := page.MustElement(`.middle`).MustElements(`.info-list`)
	if elements.Empty() {
		return nil
	}
	return b.getRisk(elements)
}
func (b *BenDiBao) openBrowser() {
	if b.browser == nil {
		la := launcher.New().Headless(b.Headless).MustLaunch()
		b.browser = rod.New().ControlURL(la).MustConnect()
	}
}

func (b *BenDiBao) Close() (err error) {
	if b.browser != nil {
		err = b.browser.Close()
	}
	b.browser = nil
	return
}

func (b *BenDiBao) getRisk(elements rod.Elements) []Risk {
	risk := []Risk{}
	for _, v := range elements {
		city := v.MustElements(`div.shi>span`)
		if city.Empty() {
			continue
		}
		province := city.First()
		shi := city.Last()
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
			Province:   province.MustText(),
			City:       shi.MustText(),
			Communitys: communitys,
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
