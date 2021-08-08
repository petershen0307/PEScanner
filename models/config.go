package models

type scanMode int

const (
	Single scanMode = iota + 1
	Concurrent
)

const MaxConcurrentNumber = 5
const MinConcurrentNumber = 1

type Config struct {
	Mode             scanMode
	ConcurrentNumber int
	EntryFolder      string
	OutputDir        string
}
