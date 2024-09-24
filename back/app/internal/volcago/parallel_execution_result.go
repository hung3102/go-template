package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/parallel_execution_result_gen.go ParallelExecutionResult

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	ID            string                               `firestore:"-" firestore_key:""` // id
	EventID       string                               ``                               // event_id
	ExecutionType int                                  ``                               // 費用按分計算請求データ
	ResultCode    int                                  ``                               // 実行結果 成功 or 失敗
	ErrorMessage  *ParallelExecutionResultErrorMessage ``                               // error object
	Meta
}
