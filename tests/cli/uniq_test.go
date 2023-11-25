package cli

import (
	internalcli "AppCLI/internal/cli"
	priorityqueue "AppCLI/internal/priorityQueue"
	"bufio"
	"bytes"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var CLI = internalcli.NewCLI()

func scanFromBuffer(b *bytes.Buffer, l int) []string {
	res := make([]string, 0, l)

	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		text := scanner.Text()

		res = append(res, text)
	}
	return res
}

// test -i -c flags

func TestIC(t *testing.T) {

	tableTest := []struct {
		inOriginalArr []string
		outArr        []string
		name          string
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks"},
			"Test01",
		},
		{[]string{"I love MUSIC.", "I love music.", "", "", "", "HeLLo", "HellO", "Hello My friend", "I love mUsic of Kartik.", "I love music of KArtik.", "Thanks"},
			[]string{"2 I love MUSIC.", "3 ", "2 HeLLo", "1 Hello My friend", "2 I love mUsic of Kartik.", "1 Thanks"},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"1 Hi my friend", "1 Hello my friend", "1 Ok my friend"},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"2 Hi my friend", "2 Hi", "1 "},
			"Test04",
		},

		{[]string{"   Hi  my friend   ", "Hi my friend", "Hi", "Hi", ""},
			[]string{"2    Hi  my friend   ", "2 Hi", "1 "},
			"Test05",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {

			CLI.QueueFlags = priorityqueue.Queue{priorityqueue.NewItem("i", CLI.PriorityMap["i"]), priorityqueue.NewItem("c", CLI.PriorityMap["c"])}

			sort.SliceStable(CLI.QueueFlags, CLI.QueueFlags.Less)

			b := CLI.Uniq(tt.inOriginalArr)
			res := scanFromBuffer(b, len(tt.outArr))

			require.Equal(t, tt.outArr, res, "Slices should be equal.")

		})
	}
}

// test -i -d flags
func TestID(t *testing.T) {

	tableTest := []struct {
		inOriginalArr []string
		outArr        []string
		name          string
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music of Kartik."},
			"Test01",
		},
		{[]string{"I love MUSIC.", "I love music.", "", "", "", "HeLLo", "HellO", "Hello My friend", "I love mUsic of Kartik.", "I love music of KArtik.", "Thanks"},
			[]string{"I love MUSIC.", "", "HeLLo", "I love mUsic of Kartik."},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi"},
			"Test04",
		},
		{[]string{"   Hi  my friend   ", "Hi my friend", "Hi", "Hi", ""},
			[]string{"   Hi  my friend   ", "Hi"},
			"Test05",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {

			CLI.QueueFlags = priorityqueue.Queue{priorityqueue.NewItem("i", CLI.PriorityMap["i"]),
				priorityqueue.NewItem("d", CLI.PriorityMap["d"])}

			sort.SliceStable(CLI.QueueFlags, CLI.QueueFlags.Less)

			b := CLI.Uniq(tt.inOriginalArr)
			res := scanFromBuffer(b, len(tt.outArr))

			require.Equal(t, tt.outArr, res, "Slices should be equal.")

		})
	}
}

// test -i -u flags

func TestIU(t *testing.T) {

	tableTest := []struct {
		inOriginalArr []string
		outArr        []string
		name          string
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"", "Thanks"},
			"Test01",
		},
		{[]string{"I love MUSIC.", "I love music.", "", "", "", "HeLLo", "HellO", "Hello My friend", "I love mUsic of Kartik.", "I love music of KArtik.", "Thanks"},
			[]string{"Hello My friend", "Thanks"},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{""},
			"Test04",
		},

		{[]string{"   Hi  my friend   ", "Hi my friend", "Hi", "Hi", ""},
			[]string{""},
			"Test05",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {

			CLI.QueueFlags = priorityqueue.Queue{priorityqueue.NewItem("i", CLI.PriorityMap["i"]),
				priorityqueue.NewItem("u", CLI.PriorityMap["u"])}

			sort.SliceStable(CLI.QueueFlags, CLI.QueueFlags.Less)

			b := CLI.Uniq(tt.inOriginalArr)
			res := scanFromBuffer(b, len(tt.outArr))

			require.Equal(t, tt.outArr, res, "Slices should be equal.")

		})
	}
}

