package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c parallel_execution_result -mockgen ../../../../../bin/mockgen -mock-output mocks/parallel_execution_result_gen.go ParallelExecutionResult

// ParallelExecutionResult - 並列実行結果
type ParallelExecutionResult struct {
	ID           string                               `firestore:"-" firestore_key:""` // ID
	EventID      string                               `firestore:"event_id"`           // イベントID
	Kind         string                               `firestore:"kind"`               // 処理種別
	ResultCode   int                                  `firestore:"result_code"`        // 実行結果
	ErrorMessage *ParallelExecutionResultErrorMessage `firestore:"error_message"`      // エラー情報
	Meta
}
