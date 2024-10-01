package valueobjects_test

import (
	"testing"

	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
)

func TestHoge(t *testing.T) {
	p := valueobjects.ParallelExecutionResultKindIndependentProportion
	s := p.String()
	if s != "independent_proportion" {
		t.Fatalf("err: s = %s", s)
	}

	p, err := valueobjects.NewParallelExecutionResultKind(s)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	s = p.String()
	if s != "independent_proportion" {
		t.Fatalf("err: s = %s", s)
	}

	p, err = valueobjects.NewParallelExecutionResultKind("a")
	if err == nil {
		t.Fatalf("err: err == nil")
	}
}
