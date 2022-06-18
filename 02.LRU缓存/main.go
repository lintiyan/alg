package main

import "fmt"

/*

	请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
	实现 LRUCache 类：
	LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
	int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
	函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

 

	示例：

	输入
	["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	输出
	[null, null, null, 1, null, -1, null, -1, 3, 4]

	解释
	LRUCache lRUCache = new LRUCache(2);
	lRUCache.put(1, 1); // 缓存是 {1=1}
	lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
	lRUCache.get(1);    // 返回 1
	lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.get(2);    // 返回 -1 (未找到)
	lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.get(1);    // 返回 -1 (未找到)
	lRUCache.get(3);    // 返回 3
	lRUCache.get(4);    // 返回 4
 

	提示：

	1 <= capacity <= 3000
	0 <= key <= 10000
	0 <= value <= 105
	最多调用 2 * 105 次 get 和 put

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/lru-cache
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/


type LRUCache struct {
	Cap int						// 容量
	Size int					// 已存入数量
	CacheMap map[int]*Node		// hash，存取O(1)
	head, tail *Node			// 头尾指针，初始化，方便链表操作，使用这之后不用考虑链表边界
}

type Node struct {
	Key int
	Value int
	PreNode *Node
	NextNode *Node
}

func initNode(key, value int) *Node{
	return &Node{
		Key: key,
		Value: value,
	}
}

func Constructor(capacity int) LRUCache {

	lru := LRUCache{
		Cap: capacity,
		Size: 0,
		CacheMap: make(map[int]*Node),
		head: initNode(0,0),
		tail: initNode(0,0),
	}

	lru.head.NextNode = lru.tail
	lru.tail.PreNode = lru.head

	return lru
}


func (this *LRUCache) Get(key int) int {
	node ,ok := this.CacheMap[key]
	if !ok {
		// 不存在直接返回-1
		return -1
	}
	// 存在，则将节点移动到队列头
	this.MoveToHead(node)
	return node.Value
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.CacheMap[key]
	if !ok {
		// 不存在，则添加到队列头
		newNode := initNode(key,value)
		this.CacheMap[key] = newNode
		this.AddToHead(newNode)
		this.Size++
		// 判断是否超过容量
		if this.Size > this.Cap {
			// 超过容量，删除链表尾部节点
			tail := this.tail.PreNode
			this.RemoveNode(tail)
			delete(this.CacheMap, tail.Key)
			this.Size--
		}
		return
	}
	// 存在，则删除旧节点，并将新节点添加到队列头
	this.RemoveNode(node)
	delete(this.CacheMap, key)
	newNode := initNode(key, value)
	this.AddToHead(newNode)
	this.CacheMap[key] = newNode
}

func (this *LRUCache) RemoveNode(node *Node){
	node.PreNode.NextNode = node.NextNode
	node.NextNode.PreNode = node.PreNode
}

func (this *LRUCache) AddToHead(node *Node){
	node.NextNode = this.head.NextNode
	this.head.NextNode.PreNode = node
	this.head.NextNode = node
	node.PreNode = this.head
}

// 移动节点到头部 = 删除节点+ 添加节点到头部
func (this *LRUCache) MoveToHead(node *Node){
	this.RemoveNode(node)
	this.AddToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TraverseMap(target map[int]*Node){
	for k, v := range target{
		fmt.Println("key=", k, "value=", v.Value)
	}
	fmt.Println("===============")
}

func main(){
	obj := Constructor(2)
	obj.Put(1,1)
	TraverseMap(obj.CacheMap)
	obj.Put(2,2)
	TraverseMap(obj.CacheMap)
	obj.Put(3,3)
	TraverseMap(obj.CacheMap)
}

