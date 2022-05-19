package main

func main() {

}


type LFUCache struct {
	recent   map[int]*listNode // freq - first node matched freq
	count    map[int]int       // freq - amount
	cache    map[int]*listNode // key - node
	head     *listNode
	tail     *listNode
	capacity int
}

func Constructor5(capacity int) LFUCache {
	l:= LFUCache{
		recent:   make(map[int]*listNode, capacity),
		count:    make(map[int]int),
		cache:    make(map[int]*listNode, capacity),
		head:     initListNode(0,0),
		tail:     initListNode(0,0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func initListNode(key,val int)*listNode{
	return &listNode{
		key: key,
		value: val,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}

	node, ok := lfu.cache[key]
	if !ok { // key不存在
		return -1
	}

	// key已存在
	next := node.next

	if lfu.count[node.frequency+1] > 0 {
		// 存在其他使用次数为n+1的缓存，将指定缓存移动到所有使用次数为n+1的节点之前
		removeNode(node)
		addBefore(lfu.recent[node.frequency+1], node)
	} else if lfu.count[node.frequency] > 1 && lfu.recent[node.frequency] != node {
		// 不存在其他使用次数为n+1的缓存，但存在其他使用次数为n的缓存，且当前节点不是最近的节点
		// 将指定缓存移动到所有使用次数为n的节点之前
		removeNode(node)
		addBefore(lfu.recent[node.frequency], node)
	}

	// 更新recent
	lfu.recent[node.frequency+1] = node
	if lfu.count[node.frequency] <= 1 { // 不存在其他freq = n的节点，recent置空
		lfu.recent[node.frequency] = nil
	} else if lfu.recent[node.frequency] == node { // 存在其他freq = n的节点，且recent = node，将recent向后移动一位
		lfu.recent[node.frequency] = next
	}

	// 更新使用次数对应的节点数
	lfu.count[node.frequency+1]++
	lfu.count[node.frequency]--

	// 更新缓存使用次数
	node.frequency++

	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	node, ok := lfu.cache[key]
	if ok { // key已存在
		lfu.Get(key)
		node.value = value

		return
	}

	// key不存在
	if len(lfu.cache) >= lfu.capacity { // 缓存已满，删除最后一个节点，相应更新cache、count、recent（条件）
		tailNode := lfu.tail.prev

		lfu.removeTail()

		if lfu.count[tailNode.frequency] <= 1 {
			lfu.recent[tailNode.frequency] = nil
		}
		lfu.count[tailNode.frequency]--
		delete(lfu.cache, tailNode.key)
	}

	newNode := &listNode{
		key:       key,
		value:     value,
		frequency: 1,
	}

	// 插入新的缓存节点
	if lfu.count[1] > 0 {
		addBefore(lfu.recent[1], newNode)
	} else {
		addBefore(lfu.tail, newNode)
	}

	// 更新recent、count、cache
	lfu.recent[1] = newNode
	lfu.count[1]++
	lfu.cache[key] = newNode
}


type listNode struct {
	key       int
	value     int
	frequency int
	prev      *listNode
	next      *listNode
}

func removeNode(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func addBefore(currNode *listNode, newNode *listNode) {
	pre := currNode.prev

	pre.next = newNode
	newNode.next = currNode

	currNode.prev = newNode
	newNode.prev = pre
}

func (this *LFUCache) removeTail() {
	node := this.tail.prev.prev
	removeNode(node)
}

func (this *LFUCache) isEmpty() bool {
	return this.head.next == this.tail
}




type LFUCache1 struct {
	keyToVal map[int]int
	keyToFreq map[int]int
	freqToKeys map[int]map[int]int
	minFreq int
	cap int
}

func Constructor4(capacity int) LFUCache1 {
	cache:=LFUCache1{
		keyToVal: map[int]int{},
		keyToFreq: map[int]int{},
		freqToKeys: map[int]map[int]int{},
		minFreq: 0,
		cap: capacity,
	}
	return cache
}

func (self *LFUCache1) get(key int) int{
	if _,ok:=self.keyToVal[key];!ok{
		return -1
	}
	self.increaseFreq(key)
	return self.keyToVal[key]
}

func (self *LFUCache1) put(key,val int){
	if self.cap <= 0  {
		return
	}
	//如果存在就更新val，并增加key的freq值
	if _,ok:=self.keyToVal[key];ok{
		self.keyToVal[key] = val
		self.increaseFreq(key)
		return
	}
	if len(self.keyToVal) >= self.cap {
		self.removeMinFreqKey()
	}
	self.keyToVal[key] = val
	self.keyToFreq[key] = 1
	self.freqToKeys[1][key] = 1
	self.minFreq =1

}

func (self *LFUCache1)removeMinFreqKey(){
	keys:=self.freqToKeys[self.minFreq]
	deleteKey:=keys[0]
	delete(keys,deleteKey)
	if len(keys) == 0 {
		delete(self.freqToKeys,self.minFreq)
	}
	delete(self.keyToVal,deleteKey)
	delete(self.keyToFreq,deleteKey)
}


func (self *LFUCache1)increaseFreq(key int){
	freq:=self.keyToFreq[key]
	//更新key的freq
	self.keyToFreq[key]+=1
	//将之前freq对应的key值删掉
	delete(self.freqToKeys[freq],key)
	//将key加入到新的freqtokeys中
	self.freqToKeys[freq+1][key] =1
	if len(self.freqToKeys[freq]) == 0 {
		delete(self.freqToKeys,freq)
		// 如果这个freq是minfreq 则更新
		if freq == self.minFreq {
			self.minFreq++
		}
	}
}