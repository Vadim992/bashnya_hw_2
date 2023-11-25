package cli

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// All "handle" func handle flags. All handle funcs named like this: handle<Flag_name in uppercase>

func HandleC(origLines, lines []string) []string {

	lineMap := make(map[string]int, len(lines))
	copyLine := make([]string, 0, len(lines))

	result := make([]string, 0, len(lines))

	for ind, line := range lines {

		if _, ok := lineMap[line]; !ok {
			result = append(result, origLines[ind])
			copyLine = append(copyLine, line)
		}
		lineMap[line]++
	}

	for i := range result {

		result[i] = fmt.Sprintf("%d %s", lineMap[copyLine[i]], result[i])
	}

	return result
}

func HandleD(origLines, lines []string) []string {

	lineMap := make(map[string]int, len(lines))

	result := make([]string, 0, len(lines))

	for _, line := range lines {

		lineMap[line]++
	}

	for ind, line := range lines {

		if val := lineMap[line]; val > 1 {
			result = append(result, origLines[ind])
			delete(lineMap, line)

		}
	}

	return result

}

func HandleU(origLines, lines []string) []string {

	lineMap := make(map[string]int, len(lines))

	result := make([]string, 0, len(lines))

	for _, line := range lines {

		lineMap[line]++
	}

	for ind, line := range lines {

		if val := lineMap[line]; val == 1 {
			result = append(result, origLines[ind])
			delete(lineMap, line)
		}
	}

	return result

}

func HandleI(lines []string) []string {

	lines = trimSpaceAround(lines)

	for ind := range lines {

		strSlice := strings.Split(lines[ind], " ")

		lines[ind] = strings.ToLower(strings.Join(strSlice, " "))
	}

	return lines

}

func HandleF(lines []string, num int) []string {

	for ind := range lines {

		if lines[ind] == "" {
			continue
		}

		strSlice := strings.Split(lines[ind], " ")

		if num > len(strSlice) {
			num = len(strSlice)
		}

		strSlice = strSlice[num:]

		if len(strSlice) > 0 {
			lines[ind] = strings.Join(strSlice, " ")
		}

	}

	return lines

}

func HandleS(lines []string, num int) []string {

	for ind := range lines {

		if lines[ind] == "" {
			continue
		}

		strSlice := strings.Split(lines[ind], " ")

		str := deleteNumChar(strSlice, num)

		lines[ind] = str

	}

	return lines

}

// deleteNumChar is called by "handleS" func for delete Char from one row
// for deleting char I use rune
// I delete only chars(symbols), whitespace isnot a char in my logic
// example: if string " I  love " && s=2 => canonic arr ["I","love"] => delte 2 chars => ["ove"] => result (string) "ove"
func deleteNumChar(arr []string, num int) string {

	result := strings.Join(arr, " ")

	for len(arr) > 0 {
		if charNum := utf8.RuneCountInString(arr[0]); charNum < num {
			num -= charNum
			arr = arr[1:]
			continue
		}

		sRune := []rune(arr[0])

		sRune = sRune[num:]
		arr[0] = string(sRune)
		if arr[0] == "" {
			arr = arr[1:]
		}

		break
	}

	if len(arr) > 0 {
		result = strings.Join(arr, " ")
	}
	return result

}

func HandleDefault(origLines, lines []string) []string {

	lineMap := make(map[string]int, len(lines))

	result := make([]string, 0, len(lines))

	for ind, line := range lines {

		if _, ok := lineMap[line]; !ok {
			result = append(result, origLines[ind])

		}
		lineMap[line]++
	}

	return result

}

// trimSpaceInside deletes all spaces inside rows
// example input:"hi   bro" -> use strings.Split - >
//->["hi","","","","bro"] - this slice is argument for this func
// output: ["hi","bro"]

func trimSpaceInside(lines []string) []string {
	result := make([]string, 0)
	for i := range lines {

		if lines[i] != "" {
			result = append(result, lines[i])
		}
	}
	return result
}

// trimSpaceInside deletes all spaces around rows
// example input (line[i]): "  hi   bro   " ->
//-> output for ONE row:"hi   bro" (line[i])

func trimSpaceAround(lines []string) []string {
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}
