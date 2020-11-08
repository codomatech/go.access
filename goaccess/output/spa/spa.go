package goaccess_analyzers

import (
	"encoding/json"
	"github.com/codomatech/go.access/common"
	"io/ioutil"
	"log"
	"os"
)

func Output(results []common.AnalysisResult, opath string) {
	var err error
	basedir := opath + "/spa"
	jsondata, _ := json.MarshalIndent(results, "", "  ")
	err = os.Mkdir(basedir, 0755)
	if err != nil {
		log.Fatalf("Failed to create output director %s (%+v)", basedir, err)
	}

	err = ioutil.WriteFile(basedir+"/data.json", jsondata, 0644)
}
