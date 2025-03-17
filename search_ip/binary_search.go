package search_ip

// 二分查找，也是要看是否能存到内存，没法一次性加载就分批加载

func binarySearch(ips []string, target string) bool {
	left, right := 0, len(ips)-1
	for left <= right {
		mid := left + (right-left)/2
		if ips[mid] == target {
			return true
		} else if ips[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
