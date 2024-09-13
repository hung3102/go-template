package invoicepdf

import (
	"os"
	"testing"
)

func Test_InvoicePDF_WritePDF(t *testing.T) {
	pdf, err := CreateInvoicePDF(&CreateInvoicePDFParam{
		OrgName: "団体名12345団体名12345団体名12345",
	})

	if err != nil {
		t.Errorf("CreateInvoicePDF err = %v", err)
	}

	err = os.WriteFile("invoice_pdf_test.pdf", pdf, 0644)
	if err != nil {
		t.Errorf("os.WriteFile err = %v", err)
	}
}
