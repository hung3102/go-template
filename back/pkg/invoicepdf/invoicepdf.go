// Package invoicepdf - 請求書PDFを作成するためのパッケージ
package invoicepdf

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/signintech/gopdf"
	"golang.org/x/xerrors"
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
	result, err := invoicePDF.execute()
	if err != nil {
		return nil, xerrors.Errorf("error in invoicePDF.execute method: %w", err)
	}
	return result, nil
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
		return nil, xerrors.Errorf("createInvoicePDF.execute: %w", err)
	}
	if err := ip.setOrgName(); err != nil {
		return nil, xerrors.Errorf("createInvoicePDF.execute: %w", err)
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
		return xerrors.Errorf("createInvoicePDF.loadFont: %w", err)
	}
	return nil
}

// setOrgName - PDFに団体名を設定する
func (ip *createInvoicePDF) setOrgName() error {
	if err := ip.pdf.SetFont(fontName, "", 14); err != nil {
		return xerrors.Errorf("createInvoicePDF.setOrgName: %w", err)
	}
	ip.pdf.SetXY(35, 57)
	ip.pdf.SetTextColor(255, 0, 0)
	if err := ip.pdf.Cell(nil, ip.param.OrgName); err != nil {
		return xerrors.Errorf("createInvoicePDF.setOrgName: %w", err)
	}
	return nil
}

// getPDF - PDFのbyte配列を取得する
func (ip *createInvoicePDF) getPDF() []byte {
	return ip.pdf.GetBytesPdf()
}
