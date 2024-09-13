// Package invoicepdf - 請求書PDFを作成するためのパッケージ
package invoicepdf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/signintech/gopdf"
)

// fontName - フォント名
const fontName = "fontName"

// template - 請求書PDFのテンプレートデータ
//
//go:embed invoice.template.pdf
var template []byte

// fontData - フォントデータ
//
//go:embed VL-PGothic-Regular.ttf
var fontData []byte

// CreateInvoicePDFParam - 請求書PDF作成のためのパラメーター
type CreateInvoicePDFParam struct {
	OrgName string //団体名
}

// CreateInvoicePDF - 請求書PDFを作成する
func CreateInvoicePDF(param *CreateInvoicePDFParam) ([]byte, error) {
	invoicePDF := createInvoicePDF{
		pdf:   &gopdf.GoPdf{},
		param: param,
	}
	return invoicePDF.execute()
}

// createInvoicePDF - 請求書PDFを作成
type createInvoicePDF struct {
	pdf   *gopdf.GoPdf
	param *CreateInvoicePDFParam
}

// execute - 請求書PDFの作成処理本体
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

// initPage - ページを作成する
func (ip *createInvoicePDF) initPage() {
	ip.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	ip.pdf.AddPage()
}

// setTemplate - PDFのテンプレートを設定する
func (ip *createInvoicePDF) setTemplate() {
	var templateReader io.ReadSeeker = bytes.NewReader(template)
	tpl := ip.pdf.ImportPageStream(&templateReader, 1, "/MediaBox")
	ip.pdf.UseImportedTemplate(tpl, 0, 0, 595, 842)
}

// loadFont - フォントファイルを読み込む
func (ip *createInvoicePDF) loadFont() error {
	if err := ip.pdf.AddTTFFontData(fontName, fontData); err != nil {
		return fmt.Errorf("InvoicePDFImpl.loadFont: ip.pdf.AddTTFFontData: %v", err)
	}
	return nil
}

// setOrgName - PDFに団体名を設定する
func (ip *createInvoicePDF) setOrgName() error {
	if err := ip.pdf.SetFont(fontName, "", 14); err != nil {
		return fmt.Errorf("InvoicePDFImpl.setOrgName: ip.pdf.SetFont: %v", err)
	}
	ip.pdf.SetXY(35, 57)
	ip.pdf.SetTextColor(255, 0, 0)
	ip.pdf.Cell(nil, ip.param.OrgName)
	return nil
}

// getPDF - PDFのbyte配列を取得する
func (ip *createInvoicePDF) getPDF() []byte {
	return ip.pdf.GetBytesPdf()
}
