package accountlist

// AccoutnListInput - アカウントリストのinput
type AccoutnListInput struct {
	EventDocID string
}

// AccountListOutput - アカウントリストのoutput
type AccountListOutput struct {
	GCASAccountDocIDs []string
}
