package models

type EntryStatus struct {
	Code int64
	Name string
}

func NewEntryStatus(code int64) EntryStatus {
	name := []string{"accepted", "processing", "registered", "error", "disabled"}[code]
	return EntryStatus{Code: code, Name: name}
}

var (
	EntryStatusAccepted   = NewEntryStatus(0)
	EntryStatusProcessing = NewEntryStatus(1)
	EntryStatusRegistered = NewEntryStatus(2)
	EntryStatusError      = NewEntryStatus(3)
	EntryStatusDisabled   = NewEntryStatus(4)
)
