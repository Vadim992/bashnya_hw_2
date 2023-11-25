package cli

import (
	"flag"
)

// check cliStruct.go to learn more CLI struct

const DefaultCase = "default" // add this in priority queue if  "-c", "-d" or "-u" flags dont used when we start cli app

func (cli *Cli) setAllFlags() {
	var (
		c, d, u, i bool
		s, f       int
	)

	flag.BoolVar(&c, "c", false,
		"Count the number of occurrences of a string in the input data.Print this number before the string separated by a space.")
	flag.BoolVar(&d, "d", false,
		"Output only those lines that are repeated in the input data.")
	flag.BoolVar(&u, "u", false,
		"Output only those lines that are not repeated in the input data.")
	flag.BoolVar(&i, "i", false,
		"Ignore case of letters.")
	flag.IntVar(&f, "f", 0,
		"Ignore the first 'num_fields' of fields in a row. A field in a string is a non-empty set of characters separated by a space.")
	flag.IntVar(&s, "s", 0,
		"Ignore the first 'num_chars' characters in the string.")

	flag.Parse()

	cli.F = f
	cli.S = s

	cli.usageFlags()

}

// usageFlags func fill map cli.Flags, validate all flags,  return the priotity queue of flags ("-i" - the highest priority)
// check cliStruct.go to learn more about priority of flags

func (cli *Cli) usageFlags() {
	numSetFlags := flag.NFlag()

	if numSetFlags == 0 {
		cli.appendDefaultCase(numSetFlags + 1)
		return
	}

	m := make(map[string]struct{}, numSetFlags)
	mUniq := make(map[string]struct{}, len(cli.UniqFlags))

	flag.Visit(func(f *flag.Flag) {
		m[f.Name] = struct{}{}

		if _, ok := cli.UniqFlags[f.Name]; ok {
			mUniq[f.Name] = struct{}{}
		}
	})

	err := cli.CheckFlagsCDU(mUniq)

	if err != nil {
		cli.ErrorLog.Fatal(err)
	}

	cli.fillQueue(m, len(mUniq), numSetFlags)

}

// StartCLI starts CLI app
func (cli *Cli) StartCLI() {

	cli.setAllFlags()
	arr := flag.Args()
	res, err := readData(arr)

	if err != nil {
		cli.ErrorLog.Fatal(err)
	}

	buf := cli.Uniq(res)

	if err = outData(arr, buf); err != nil {
		cli.ErrorLog.Fatal(err)
	}

}
