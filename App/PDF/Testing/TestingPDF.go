package Testing

import (
	"Service/App/PDF"
	"Service/App/Service/Constant/Path"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)

type TestingPDF struct {
	Name string
}

func (pdf TestingPDF) Generate() string {
	xpdf := PDF.PDF{}
	xpdf.NewGenerator("PDFTemplate.html", pdf)

	xpdf.PDFG.Dpi.Set(300)
	xpdf.PDFG.PageSize.Set("A4")
	xpdf.PDFG.Orientation.Set(wkhtml.OrientationPortrait)

	path := Path.PathPDFTesting()
	filename := "hasil.pdf"

	err := xpdf.Save(path, filename)
	if err != nil {
		log.Panicf("Unable to generate PDF: %s", err)
	}

	return path + filename
}
