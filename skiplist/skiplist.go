package main

import (
	"fmt"
	"math/rand"
	"time"
)

//定义每个节点类型：
type nodeStructure struct {
	key     int // key值
	value   int // value值
	forward []*nodeStructure
}

// 定义跳表数据类型
type listStructure struct {
	level  int            /* Maximum level of the list (1 more than the number of levels in the list) */
	header *nodeStructure /* pointer to header */
}

const (
	MaxNumberOfLevels = 11
	MaxLevel          = 10
)

// newNodeOfLevel生成一个nodeStructure结构体，同时生成l个*nodeStructure数组指针
//#define newNodeOfLevel(l) (*nodeStructure)malloc(sizeof(struct nodeStructure)+(l)*sizeof(nodeStructure *))

func newNodeOfLevel(level int) *nodeStructure {
	nodearr := make([]*nodeStructure, level) //new([level]*node)
	return &nodeStructure{forward: nodearr}
}

//跳表初始化
func newList() *listStructure {
	var l *listStructure
	var i int
	// 申请list类型大小的内存
	l = &listStructure{}
	// 设置跳表的层level，初始的层为0层（数组从0开始）
	l.level = 0

	// 生成header部分
	l.header = newNodeOfLevel(MaxNumberOfLevels)
	// 将header的forward数组清空
	for i = 0; i < MaxNumberOfLevels; i++ {
		l.header.forward[i] = nil
	}
	return l
}

func randomLevel() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(MaxLevel)
}

func insert(l *listStructure, key int, value int) bool {
	var k int
	// 使用了update数组
	var update [MaxNumberOfLevels]*nodeStructure
	var p, q *nodeStructure
	p = l.header
	k = l.level

	for ; k >= 0; k-- {
		// 查找插入位置
		q = p.forward[k]
		for q != nil && q.key < key {
			p = q
			q = p.forward[k]
		}

		// 设置update数组
		update[k] = p

	} // 对于每一层进行遍历 一直到最低层

	// 这里已经查找到了合适的位置，并且update数组已经
	// 填充好了元素 貌似不插入重复元素
	if q != nil && q.key == key {
		q.value = value
		return false
	}

	// 随机生成一个层数
	k = randomLevel()

	if k > l.level {
		// 如果新生成的层数比跳表的层数大的话
		// 增加整个跳表的层数
		l.level++
		k = l.level
		// 在update数组中将新添加的层指向l->header
		update[k] = l.header
	}

	// 生成层数个节点数目
	q = newNodeOfLevel(k + 1)
	q.key = key
	q.value = value

	// 每层插入节点
	for ; k >= 0; k-- {
		p = update[k]
		q.forward[k] = p.forward[k]
		p.forward[k] = q
	}

	// 如果程序运行到这里，程序已经插入了该节点
	return true
}

func delete(l *listStructure, key int) bool {
	var k, m int
	// 生成一个辅助数组update
	var update [MaxNumberOfLevels]*nodeStructure
	var p, q *nodeStructure
	p = l.header
	k = l.level
	m = l.level

	//指向该节点对应层的前驱节点 一直到最低层
	for ; k >= 0; k-- {

		q = p.forward[k]
		for q != nil && q.key < key {
			p = q
			q = p.forward[k]
		}
		update[k] = p

	}

	// 如果找到了该节点，才进行删除的动作
	if q != nil && q.key == key {
		// 指针运算
		for k = 0; k <= m && update[k].forward[k] == q; k++ {
			// 这里可能修改l.header.forward数组的值的
			p = update[k]
			p.forward[k] = q.forward[k]
		}

		// 如果删除的是最大层的节点，那么需要重新维护跳表的
		// 层数level
		for l.header.forward[m] == nil && m > 0 {
			m--
		}

		l.level = m
		return true

	} else {
		return false

	}
}

func search(l *listStructure, key int) int {
	var k int

	var p, q *nodeStructure
	p = l.header
	k = l.level

	//指向该节点对应层的前驱节点 一直到最低层
	for ; k >= 0; k-- {

		q = p.forward[k]
		for q != nil && q.key < key {
			q = q.forward[k]
		}

		if q != nil && q.key == key {
			return q.value
		}
	}

	if q != nil && q.key == key {
		return q.value
	} else {
		return -1
	}
}

func main() {
	l := newList()

	insert(l, 3, 3)
	insert(l, 6, 6)
	insert(l, 7, 7)
	insert(l, 9, 9)
	insert(l, 10, 10)
	insert(l,13 , 13)
	insert(l, 18, 18)
	insert(l, 8, 8)

	fmt.Printf("skiplist:%v\n", search(l, 12))
	fmt.Printf("skiplist:%v\n", search(l, 9))
}