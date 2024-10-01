package valueobjects

import "golang.org/x/xerrors"

// ParallelExecutionResultKind - ParallelExecutionResultコレクションのKindの値
// https://www.notion.so/Status-Type-fcd964996da74dcab47dabdc6b35434e#71100b4022104541aeffbdbcc02f7127
type ParallelExecutionResultKind interface {
	// String - ParallelExecutionResultKindを文字列に変換する
	String() string
}

var (
	// 他環境参照しない費用按分処理
	ParallelExecutionResultKindIndependentProportion ParallelExecutionResultKind = parallelExecutionResultKind{value: "independent_proportion"}
	// 他環境参照の費用按分処理
	ParallelExecutionResultKindDependentProportion ParallelExecutionResultKind = parallelExecutionResultKind{value: "dependent_proportion"}
	// 請求データーを作成する
	ParallelExecutionResultKindBillDataCreate ParallelExecutionResultKind = parallelExecutionResultKind{value: "bill_data_create"}
	// 収納番号を取得する
	ParallelExecutionResultKindPaymentNum ParallelExecutionResultKind = parallelExecutionResultKind{value: "payment_num"}
	// 請求書のPDFファイルを作成する
	ParallelExecutionResultKindPdfCreate ParallelExecutionResultKind = parallelExecutionResultKind{value: "pdf_create"}
	// メール送信する
	ParallelExecutionResultKindSendEmail ParallelExecutionResultKind = parallelExecutionResultKind{value: "send_email"}

	// ParallelExecutionResultKindのマップ
	parallelExecutionResultKindMap = make(map[string]ParallelExecutionResultKind)
)

func init() {
	parallelExecutionResultKindMap["independent_proportion"] = ParallelExecutionResultKindIndependentProportion
	parallelExecutionResultKindMap["dependent_proportion"] = ParallelExecutionResultKindDependentProportion
	parallelExecutionResultKindMap["bill_data_create"] = ParallelExecutionResultKindBillDataCreate
	parallelExecutionResultKindMap["payment_num"] = ParallelExecutionResultKindPaymentNum
	parallelExecutionResultKindMap["pdf_create"] = ParallelExecutionResultKindPdfCreate
	parallelExecutionResultKindMap["send_email"] = ParallelExecutionResultKindSendEmail
}

// parallelExecutionResultKind - ParallelExecutionResultKindの実装
type parallelExecutionResultKind struct {
	value string // 文字列の値
}

// NewParallelExecutionResultKind - NewParallelExecutionResultKindを作成する
func NewParallelExecutionResultKind(str string) (ParallelExecutionResultKind, error) {
	parallelExecutionResultKind, ok := parallelExecutionResultKindMap[str]
	if !ok {
		return nil, xerrors.Errorf("err in NewParallelExecutionResultKindFromString: str = %s", str)
	}
	return parallelExecutionResultKind, nil
}

// String - ParallelExecutionResultKindを文字列に変換する
func (v parallelExecutionResultKind) String() string {
	return v.value
}
