package getExternalIP

import (
	"regexp"
	"strings"
	"testing"
)

// Check for valid IP
func validIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}

func TestGetIP(t *testing.T) {

	IP := GetIP()
	if !(validIP4(IP)) {
		t.Error("get invalid IP, got ", IP)
	}
}

func TestGetIPinfo(t *testing.T) {

	IPinfo := GetIPinfo()
	if !(validIP4(IPinfo.IP)) {
		t.Error("get invalid IP, got ", IPinfo.IP)
	}
	if len(IPinfo.City) < 2 {
		t.Error("get invalid city, got ", IPinfo.City)
	}
}
