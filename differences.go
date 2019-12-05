package microtest

type Differences struct {
	slice []string
}

func (this *Differences) Push(difference string) {
	this.slice = append(this.slice, difference)
}

func (this *Differences) isEmpty() bool {
	return len(this.slice) == 0
}
