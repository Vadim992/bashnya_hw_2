package priorityqueue

const MinPriority = 0

type Item struct {
	Value    string
	Priority int
}

// A Queue - I use this type to realise priority queue
// (in fact it is the slice  sorted that sorted in ascending order)
type Queue []*Item

func NewItem(value string, priority int) *Item {
	return &Item{
		Value:    value,
		Priority: priority,
	}
}

func NewQueue(cap int) Queue {
	queue := make(Queue, 0, cap)
	return queue
}

func (pq Queue) Len() int { return len(pq) }

func (pq Queue) Less(i, j int) bool {

	return pq[i].Priority < pq[j].Priority
}

func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

}
