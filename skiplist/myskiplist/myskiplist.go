package myskiplist

import (
	"math/rand"
	"time"
)

type node struct {
	key int
	value int
	forward []*node
}

type SkipList struct {
	level int
	hand *node
}

func NewNodeOfLevel(level int)(*node){
	l := make([]*node, level)
	return &node{forward:l}
}

const (
	MaxNumberOfLevels = 11
	MaxLevel          = 10
)

func NewSkipList()(*SkipList){
	var(
		l *SkipList
	)

	l = &SkipList{}
	l.level = 0
	f := NewNodeOfLevel(MaxNumberOfLevels)
	for i := 0;i < MaxNumberOfLevels;i++{
		f.forward[i] = nil
	}
	l.hand = f

	return l
}

func randomLevel() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(MaxLevel)
}

func inster(list *SkipList, key,value int) bool{
	var(
		k int
		p, q *node
		update [MaxNumberOfLevels]*node
	)
	k = list.level
	p = list.hand

	for ;k >=0 ;k--{
		q = p.forward[k]
		for q != nil && q.key < key{
			p = q
			q = p.forward[k]
		}
		update[k] = p
	}

	if q!= nil && q.value == value{
		return false
	}

	k = randomLevel()
	if k > list.level{
		list.level ++
		k = list.level
		update[k] = list.hand
	}

	q = NewNodeOfLevel(k +1)
	q.key = key
	q.value = value

	// 每层插入节点
	for ; k >= 0; k-- {
		p = update[k]
		q.forward[k] = p.forward[k]
		p.forward[k] = q
	}

	return true
}


func search(list *SkipList, key int) int {
	var(
		k int
		p,q * node
	)

	k = list.level
	p = list.hand
	for ;k >=0;k--{
		q = p.forward[k]
		for q != nil && q.key < key{
			q = q.forward[k]
		}

		if q != nil && q.key == key{
			return q.value
		}
	}

	if q != nil && q.key == key {
		return q.value
	} else {
		return -1
	}
}