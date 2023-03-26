package jache

import "container/list"

func New(maxBytes int64, onEvicted func(string, Value)) *Jache {
	return &Jache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (j *Jache) Get(key string) (value Value, ok bool) {
	if ele, ok := j.cache[key]; ok {
		j.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (j *Jache) RemoveOldest() {
	ele := j.ll.Back()
	if ele != nil {
		j.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(j.cache, kv.key)
		j.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if j.OnEvicted != nil {
			j.OnEvicted(kv.key, kv.value)
		}
	}
}

func (j *Jache) Add(key string, value Value) {
	if ele, ok := j.cache[key]; ok {
		j.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		j.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := j.ll.PushFront(&entry{key, value})
		j.cache[key] = ele
		j.nbytes += int64(len(key)) + int64(value.Len())
	}

	for j.maxBytes != 0 && j.maxBytes < j.nbytes {
		j.RemoveOldest()
	}
}
