
type MinHeap []int
 
 func (h MinHeap) Len() int { return len(h)}
 func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
 func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
 func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
 }
 func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
 }
 
type MaxHeap []int
 
 func (h MaxHeap) Len() int { return len(h)}
 func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
 func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
 }
  func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
 }
type MedianFinder struct {
    lower *MaxHeap
    upper *MinHeap
}


func Constructor() MedianFinder {
    minHeap := &MinHeap{}
    maxHeap := &MaxHeap{}
    heap.Init(minHeap)
    heap.Init(maxHeap)
    return MedianFinder{lower: maxHeap, upper: minHeap}
}

func (m *MedianFinder) AddNum(num int)  {
    heap.Push(m.lower, num)

    val := heap.Pop(m.lower).(int)
    heap.Push(m.upper, val)

    if m.upper.Len() > m.lower.Len() {
        val := heap.Pop(m.upper)
        heap.Push(m.lower, val)
    }
}


func (m *MedianFinder) FindMedian() float64 {
    if m.lower.Len() > m.upper.Len() {
        return float64((*m.lower)[0])
    }
    
    return float64((*m.lower)[0] + (*m.upper)[0])/2.0
}
