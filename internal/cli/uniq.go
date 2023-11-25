package cli

import (
	"bytes"
	"fmt"
	"strings"
)

// Я создаю массив (приоритетеную очередь),
//отсротированный потприоритетам обработки флагов: cli.QueueFlags
// флаги обрабатываются в слудющей последовательности: i, f, s, затем остальные
// массив copyLines содержит копию элементов оригинального массива
//(оригинальный массив - массив строк входных данных,
// один элемент этого массива - это строка во входных данных)

// "Эталонный вид" массива - это массив состоящий из строк в которых нет пробелов вокруг (слева и справа),
//а внутри слова разделены одним пробелом

// Пример: после считывания данных получаю массив ["   I    love   ", "I    hate       "],
// Далее создаю копию этого массива (исходные данные не меняю)
// После преобразований к эталонному виду послучаю массив ["I love", "I hate"]
// В выходных данных уже использую элементы ИСХОДНОГО массива, то есть
// если программа была запущена с флагом -с, то на выходе получу
// 1    I    love
// 1 I    hate

// Таким образом в логику программы я заложил сравнение толоько СЛОВ (все пробелы, знаки табуляции и тд не учитывал)

func (cli *Cli) Uniq(lines []string) *bytes.Buffer {

	copyLines := make([]string, len(lines))

	copy(copyLines, lines)           // создаю массив копий
	copyLines = trimSpace(copyLines) // привожу массив копий к "эталонному виду"

	for len(cli.QueueFlags) > 0 {

		item := cli.QueueFlags[len(cli.QueueFlags)-1]

		switch item.Value {
		case "c":

			copyLines = HandleC(lines, copyLines)

		case "d":

			copyLines = HandleD(lines, copyLines)

		case "u":

			copyLines = HandleU(lines, copyLines)

		case "i":

			copyLines = HandleI(copyLines)

		case "f":

			copyLines = HandleF(copyLines, cli.F)

		case "s":

			copyLines = HandleS(copyLines, cli.S)

		case DefaultCase:

			copyLines = HandleDefault(lines, copyLines)
		}

		cli.deleteLast()
	}

	return fillBuffer(copyLines)

}

// fillBuffer fills bytes.Buffer that I used to out data
func fillBuffer(lines []string) *bytes.Buffer {
	var b bytes.Buffer
	for _, val := range lines {

		b.WriteString(fmt.Sprintf("%s\r\n", val))
	}
	return &b
}

// trimSpace trims space around ands inside string
// (use d this func to create canonical slise of string)
func trimSpace(lines []string) []string {
	lines = trimSpaceAround(lines)
	for i := range lines {

		strSlice := strings.Split(lines[i], " ")
		strSlice = trimSpaceInside(strSlice)

		lines[i] = strings.Join(strSlice, " ")
	}

	return lines
}
