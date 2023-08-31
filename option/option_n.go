package option

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Option_n(chunks string, file *os.File) error {
	inOption := strings.Split(chunks, "/")
	argsLen := len(inOption)
	if 1 > argsLen || argsLen > 3 {
		return fmt.Errorf("[ERROR] 引数無効")
	}

	if err := validateArgs(inOption[1:]); err != nil {
		return err
	}

	if inOption[0] == "r" {
		return fmt.Errorf("option 'r' 未実装")
	}
	if inOption[0] == "l" {
		return fmt.Errorf("option 'l' 未実装")
	}

	intOption, err := strconv.Atoi(inOption[0])
	if err != nil {
		return err
	}

	if argsLen == 1 {
		return option_n_default(intOption, file)
	} else {
		return option_n_k(intOption)
	}
}

func validateArgs(args []string) error {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		_, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func option_n_default(splitNum int, file *os.File) error {
	fmt.Println("option_n_default")
	// xaa ~ xzz まで対応
	if splitNum < 1 || splitNum > 676 {
		return fmt.Errorf("[ERROR] 分割数無効")
	}

	var prefix string = "x"
	idx := []int{0, 0}
	asciiInt := 97 // `a`
	fileInfo, _ := file.Stat()
	splitSize := make([]byte, int64(math.Ceil(float64(fileInfo.Size())/float64(splitNum))))

	for {
		n, err := file.Read(splitSize)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		outputFileName := fmt.Sprintf("%s%s%s", prefix, string(rune(asciiInt+idx[1])), string(rune(asciiInt+idx[0])))
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		_, err = outputFile.Write(splitSize[:n])
		if err != nil {
			return err
		}

		idx[0]++
		if idx[0] == 26 {
			idx[0] = 0
			idx[1]++
		}
	}

	return nil
}

func option_n_k(k int) error {
	return nil
}
