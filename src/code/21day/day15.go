package main

func main() {

}

//双向链表
type DLinkedNode struct {
	key int
	value int
	prev *DLinkedNode
	next *DLinkedNode
}

func initDLinkedNode(key,value int)*DLinkedNode{
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head *DLinkedNode
	tail *DLinkedNode
}




//使用伪头部和伪尾部标记界限
func Constructor(capacity int) LRUCache {
	cache:=LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0,0),
		tail: initDLinkedNode(0,0),
		capacity: capacity,
	}
	//头指向尾 尾指向头
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}


func (this *LRUCache) Get(key int) int {
	if _,ok:=this.cache[key];!ok{
		return -1
	}
	node:=this.cache[key]
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok:=this.cache[key];!ok{
		node:=initDLinkedNode(key,value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			deleteNode:=this.removeTail()
			delete(this.cache,deleteNode.key)
			this.size--
		}
	}else{
		node:=this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) removeNode(node *DLinkedNode){
	//将当前节点上一个的下一个变成 当前节点的下一下
	node.prev.next = node.next
	//将当前节点下一个的上一个变成 当前节点的上一个
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkedNode){
	//添加到头部 上一个节点就是this.head
	node.prev = this.head
	node.next = this.head.next
	//将头部节点的下一个改成当前节点，之前这个位置的节点的上一个改成当前节点
	this.head.next.prev = node
	this.head.next = node

}

//将已经存在的节点移动到头部
func (this *LRUCache) moveToHead(node *DLinkedNode){
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail()*DLinkedNode{
	node:=this.tail.prev
	this.removeNode(node)
	return node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */