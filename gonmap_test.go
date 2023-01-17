package gonmap

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	var scanner = New()
	host := "172.16.130.138"
	port := 9200
	status, response := scanner.ScanTimeout(host, port, time.Second*30)
	println(status, response.FingerPrint.Service, host, ":", port)

}

func TestPattern(t *testing.T) {
	pattern := "[a-zA-Z0-9\u2460-\u24FF]"
	check(pattern)

}

func check(pattern string) {
	src := "abc123一二三①②③"
	fmt.Println(src)
	reg := regexp.MustCompile(pattern)
	fmt.Println(src)
	dst := reg.ReplaceAllString(src, "*")
	fmt.Println(dst)
}
