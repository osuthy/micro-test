package runner

type DifferenceQueue struct {
	queue []string
}

func (this *DifferenceQueue) Push(difference string) {
	this.queue = append(this.queue, difference)
}

func (this *DifferenceQueue) isEmpty() bool {
	return len(this.queue) == 0
}
