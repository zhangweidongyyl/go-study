package search_ip

import "net"

// IPv4ToInt64 ipv4 4个字节，每个字节8个bit位，所以要左移到高位
func IPv4ToInt64(ip net.IP) int64 {
	ip = ip.To4() // 确保是 IPv4 地址
	if ip == nil {
		return 0 // 如果不是 IPv4 地址，返回 0
	}
	return int64(ip[0])<<24 | int64(ip[1])<<16 | int64(ip[2])<<8 | int64(ip[3])
}

// IPv6ToInt64s 8个字节，分高位低位
func IPv6ToInt64s(ip net.IP) (int64, int64) {
	ip = ip.To16() // 确保是 IPv6 地址
	if ip == nil {
		return 0, 0 // 如果不是 IPv6 地址，返回 0
	}
	// 高 64 位
	high := int64(ip[0])<<56 | int64(ip[1])<<48 | int64(ip[2])<<40 | int64(ip[3])<<32 |
		int64(ip[4])<<24 | int64(ip[5])<<16 | int64(ip[6])<<8 | int64(ip[7])
	// 低 64 位
	low := int64(ip[8])<<56 | int64(ip[9])<<48 | int64(ip[10])<<40 | int64(ip[11])<<32 |
		int64(ip[12])<<24 | int64(ip[13])<<16 | int64(ip[14])<<8 | int64(ip[15])
	return high, low
}
