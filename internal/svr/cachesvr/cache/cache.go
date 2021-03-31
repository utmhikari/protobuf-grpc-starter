package cache

import (
	"bytes"
	"fmt"
	pb "github.com/utmhikari/protobuf-grpc-starter/api/pb/base"
	"log"
	"sync"
)

/**
	A simple implementation of lru cache
	key: shortLink
	value: document node
 */

type Node struct {
	document *pb.Document

	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node

	size int
}

func (l *LinkedList) AppendToFront(n *Node) {
	if n == nil {
		return
	}

	if l.head == nil {
		l.tail = n
	} else {
		n.next = l.head
		l.head.prev = n
	}

	l.head = n
	l.size++
}


func (l*LinkedList) MoveToFront(n *Node) {
	if n == nil || l.head == n {
		return
	}

	prev, next := n.prev, n.next
	prev.next = next
	if l.tail == n {
		l.tail = prev
	}
	if next != nil {
		next.prev = prev
	}
	n.next = l.head
	l.head.prev = n
	n.prev = nil
	l.head = n
}


func (l*LinkedList) RemoveFromTail() {
	if l.tail == nil {
		return
	}

	prev := l.tail.prev
	if prev != nil {
		prev.next = nil
	}

	l.tail.prev = nil
	l.tail = prev
	l.size--
}


func (l*LinkedList) ToString() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintf("<LinkedList(%d)> <%v>:", l.size, l))
	if l.head == nil {
		s.WriteString(" nil")
		return s.String()
	}

	p, cnt := l.head, 0
	for p != nil {
		s.WriteString(fmt.Sprintf("\n\t<%d> document: %+v", cnt, p.document))
		cnt++
		p = p.next
	}

	return s.String()
}

func (l*LinkedList) ToStringReverse() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintf("<LinkedList(%d)> <%v> <Reverse>:", l.size, l))
	if l.tail == nil {
		s.WriteString(" nil")
		return s.String()
	}

	p, cnt := l.tail, -1
	for p != nil {
		s.WriteString(fmt.Sprintf("\n\t<%d> document: %+v", cnt, p.document))
		cnt--
		p = p.prev
	}

	return s.String()
}


type LRUCache struct {
	maxSize int

	nodes LinkedList
	mp    map[string]*Node  // ShortLink -> NodePtr

	mu sync.RWMutex
}


var cache *LRUCache = nil


// Init cache instance
func Init(c *Config) error {
	err := c.Check()
	if err != nil {
		return err
	}

	cache = &LRUCache{maxSize: c.MaxSize}
	return nil
}


func Get(shortLink string) *pb.Document {
	if cache == nil {
		return nil
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	node, ok := cache.mp[shortLink]
	if !ok {
		return nil
	}

	cache.nodes.MoveToFront(node)

	log.Printf("Get %s:\n%s\n%s\n",
		shortLink, cache.nodes.ToString(), cache.nodes.ToStringReverse())
	return node.document
}


func Set(doc *pb.Document) {
	if cache == nil || doc == nil || len(doc.ShortLink) == 0 {
		return
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	if nil == cache.mp {
		cache.mp = make(map[string]*Node)
	}

	node, ok := cache.mp[doc.ShortLink]
	if !ok {
		if cache.nodes.size == cache.maxSize {
			delete(cache.mp, cache.nodes.tail.document.ShortLink)
			cache.nodes.RemoveFromTail()
		}
		newNode := &Node{
			document: doc,
		}
		cache.nodes.AppendToFront(newNode)
		cache.mp[doc.ShortLink] = newNode
	} else {
		node.document = doc
		cache.nodes.MoveToFront(node)
	}

	log.Printf("Set %s:\n%s\n%s\n",
		doc.ShortLink, cache.nodes.ToString(), cache.nodes.ToStringReverse())
}
