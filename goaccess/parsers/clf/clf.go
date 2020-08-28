package goaccess_parsers

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/codomatech/goaccess.go/common"
)

func Parse(line string) common.AccessRecord {
	var record common.AccessRecord
	var request common.Request
	re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) .*?\[(.*?)\] "(.*?)" (\d+) (\d+) "(.*?)" "(.*?)"`)

	result := re.FindStringSubmatch(line)

	record.Ip = result[1]

	t, _ := time.Parse("2/Jan/2006:15:04:05 -0700", result[2])
	record.Timestamp = t.Unix()

	req := strings.Split(result[3], " ")
	request.Method = req[0]
	request.Path = req[1]
	request.Protocol = req[2]
	record.Request = request

	record.Status, _ = strconv.Atoi(result[4])
	record.Size, _ = strconv.ParseInt(result[5], 0, 64)
	record.Referrer = result[6]
	record.UserAgent = result[7]

	return record
}
