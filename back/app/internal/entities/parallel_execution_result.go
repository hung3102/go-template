package entities

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	id            string                               // id
	eventDocID    string                               // event_doc_id
	executionType int                                  // 実行種別
	identifierID  string                               // account_id or org_csp_id
	resultCode    int                                  // 実行結果 成功 or 失敗
	errorMessage  *ParallelExecutionResultErrorMessage // error object
	meta          *Meta                                // メタ
}

// NewParallelExecutionResultParam - 並列実行結果作成パラメータ
type NewParallelExecutionResultParam struct {
	ID            string                               // id
	EventDocID    string                               // event_doc_id
	ExecutionType int                                  // 実行種別
	IdentifierID  string                               // account_id or org_csp_id
	ResultCode    int                                  // 実行結果 成功 or 失敗
	ErrorMessage  *ParallelExecutionResultErrorMessage // error object
	Meta          *Meta                                // メタ
}

// NewParallelExecutionResult - 並列実行結果作成
func NewParallelExecutionResult(param *NewParallelExecutionResultParam) *ParallelExecutionResult {
	return &ParallelExecutionResult{
		id:            param.ID,
		eventDocID:    param.EventDocID,
		executionType: param.ExecutionType,
		identifierID:  param.IdentifierID,
		resultCode:    param.ResultCode,
		errorMessage:  param.ErrorMessage,
		meta:          param.Meta,
	}
}

// ID - ID のゲッター
func (e *ParallelExecutionResult) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *ParallelExecutionResult) EventDocID() string {
	return e.eventDocID
}

// ExecutionType - ExecutionType のゲッター
func (e *ParallelExecutionResult) ExecutionType() int {
	return e.executionType
}

// IdentifierID - IdentifierID のゲッター
func (e *ParallelExecutionResult) IdentifierID() string {
	return e.identifierID
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
