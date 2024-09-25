package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c parallel_execution_result -mockgen ../../../../../bin/mockgen -mock-output mocks/parallel_execution_result_gen.go ParallelExecutionResult

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	ID            string                               `firestore:"-" firestore_key:""` // id
	EventID       string                               `firestore:"event_id"`           // event_id
	ExecutionType int                                  `firestore:"execution_type"`     // 費用按分計算請求データ
	ResultCode    int                                  `firestore:"result_code"`        // 実行結果 成功 or 失敗
	ErrorMessage  *ParallelExecutionResultErrorMessage `firestore:"error_message"`      // error object
	Meta
}
