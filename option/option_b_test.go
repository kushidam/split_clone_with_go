package option

import (
	"os"
	"testing"
)

func TestCreateOutputFile_b(t *testing.T) {
	// テストデータ
	prefix := "x"
	idx := []int{0, 0}
	expectFileName := "xaa"

	// 関数を呼び出し、出力されるファイル名を検証
	outputFile, err := createOutputFile_b(prefix, idx)
	if err != nil {
		t.Fatalf("createOutputFile ERROR: %v", err)
	}
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	outputFileName := outputFile.Name()
	if outputFileName != expectFileName {
		t.Errorf("\nファイル名不一致\nExpect: %s, Result: %s",
			expectFileName, outputFileName)
	}
	
	idx = []int{25, 25}
	expectFileName = "xzz"

	// 関数を呼び出し、出力されるファイル名を検証
	outputFile, err = createOutputFile_b(prefix, idx)
	if err != nil {
		t.Fatalf("createOutputFile ERROR: %v", err)
	}
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	outputFileName = outputFile.Name()
	if outputFileName != expectFileName {
		t.Errorf("\nファイル名不一致\nExpect: %s, Result: %s",
			expectFileName, outputFileName)
	}
}