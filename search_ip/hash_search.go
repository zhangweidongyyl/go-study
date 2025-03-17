package search_ip

import "fmt"

// 存到内存里，这时要考虑 每个ip地址占用字节数，看是否能存下

func SearchByHash() {
	ipSet := make(map[string]bool)
	ips := []string{"192.168.1.1", "10.0.0.1", "172.16.254.1"}

	for _, ip := range ips {
		ipSet[ip] = true
	}

	fmt.Println(ipSet["192.168.1.1"]) // true
	fmt.Println(ipSet["10.0.0.2"])    // false
}
