package goaccess_parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := `123.45.678.90 - - [14/Aug/2020:06:33:07 +0200] "GET /page1.html HTTP/2.0" 200 3784 "https://googleads.g.doubleclick.net/pagead/123456" "Mozilla/5.0 (Linux; Android 8.1.0; CPH1803) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.111 Mobile Safari/537.36"`

	accessRecord := Parse(line)
	assert.Equal(t, accessRecord.Timestamp, int64(1597379587))
	assert.Equal(t, accessRecord.Status, 200)
	assert.Equal(t, accessRecord.Size, int64(3784))
	assert.Equal(t, accessRecord.Ip, "123.45.678.90")
	assert.Equal(t, accessRecord.Referrer, "https://googleads.g.doubleclick.net/pagead/123456")
	assert.Equal(t, accessRecord.UserAgent, "Mozilla/5.0 (Linux; Android 8.1.0; CPH1803) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.111 Mobile Safari/537.36")
	assert.Equal(t, accessRecord.Request.Path, "/page1.html")
	assert.Equal(t, accessRecord.Request.Protocol, "HTTP/2.0")
	assert.Equal(t, accessRecord.Request.Method, "GET")
}
