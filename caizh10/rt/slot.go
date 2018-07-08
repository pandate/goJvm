package rt

import "./heap"
/*
定义一个结构体，让它可以同时容纳一个int值和一个引用值。
 */

type Slot struct {
	num int32
	ref *heap.Object
}