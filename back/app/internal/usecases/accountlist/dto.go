package accountlist

// AccountListInput - アカウントリストのinput
type AccountListInput struct {
	EventDocID string
}

// AccountListOutput - アカウントリストのoutput
type AccountListOutput struct {
	GCASAccountDocIDs []string
}
