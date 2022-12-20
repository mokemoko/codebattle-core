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

type MatchStatus struct {
	Code int64
	Name string
}

func NewMatchStatus(code int64) MatchStatus {
	name := []string{"requested", "ongoing", "finished", "error"}[code]
	return MatchStatus{Code: code, Name: name}
}

var (
	MatchStatusRequested = NewMatchStatus(0)
	MatchStatusOngoing   = NewMatchStatus(1)
	MatchStatusFinished  = NewMatchStatus(2)
	MatchStatusError     = NewMatchStatus(3)
)

type MatchType struct {
	Code int64
	Name string
}

func NewMatchType(code int64) MatchType {
	name := []string{"rated", "unrated"}[code]
	return MatchType{Code: code, Name: name}
}

var (
	MatchTypeRated   = NewMatchType(0)
	MatchTypeUnrated = NewMatchType(1)
)
