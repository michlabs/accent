package accent

import (
	"bufio"
	"fmt"
	"os"
)

var sourceCharacters = []rune{
			'À', 'Á', 'Â', 'Ã', 'È', 'É', 'Ê', 'Ì', 'Í', 'Ò', 'Ó', 'Ô', 'Õ', 
			'Ù', 'Ú', 'Ý', 'à', 'á', 'â', 'ã', 'è', 'é', 'ê', 'ì', 'í', 'ò', 
			'ó', 'ô', 'õ', 'ù', 'ú', 'ý', 'Ă', 'ă', 'Đ', 'đ', 'Ĩ', 'ĩ', 'Ũ', 
			'ũ', 'Ơ', 'ơ', 'Ư', 'ư', 'Ạ', 'ạ', 'Ả', 'ả', 'Ấ', 'ấ', 'Ầ', 'ầ', 
			'Ẩ', 'ẩ', 'Ẫ', 'ẫ', 'Ậ', 'ậ', 'Ắ', 'ắ', 'Ằ', 'ằ', 'Ẳ', 'ẳ', 'Ẵ', 
			'ẵ', 'Ặ', 'ặ', 'Ẹ', 'ẹ', 'Ẻ', 'ẻ', 'Ẽ', 'ẽ', 'Ế', 'ế', 'Ề', 'ề',
			'Ể', 'ể', 'Ễ', 'ễ', 'Ệ', 'ệ', 'Ỉ', 'ỉ', 'Ị', 'ị', 'Ọ', 'ọ', 'Ỏ', 
			'ỏ', 'Ố', 'ố', 'Ồ', 'ồ', 'Ổ','ổ', 'Ỗ', 'ỗ', 'Ộ', 'ộ', 'Ớ', 'ớ', 
			'Ờ', 'ờ', 'Ở', 'ở', 'Ỡ', 'ỡ', 'Ợ', 'ợ', 'Ụ', 'ụ', 'Ủ', 'ủ', 'Ứ', 
			'ứ', 'Ừ', 'ừ', 'Ử', 'ử', 'Ữ', 'ữ', 'Ự', 'ự'}

var destinationCharacters = []rune{
			'A', 'A', 'A', 'A', 'E', 'E', 'E', 'I', 'I', 'O', 'O', 'O', 'O', 
			'U', 'U', 'Y', 'a', 'a', 'a', 'a', 'e', 'e', 'e', 'i', 'i', 'o', 
			'o', 'o', 'o', 'u', 'u', 'y', 'A', 'a', 'D', 'd', 'I', 'i', 'U', 
			'u', 'O', 'o', 'U', 'u', 'A', 'a', 'A', 'a', 'A', 'a', 'A', 'a', 
			'A', 'a', 'A', 'a', 'A', 'a', 'A', 'a', 'A', 'a', 'A', 'a', 'A', 
			'a', 'A', 'a', 'E', 'e', 'E', 'e', 'E', 'e', 'E', 'e', 'E', 'e', 
			'E', 'e', 'E', 'e', 'E', 'e', 'I', 'i', 'I', 'i', 'O', 'o', 'O', 
			'o', 'O', 'o', 'O', 'o', 'O', 'o', 'O', 'o', 'O', 'o', 'O', 'o', 
			'O', 'o', 'O', 'o', 'O', 'o', 'O', 'o', 'U', 'u', 'U', 'u', 'U', 
			'u', 'U', 'u', 'U', 'u', 'U', 'u', 'U', 'u', }

var mappingTable map[rune]rune = make(map[rune]rune)

func init() {
	for i, s := range sourceCharacters {
		mappingTable[s] = destinationCharacters[i]
	}
}

// Remove removes Vietnamese accents from a string
func Remove(s string) string {
	arr := []rune(s)
	for i, r := range arr {
		if v, ok := mappingTable[r]; ok {
			arr[i] = v
		}
	}
	return string(arr)
}

// RemoveFromFile removes Vietnamese accents from an input file and 
// writes result to the output file
func RemoveFromFile(inputFP, outputFP string) error {
	input, err := os.Open(inputFP)
    if err != nil {
        return fmt.Errorf("failed to read input file. Error: %s", err.Error())
    }
    defer input.Close()

    output, err := os.OpenFile(outputFP, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0775)
    if err != nil {
    	return fmt.Errorf("failed to create output file. Error: %s", err.Error())
    }
    defer output.Close()

    scanner := bufio.NewScanner(input)
    writer := bufio.NewWriter(output)
    defer writer.Flush()
    for scanner.Scan() {
        fmt.Fprintln(writer, Remove(scanner.Text()))
    }

    if err := scanner.Err(); err != nil {
        return fmt.Errorf("failed to scan input file %s. Error: %s", inputFP, err.Error())
    }

    return nil
}