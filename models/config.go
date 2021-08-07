package models

type scanMode int

const Single scanMode = 1
const Concurrent scanMode = 2

const MaxConcurrentNumber = 5
const MinConcurrentNumber = 1

type Config struct {
	Mode             scanMode
	ConcurrentNumber int
	EntryFolder      string
	OutputDir        string
}
