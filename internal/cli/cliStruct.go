package cli

import (
	priorityqueue "AppCLI/internal/priorityQueue"
	"AppCLI/internal/validator"
	"log"
	"os"
)

type Cli struct {
	QueueFlags          priorityqueue.Queue // this field contains priority queue of flags (sorted slice of flags)
	PriorityMap         map[string]int      // this map contains priority of flags
	UniqFlags           map[string]struct{} // this map contains flags that we cant use in the same time (flags, -c, -d, -u)
	ErrorLog            *log.Logger         // custom logger fo errors
	F                   int                 // contains the value of -f flag
	S                   int                 // contains the value of -s flag
	validator.Validator                     // check that we dont use flags -c, -d, -u in the same time
}

func NewCLI() *Cli {
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lshortfile)
	uniqFlags := map[string]struct{}{
		"c": {},
		"d": {},
		"u": {},
	}

	priority := map[string]int{
		"i": 3,
		"f": 2,
		"s": 1,
	}

	return &Cli{
		PriorityMap: priority,
		UniqFlags:   uniqFlags,
		ErrorLog:    errorLog,
	}
}
