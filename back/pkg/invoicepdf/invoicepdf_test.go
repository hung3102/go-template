package invoicepdf_test

import (
	"fmt"
	"testing"

	"github.com/topgate/gcim-temporary/back/pkg/invoicepdf"
)

func Test_CreateInvoicePDF(t *testing.T) {
	pdf, err := invoicepdf.CreateInvoicePDF(&invoicepdf.CreateInvoicePDFParam{
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
