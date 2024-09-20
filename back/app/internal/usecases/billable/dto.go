package billable

// Input - 請求書作成の開始判定のinput
type Input struct {
	EventDocID string
}

// Output - 請求書作成の開始判定のoutput
type Output struct {
	GCASAccountDocIDs []string
}