// test -f -s -i flags

func TestFSI(t *testing.T) {

	tableTest := []struct {
		inOriginalArr []string
		outArr        []string
		name          string
		f             int
		s             int
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "", "I love music of Kartik.", "Thanks"},
			"Test01",
			1,
			1,
		},
		{[]string{"I love MUSIC.", "I love music.", "", "", "", "HeLLo", "HellO", "Hello My friend", "I love mUsic of Kartik.", "I love music of KArtik.", "Thanks"},
			[]string{"I love MUSIC.", "", "HeLLo", "Hello My friend", "I love mUsic of Kartik.", "Thanks"},
			"Test02",
			1000,
			3, // если значение f больше, чем число полей, то при работе с флагом s начинаю идти с НАЧАЛА строки, а НЕ после флага f
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend"},
			"Test03",
			1,
			1,
		},
		{[]string{"  Hi    my    friend", "Hi my friend", "     Hi", "Hi", ""},
			[]string{"  Hi    my    friend", "     Hi", ""},
			"Test04",
			1,
			1,
		},

		{[]string{"   Hi  my friend   ", "Hi my friend", "Hi", "Hi", ""},
			[]string{"   Hi  my friend   ", "Hi", ""},
			"Test05",
			2,
			2,
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {

			CLI.QueueFlags = priorityqueue.Queue{priorityqueue.NewItem("f", CLI.PriorityMap["f"]),
				priorityqueue.NewItem("s", CLI.PriorityMap["s"]),
				priorityqueue.NewItem(internalcli.DefaultCase, CLI.PriorityMap[internalcli.DefaultCase]),
				priorityqueue.NewItem("i", CLI.PriorityMap["i"])}
			CLI.F = tt.f
			CLI.S = tt.s

			sort.SliceStable(CLI.QueueFlags, CLI.QueueFlags.Less)

			b := CLI.Uniq(tt.inOriginalArr)
			res := scanFromBuffer(b, len(tt.outArr))

			require.Equal(t, tt.outArr, res, "Slices should be equal.")

		})
	}
}

// test -f -s -i -c flags

func TestFSIC(t *testing.T) {

	tableTest := []struct {
		inOriginalArr []string
		outArr        []string
		name          string
		f             int
		s             int
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks"},
			"Test01",
			1,
			1,
		},
		{[]string{"I love MUSIC.", "I love music.", "", "", "", "HeLLo", "HellO", "Hello My friend", "I love mUsic of Kartik.", "I love music of KArtik.", "Thanks"},
			[]string{"2 I love MUSIC.", "3 ", "2 HeLLo", "1 Hello My friend", "2 I love mUsic of Kartik.", "1 Thanks"},
			"Test02",
			1000,
			3, // если значение f больше, чем число полей, то при работе с флагом s начинаю идти с НАЧАЛА строки, а НЕ после флага f
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"3 Hi my friend"},
			"Test03",
			1,
			1,
		},
		{[]string{"  Hi    my    friend", "Hi my friend", "     Hi", "Hi", ""},
			[]string{"2   Hi    my    friend", "2      Hi", "1 "},
			"Test04",
			1,
			1,
		},

		{[]string{"   Hi  my friend   ", "Hi my friend", "Hi", "Hi", ""},
			[]string{"2    Hi  my friend   ", "2 Hi", "1 "},
			"Test05",
			2,
			2,
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {

			CLI.QueueFlags = priorityqueue.Queue{priorityqueue.NewItem("f", CLI.PriorityMap["f"]),
				priorityqueue.NewItem("s", CLI.PriorityMap["s"]),
				priorityqueue.NewItem("i", CLI.PriorityMap["i"]),
				priorityqueue.NewItem("c", CLI.PriorityMap["c"])}
			CLI.F = tt.f
			CLI.S = tt.s

			sort.SliceStable(CLI.QueueFlags, CLI.QueueFlags.Less)

			b := CLI.Uniq(tt.inOriginalArr)
			res := scanFromBuffer(b, len(tt.outArr))

			require.Equal(t, tt.outArr, res, "Slices should be equal.")

		})
	}
}
