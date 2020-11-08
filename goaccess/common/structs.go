package common

import "strconv"

type AccessRecord struct {
	Day       string
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
	Name string
	X    []string
	Ys   map[string][]string
}

func (r *AnalysisResult) Init(capacity int, ynames ...string) {
	// allocate non-atomic members
	r.X = make([]string, 0, capacity)
	r.Ys = make(map[string][]string)
	for _, name := range ynames {
		r.Ys[name] = make([]string, 0, capacity)
	}
}

func (r *AnalysisResult) AddX(v string) {
	r.X = append(r.X, v)
}

func (r *AnalysisResult) AddY(name string, v float64) {
	r.Ys[name] = append(r.Ys[name], strconv.FormatFloat(v, 'f', 4, 64))
}
