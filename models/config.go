package models

type ScanMode int

func (mode *ScanMode) String() string {
	switch *mode {
	case Single:
		return "single"
	case Concurrent:
		return "concurrent"
	}
	return "None"
}

const (
	Single ScanMode = iota + 1
	Concurrent
)

const MaxConcurrentNumber = 5
const MinConcurrentNumber = 1

type Config struct {
	Mode             ScanMode
	ConcurrentNumber int
	EntryFolder      string
	OutputDir        string
}
