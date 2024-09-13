package invoicepdf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/signintech/gopdf"
)

const fontName = "fontName"

//go:embed invoice.template.pdf
var template []byte

//go:embed VL-PGothic-Regular.ttf
var fontData []byte

type CreateInvoicePDFParam struct {
	OrgName string //団体名
}

func CreateInvoicePDF(param *CreateInvoicePDFParam) ([]byte, error) {
	invoicePDF := createInvoicePDF{
		pdf:   &gopdf.GoPdf{},
		param: param,
	}
	return invoicePDF.execute()
}

type createInvoicePDF struct {
	pdf   *gopdf.GoPdf
	param *CreateInvoicePDFParam
}

func (ip *createInvoicePDF) execute() ([]byte, error) {
	ip.initPage()
	ip.setTemplate()
	if err := ip.loadFont(); err != nil {
		return nil, err
	}
	if err := ip.setOrgName(); err != nil {
		return nil, err
	}
	return ip.getPDF(), nil
}

func (ip *createInvoicePDF) initPage() {
	ip.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	ip.pdf.AddPage()
}

func (ip *createInvoicePDF) setTemplate() {
	var templateReader io.ReadSeeker = bytes.NewReader(template)
	tpl := ip.pdf.ImportPageStream(&templateReader, 1, "/MediaBox")
	ip.pdf.UseImportedTemplate(tpl, 0, 0, 595, 842)
}

func (ip *createInvoicePDF) loadFont() error {
	if err := ip.pdf.AddTTFFontData(fontName, fontData); err != nil {
		return fmt.Errorf("InvoicePDFImpl.loadFont: ip.pdf.AddTTFFontData: %v", err)
	}
	return nil
}

func (ip *createInvoicePDF) setOrgName() error {
	if err := ip.pdf.SetFont(fontName, "", 14); err != nil {
		return fmt.Errorf("InvoicePDFImpl.setOrgName: ip.pdf.SetFont: %v", err)
	}
	ip.pdf.SetXY(35, 57)
	ip.pdf.SetTextColor(255, 0, 0)
	ip.pdf.Cell(nil, ip.param.OrgName)
	return nil
}

func (ip *createInvoicePDF) getPDF() []byte {
	return ip.pdf.GetBytesPdf()
}
