type MedianFinder struct {
    data []float64
}


func Constructor() MedianFinder {
    return MedianFinder{data: make([]float64, 0)}
}


func (this *MedianFinder) AddNum(num int)  {
    this.data = append(this.data, float64(num))
}


func (this *MedianFinder) FindMedian() float64 {
    sort.Float64s(this.data)
    n := len(this.data)
    if n % 2 == 0 {
        res := (this.data[n/2] + this.data[n/2-1])/2
        return res
    } else {
        return this.data[n/2]
    }
}
