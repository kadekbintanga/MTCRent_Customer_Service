package core

import (
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/globalxtreme/gobaseconf/helpers"
	"log"
)

type PDF struct {
	PDFG *wkhtml.PDFGenerator
}

func (x *PDF) NewGenerator(layout string, data interface{}) {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Panicf("New Generator invalid: %s", err)
	}

	x.PDFG = pdfg

	buffer := helpers.PDFHTMLTemplate(layout, data)
	page := wkhtml.NewPageReader(&buffer)
	page.DisableExternalLinks.Set(true)

	page.HeaderHTML.Set("internal/pkg/layout/components/header.html")
	page.FooterHTML.Set("internal/pkg/layout/components/footer.html")

	x.PDFG.AddPage(page)

	x.PDFG.MarginTop.Set(25)
	x.PDFG.MarginBottom.Set(29)
	x.PDFG.MarginLeft.Set(0)
	x.PDFG.MarginRight.Set(0)
}

func (x *PDF) Save(path string, filename string) error {
	path = helpers.SetStorageAppDir(path)
	helpers.CheckAndCreateDirectory(path)

	err := x.PDFG.Create()
	if err != nil {
		return err
	}

	err = x.PDFG.WriteFile(path + filename)
	if err != nil {
		return err
	}

	return nil
}
