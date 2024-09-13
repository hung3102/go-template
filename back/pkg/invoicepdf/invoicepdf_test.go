package invoicepdf

import (
	"fmt"
	"testing"
)

func Test_CreateInvoicePDF(t *testing.T) {
	pdf, err := CreateInvoicePDF(&CreateInvoicePDFParam{
		OrgName: "団体名12345団体名12345団体名12345",
	})

	if err != nil {
		t.Errorf("CreateInvoicePDF err = %v", err)
	}

	fmt.Println(pdf)

	// err = os.WriteFile("invoice_pdf_test.pdf", pdf, 0644)
	// if err != nil {
	// 	t.Errorf("os.WriteFile err = %v", err)
	// }
}
