package queue

/**
No.933 最近的请求次数
思路：
    1、(莫名其妙的题目)
*/

type RecentCounter struct {
	Q *Queue
}


func RecentCounterConstructor() RecentCounter {
	return RecentCounter{Q:New()}
}


func (this *RecentCounter) Ping(t int) int {
	this.Q.Enqueue(t)
	for this.Q.Peek().(int) < t - 3000{
		this.Q.Dequeue()
	}
	return this.Q.Len()
}