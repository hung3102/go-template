package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	id            valueobjects.ParallelExecutionResultID // id
	eventID       valueobjects.EventID                   // event_id
	executionType int                                    // 費用按分計算請求データ
	resultCode    int                                    // 実行結果 成功 or 失敗
	errorMessage  *ParallelExecutionResultErrorMessage   // error object
	meta          *Meta                                  // メタ
}

// NewParallelExecutionResultParam - 並列実行結果作成パラメータ
type NewParallelExecutionResultParam struct {
	ID            valueobjects.ParallelExecutionResultID // id
	EventID       valueobjects.EventID                   // event_id
	ExecutionType int                                    // 費用按分計算請求データ
	ResultCode    int                                    // 実行結果 成功 or 失敗
	ErrorMessage  *ParallelExecutionResultErrorMessage   // error object
	Meta          *Meta                                  // メタ
}

// NewParallelExecutionResult - 並列実行結果作成
func NewParallelExecutionResult(param *NewParallelExecutionResultParam) *ParallelExecutionResult {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewParallelExecutionResultID()
	}
	return &ParallelExecutionResult{
		id:            id,
		eventID:       param.EventID,
		executionType: param.ExecutionType,
		resultCode:    param.ResultCode,
		errorMessage:  param.ErrorMessage,
		meta:          param.Meta,
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

// ExecutionType - ExecutionType のゲッター
func (e *ParallelExecutionResult) ExecutionType() int {
	return e.executionType
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
