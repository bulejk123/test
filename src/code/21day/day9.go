package main

func main() {

}


type MyStack struct {
	queue []int
}


func Constructor2() MyStack {
	return MyStack{queue: make([]int, 0)}
}


func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}


func (this *MyStack) Pop() int {
	pop := this.queue[len(this.queue)-1]
	this.queue = this.queue[0:len(this.queue)-1]
	return pop
}


func (this *MyStack) Top() int {
	pop := this.queue[len(this.queue)-1]
	return pop
}


func (this *MyStack) Empty() bool {
	return len(this.queue) <= 0
}

type MyQueue struct {
	queue []int
}


func Constructor3() MyQueue {
	return MyQueue{queue: make([]int, 0)}
}


func (this *MyQueue) Push(x int)  {
	this.queue = append(this.queue, x)
}


func (this *MyQueue) Pop() int {
	p := this.queue[0]
	this.queue = this.queue[1:len(this.queue)]
	return p
}


func (this *MyQueue) Peek() int {
	return this.queue[0]
}


func (this *MyQueue) Empty() bool {
	return len(this.queue) <=0
}
