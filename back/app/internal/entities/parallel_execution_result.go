package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	id           valueobjects.ParallelExecutionResultID   // ID
	eventID      valueobjects.EventID                     // イベントID
	kind         valueobjects.ParallelExecutionResultKind // 処理種別
	resultCode   int                                      // 実行結果
	errorMessage *ParallelExecutionResultErrorMessage     // エラー情報
	meta         *Meta                                    // メタ
}

// NewParallelExecutionResultParam - 並列実行結果作成パラメータ
type NewParallelExecutionResultParam struct {
	ID           valueobjects.ParallelExecutionResultID   // ID
	EventID      valueobjects.EventID                     // イベントID
	Kind         valueobjects.ParallelExecutionResultKind // 処理種別
	ResultCode   int                                      // 実行結果
	ErrorMessage *ParallelExecutionResultErrorMessage     // エラー情報
	Meta         *Meta                                    // メタ
}

// NewParallelExecutionResult - 並列実行結果作成
func NewParallelExecutionResult(param *NewParallelExecutionResultParam) *ParallelExecutionResult {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewParallelExecutionResultID()
	}
	return &ParallelExecutionResult{
		id:           id,
		eventID:      param.EventID,
		kind:         param.Kind,
		resultCode:   param.ResultCode,
		errorMessage: param.ErrorMessage,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *ParallelExecutionResult) ID() valueobjects.ParallelExecutionResultID {
	return e.id
}

// EventID - EventID のゲッター
func (e *ParallelExecutionResult) EventID() valueobjects.EventID {
	return e.eventID
}

// Kind - Kind のゲッター
func (e *ParallelExecutionResult) Kind() valueobjects.ParallelExecutionResultKind {
	return e.kind
}

// ResultCode - ResultCode のゲッター
func (e *ParallelExecutionResult) ResultCode() int {
	return e.resultCode
}

// ErrorMessage - ErrorMessage のゲッター
func (e *ParallelExecutionResult) ErrorMessage() *ParallelExecutionResultErrorMessage {
	return e.errorMessage
}

// Meta - Meta のゲッター
func (e *ParallelExecutionResult) Meta() *Meta {
	return e.meta
}
