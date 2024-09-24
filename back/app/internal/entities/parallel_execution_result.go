package entities

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	id            string                               // id
	eventID       string                               // event_id
	executionType int                                  // 費用按分計算請求データ
	resultCode    int                                  // 実行結果 成功 or 失敗
	errorMessage  *ParallelExecutionResultErrorMessage // error object
	meta          *Meta                                // メタ
}

// NewParallelExecutionResultParam - 並列実行結果作成パラメータ
type NewParallelExecutionResultParam struct {
	ID            string                               // id
	EventID       string                               // event_id
	ExecutionType int                                  // 費用按分計算請求データ
	ResultCode    int                                  // 実行結果 成功 or 失敗
	ErrorMessage  *ParallelExecutionResultErrorMessage // error object
	Meta          *Meta                                // メタ
}

// NewParallelExecutionResult - 並列実行結果作成
func NewParallelExecutionResult(param *NewParallelExecutionResultParam) *ParallelExecutionResult {
	return &ParallelExecutionResult{
		id:            param.ID,
		eventID:       param.EventID,
		executionType: param.ExecutionType,
		resultCode:    param.ResultCode,
		errorMessage:  param.ErrorMessage,
		meta:          param.Meta,
	}
}

// ID - ID のゲッター
func (e *ParallelExecutionResult) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *ParallelExecutionResult) EventID() string {
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
