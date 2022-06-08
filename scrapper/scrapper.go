package scrapper

import (
	"fmt"
	"os"
	"sync"

	"github.com/gocolly/colly"
	"github.com/msam1r/natiga22/result"
)

type Scrapper struct {
	From int
	To   int
	File *os.File
}

var wg sync.WaitGroup

func (sc *Scrapper) Start() {
	gc := colly.NewCollector(colly.AllowedDomains("www.controlbm.com"))
	pc := colly.NewCollector(colly.AllowedDomains("www.controlbm.com"))

	for i := sc.From; i <= sc.To; i++ {
		wg.Add(1)
		fmt.Println(i)
		sc.scrapeStudentResult(i, gc, pc)
	}

	wg.Wait()
}

func (sc *Scrapper) scrapeStudentResult(number int, gc, pc *colly.Collector) {
	gc.OnHTML("form", func(h *colly.HTMLElement) {
		data := collectFormData(number, h)

		pc.OnHTML("form", func(h *colly.HTMLElement) {
			defer wg.Done()
			s := collectStudentResult(h)
			s.ToCSV(sc.File)
		})

		pc.Post("https://www.controlbm.com/alnatiga12345", data)
	})

	gc.Visit("https://www.controlbm.com/alnatiga12345")
}

func collectFormData(number int, h *colly.HTMLElement) map[string]string {
	views_state := h.ChildAttr("input[name=__VIEWSTATE]", "value")
	views_state_generator := h.ChildAttr("input[name=__VIEWSTATEGENERATOR]", "value")
	event_validation := h.ChildAttr("input[name=__EVENTVALIDATION]", "value")

	body := map[string]string{
		"__VIEWSTATE":          views_state,
		"__VIEWSTATEGENERATOR": views_state_generator,
		"__EVENTVALIDATION":    event_validation,
		"SearchButton":         "بحث",
		"GeloseNumberTextBox":  fmt.Sprint(number),
	}

	return body
}

func collectStudentResult(h *colly.HTMLElement) *result.Student {
	student := result.NewStudent(
		h.ChildText("#GeloseNumberLabel"),
		h.ChildText("#StudentName"),
		h.ChildText("#SchoolNameLabel"),
		h.ChildText("#StudentTypeLabel"),
	)

	result := &result.Result{
		Arabic:     h.ChildText("#ArabicLabel"),
		English:    h.ChildText("#EnglishLabel"),
		Studies:    h.ChildText("#DrasatLabel"),
		Algebra:    h.ChildText("#GabrLabel"),
		Geometry:   h.ChildText("#HandasaLabel"),
		Total_math: h.ChildText("#MagReadyatLabel"),
		Science:    h.ChildText("#ScienceLabel"),
		Total:      h.ChildText("#MagKollyLabel"),
		Religion:   h.ChildText("#DeanLabel"),
		Art:        h.ChildText("#FanyaLabel"),
		Computer:   h.ChildText("#HasebLabel"),
		Sport:      h.ChildText("#TarbyaReadyaLabel"),
	}

	student.AttachResult(result)

	return student
}
