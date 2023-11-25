package cli

import (
	testcli "AppCLI/internal/cli"
	"testing"

	"github.com/stretchr/testify/require"
)

// test flag -c
func TestHandleC(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
	}{
		{[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks"},
			[]string{"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks"},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks"},
			"Test01",
		},
		{[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"2 I love music.", "3 ", "2 Hello", "1 Hello my friend", "2 I love music of Kartik.", "1 Thanks"},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"1 Hi my friend", "1 Hello my friend", "1 Ok my friend"},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"2 Hi my friend", "2 Hi", "1 "},
			"Test04",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleC(tt.inOriginalArr, tt.inArr)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}

// test flag -d
func TestHandleD(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music of Kartik."},
			"Test01",
		},
		{[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "", "Hello", "I love music of Kartik."},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi"},
			"Test04",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleD(tt.inOriginalArr, tt.inArr)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}

// test flag -u
func TestHandleU(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"", "Thanks"},
			"Test01",
		},
		{[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"Hello my friend", "Thanks"},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{""},
			"Test04",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleU(tt.inOriginalArr, tt.inArr)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}

// test WITHOUT flags
func TestHandleDefault(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "", "I love music of Kartik.", "Thanks"},
			"Test01",
		},
		{[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "", "Hello", "Hello my friend", "I love music of Kartik.", "Thanks"},
			"Test02",
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			"Test03",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi", ""},
			"Test04",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleDefault(tt.inOriginalArr, tt.inArr)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}

// test flag -f
func TestHandleF(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
		f                    int
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"love music.", "love music.", "love music.", "", "love music of Kartik.", "love music of Kartik.", "Thanks"},
			"Test01",
			1,
		},
		{[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"music.", "music.", "", "", "", "Hello", "Hello", "my friend", "love music of Kartik.", "love music of Kartik.", "Thanks"},
			"Test02",
			2,
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			"Test03",
			5,
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"my friend", "my friend", "Hi", "Hi", ""},
			"Test04",
			1,
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleF(tt.inArr, tt.f)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}

// test flag -s
// Удаляю только символы (пробелы - НЕ символы по моей логике)
func TestHandleS(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
		s                    int
	}{
		{[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"love music.", "love music.", "love music.", "", "love music of Kartik.", "love music of Kartik.", "hanks"},
			"Test01",
			1,
		},
		{[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"I love music.", "I love music.", "", "", "", "Hello", "Hello", "Hello my friend", "I love music of Kartik.", "I love music of Kartik.", "Thanks"},
			[]string{"ove music.", "ove music.", "", "", "", "llo", "llo", "llo my friend", "ove music of Kartik.", "ove music of Kartik.", "anks"},
			"Test02",
			2,
		},
		{[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			[]string{"Hi my friend", "Hello my friend", "Ok my friend"},
			"Test03",
			10000,
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"y friend", "y friend", "Hi", "Hi", ""},
			"Test04",
			3},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleS(tt.inArr, tt.s)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}

// test -i flag
func TestHandleI(t *testing.T) {

	tableTest := []struct {
		inOriginalArr, inArr []string
		outArr               []string
		name                 string
	}{
		{[]string{"I love music.",
			"I loVe music.",
			"I love music.",
			"",
			"I love mUSic of Kartik.",
			"I LOVE MUSIC of kartik.",
			"Thanks"},

			[]string{"I love music.",
				"I loVe music.",
				"I love music.",
				"",
				"I love mUSic of Kartik.",
				"I LOVE MUSIC of kartik.",
				"Thanks"},
			[]string{"i love music.",
				"i love music.",
				"i love music.",
				"",
				"i love music of kartik.",
				"i love music of kartik.",
				"thanks"},
			"Test01",
		},

		{[]string{"Hi MY friend", "HelLo my friend", "Ok my friEnd"},
			[]string{"Hi MY friend", "HelLo my friend", "Ok my friEnd"},
			[]string{"hi my friend", "hello my friend", "ok my friend"},
			"Test02",
		},
		{[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"Hi my friend", "Hi my friend", "Hi", "Hi", ""},
			[]string{"hi my friend", "hi my friend", "hi", "hi", ""},
			"Test04",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			res := testcli.HandleI(tt.inArr)

			require.Equal(t, tt.outArr, res, "Slices should be equal")
		})
	}
}
