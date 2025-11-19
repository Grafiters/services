package lib

import (
	"net"
	"strings"
)

func GetIPClient(ip string) string {
	if strings.Contains(ip, ":") {
		return FilterIPv4(ip)
	}
	return ip
}

func FilterIPv4(ip string) string {
	// Kena DOS attack
	// ips, err := net.LookupIP(ip)
	// if err != nil {
	// 	return ip // Return the original IP if any error occurs
	// }

	// for _, ip := range ips {
	// 	if ipv4 := ip.To4(); ipv4 != nil {
	// 		return ipv4.String()
	// 	}
	// }

	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		if ipv4 := parseIP.To4(); ipv4 != nil {
			return ipv4.String()
		}

		return ip
	}

	return ip
}
