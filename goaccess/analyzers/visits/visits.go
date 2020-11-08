/**
 * Analysis for visits and hits
 */

package goaccess_analyzers

import (
	"github.com/codomatech/go.access/common"
)

func Analyze(records []common.AccessRecord) common.AnalysisResult {
	var result common.AnalysisResult

	result.Init(len(records)/512, "visits", "hits")
	result.Name = "Visits & Hits"

	var dayvisitors map[string]bool
	var dayhits uint
	var curday = ""

	for _, record := range records {
		if record.Day != curday {
			if len(curday) > 0 {
				result.AddX(curday)
				result.AddY("hits", float64(dayhits))
				result.AddY("visits", float64(len(dayvisitors)))
			}
			curday = record.Day
			dayhits = 0
			dayvisitors = make(map[string]bool)
		}
		dayhits += 1
		dayvisitors[record.Ip] = true
	}

	return result
}
