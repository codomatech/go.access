package common

type AccessRecord struct {
	Timestamp int64
	Status    int
	Size      int64
	Referrer  string
	UserAgent string
	Ip        string
	Request   Request
}

type Request struct {
	Path     string
	Method   string
	Protocol string
}

type AnalysisResult struct {
	Name     string
}
