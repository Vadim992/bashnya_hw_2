package cli

import (
	priorityqueue "AppCLI/internal/priorityQueue"
	"sort"
)

// fillQueue fills queue - max element placed in the end of slice (queue)

func (cli *Cli) fillQueue(m map[string]struct{}, uniqFlag, allFlags int) {

	// uniqFlag is 1 or 0 I used it to know: we used -c. -d. -u flags or not
	if uniqFlag == 0 {
		cli.appendDefaultCase(allFlags + 1)
	} else {
		cli.QueueFlags = priorityqueue.NewQueue(allFlags)
	}

	for key := range m {
		item := priorityqueue.NewItem(key, cli.PriorityMap[key])
		cli.QueueFlags = append(cli.QueueFlags, item)
	}

	sort.SliceStable(cli.QueueFlags, cli.QueueFlags.Less)

}

// appendDefaultCase appends default case (we should do this if we wasnt set  one of rhis flags: -c, -d or -u )

func (cli *Cli) appendDefaultCase(num int) {
	if cli.QueueFlags == nil {
		cli.QueueFlags = priorityqueue.NewQueue(num)
	}

	defaultCase := priorityqueue.NewItem(DefaultCase, priorityqueue.MinPriority)
	cli.QueueFlags = append(cli.QueueFlags, defaultCase)
}

// deleteLast deletes last element from queue (slice)
func (cli *Cli) deleteLast() {

	if len(cli.QueueFlags) == 0 {
		return
	}

	l := len(cli.QueueFlags)

	cli.QueueFlags[l-1] = nil

	cli.QueueFlags = cli.QueueFlags[:l-1]
}
