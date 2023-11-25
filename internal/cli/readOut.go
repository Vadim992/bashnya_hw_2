package cli

import (
	"AppCLI/customerr"
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

// readData reads data reads data from Stdin or input file (input.txt), first argument is slice from func flags.Args()
// Func return slice of string (data thet we read) and error

func readData(args []string) ([]string, error) {

	var data []string

	if len(args) == 0 {
		data = read(os.Stdin)
		return data, nil
	}

	path := args[0]

	ok, err := isFile(path)

	if !ok {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0777)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	data = read(file)

	return data, nil
}

// isFile check is path is filePath (it also maybe the Dir).

func isFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false, err
	} else if fileInfo.IsDir() {
		return false, customerr.ErrPathFile
	}

	return true, nil

}

// read reads data from input
// This func use bufio.Scanner for reading data

func read(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string

	for scanner.Scan() {

		line := scanner.Text()

		lines = append(lines, line)

	}

	return lines
}

// outData out data in output

func outData(args []string, b *bytes.Buffer) error {

	if len(args) < 2 {
		outResult(os.Stdout, b)
		return nil
	}

	path := args[1]

	ok, err := isFile(path)

	if !ok && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return err
	}
	defer file.Close()

	outResult(file, b)

	return nil

}

// outResult out data in CORRECT (Stdin or file) output

func outResult(w io.Writer, b *bytes.Buffer) {
	b.WriteTo(w)
}
