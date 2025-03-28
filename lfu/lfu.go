package lfu

type LFUCache struct {
	keyToVal   map[int]int   // kv表，存储key和值的映射
	keyToFreq  map[int]int   // kf表，存储key和频率的映射
	freqToKeys map[int][]int // fk表，存储 频率 到key的 list 映射,同样频率的key 映射关系存储到freqToKeys map value 为 栈，最早的在栈低，最晚的在栈头
	minFreq    int           // 最小频次
	cap        int           // 容量
}

func NewLFUCache(cap int) *LFUCache {
	lfuCache := LFUCache{
		keyToVal:   make(map[int]int),
		keyToFreq:  make(map[int]int),
		freqToKeys: make(map[int][]int),
		minFreq:    0,
		cap:        cap,
	}
	return &lfuCache
}
func (this *LFUCache) Get(k int) (v int) {

	if v, ok := this.keyToVal[k]; ok {
		this.keyToFreq[k] = this.keyToFreq[k] + 1
		return v
	}
	return -1

}
func (this *LFUCache) Put(k, v int) bool {

	if _, ok := this.keyToVal[k]; ok {
		this.keyToFreq[k] = this.keyToFreq[k] + 1
		return true
	}

	// 此时已经满了 : 删除使用频率最小的且最老的key
	if this.cap <= len(this.keyToVal) {

		keyList := this.freqToKeys[this.minFreq]
		needRemoveKey := keyList[0]

		keyList = keyList[1:]

		// 删除
		this.freqToKeys[this.minFreq] = keyList
		delete(this.keyToVal, needRemoveKey)
		delete(this.keyToFreq, needRemoveKey)

	}
	// 插入
	this.keyToVal[k] = v
	this.keyToFreq[k] = 1
	this.minFreq = 1

	newKeyList := this.freqToKeys[this.minFreq]
	newKeyList = append(newKeyList, k)
	this.freqToKeys[this.minFreq] = newKeyList

	return true
}
