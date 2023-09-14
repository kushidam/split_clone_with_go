package option

import (
	"fmt"
	"io"
	"os"
)
func createOutputFile_b(prefix string, idx []int) (*os.File, error) {
	asciiOffset := 97 // 'a'
	asciiSuffix := []string{string(rune(asciiOffset + idx[0])), string(rune(asciiOffset + idx[1]))}
	outputFileName := fmt.Sprintf("%s%s%s", prefix, asciiSuffix[1], asciiSuffix[0])

    outputFile, errOsCreate := os.Create(outputFileName)
    if errOsCreate != nil {
        return nil, errOsCreate
    }
    return outputFile, nil
}

func writeToFile(file *os.File, content []byte) error {
    _, err := file.Write(content)
    return err
}

func writeFinalFile(fileName string, content []byte) error {
    finalFile, errOsCreate := os.Create(fileName)
    if errOsCreate != nil {
        return errOsCreate
    }
    defer finalFile.Close()

    if _, errIoWrite := finalFile.Write(content); errIoWrite != nil {
        return errIoWrite
    }
    return nil
}

func Option_b(bufSize []byte, file *os.File) error {
    var prefix string = "x"
    idx := []int{0, 0}

    fileInfo, errStat := file.Stat()
    if errStat != nil {
        return errStat
    }
    remainSize := fileInfo.Size()

    for {
        n, errIoRead := file.Read(bufSize)
        if errIoRead != nil && errIoRead != io.EOF {
            return errIoRead
        }
        if n == 0 {
            break
        }

        outputFile, errIoCreate := createOutputFile_b(prefix, idx)
        if errIoCreate != nil {
            return errIoCreate
        }
        defer outputFile.Close()

        if errIoWrite := writeToFile(outputFile, bufSize[:n]); errIoWrite != nil {
            return errIoWrite
        }

        remainSize -= int64(n)

        idx[0]++
        if idx[0] == 26 {
            idx[0] = 0
            idx[1]++
        }

        if idx[0] == 25 && idx[1] == 25 && remainSize > 0 {
			finalFileName := "xzz" // 最終ファイル名

			lastSize := make([]byte, remainSize)
			n, errIoRead := file.Read(lastSize)
			if errIoRead != nil && errIoRead != io.EOF {
				return errIoRead
			}
			if n == 0 {
				break
			}
			if errFinalWrite := writeFinalFile(finalFileName, lastSize[:n]); errFinalWrite != nil {
				return errFinalWrite
			}
		}
	}
	return nil
}