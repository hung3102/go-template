package entities

// ParallelExecution - 並列実行
type ParallelExecution struct {
	id            string // id
	eventDocID    string // event_doc_id
	executionType int    // 実行種別
	identifierID  string // account_id or org_csp_id
	meta          *Meta  // メタ
}

// NewParallelExecutionParam - 並列実行作成パラメータ
type NewParallelExecutionParam struct {
	ID            string // id
	EventDocID    string // event_doc_id
	ExecutionType int    // 実行種別
	IdentifierID  string // account_id or org_csp_id
	Meta          *Meta  // メタ
}

// ParallelExecution - 並列実行結果作成
func NewParallelExecution(param *NewParallelExecutionParam) *ParallelExecution {
	return &ParallelExecution{
		id:            param.ID,
		eventDocID:    param.EventDocID,
		executionType: param.ExecutionType,
		identifierID:  param.IdentifierID,
		meta:          param.Meta,
	}
}

// ID - ID のゲッター
func (e *ParallelExecution) ID() string {
	return e.id
}

// EventDocID - EventDocID のゲッター
func (e *ParallelExecution) EventDocID() string {
	return e.eventDocID
}

// ExecutionType - ExecutionType のゲッター
func (e *ParallelExecution) ExecutionType() int {
	return e.executionType
}

// IdentifierId - IdentifierId のゲッター
func (e *ParallelExecution) IdentifierId() string {
	return e.identifierID
}

// Meta - meta のゲッター
func (e *ParallelExecution) Meta() *Meta {
	return e.meta
}
