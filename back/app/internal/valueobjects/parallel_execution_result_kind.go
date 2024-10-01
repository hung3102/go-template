package valueobjects

import "golang.org/x/xerrors"

// ParallelExecutionResultKind - ParallelExecutionResultコレクションのKindの値
// https://www.notion.so/Status-Type-fcd964996da74dcab47dabdc6b35434e#71100b4022104541aeffbdbcc02f7127
type ParallelExecutionResultKind string

const (
	// 他環境参照しない費用按分処理
	ParallelExecutionResultKindIndependentProportion = ParallelExecutionResultKind("independent_proportion")
	// 他環境参照の費用按分処理
	ParallelExecutionResultKindDependentProportion = ParallelExecutionResultKind("dependent_proportion")
	// 請求データーを作成する
	ParallelExecutionResultKindBillDataCreate = ParallelExecutionResultKind("bill_data_create")
	// 収納番号を取得する
	ParallelExecutionResultKindPaymentNum = ParallelExecutionResultKind("payment_num")
	// 請求書のPDFファイルを作成する
	ParallelExecutionResultKindPdfCreate = ParallelExecutionResultKind("pdf_create")
	// メール送信する
	ParallelExecutionResultKindSendEmail = ParallelExecutionResultKind("send_email")
)

var (
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

// NewParallelExecutionResultKindFromString - 文字列をParallelExecutionResultKindに変換する
func NewParallelExecutionResultKindFromString(str string) (ParallelExecutionResultKind, error) {
	parallelExecutionResultKind, ok := parallelExecutionResultKindMap[str]
	if !ok {
		return ParallelExecutionResultKind(""), xerrors.Errorf("err in NewParallelExecutionResultKindFromString: str = %s", str)
	}
	return parallelExecutionResultKind, nil
}

// String - ParallelExecutionResultKindを文字列に変換する
func (v ParallelExecutionResultKind) String() string {
	return string(v)
}
