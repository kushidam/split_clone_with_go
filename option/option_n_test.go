package option

import (
	"os"
	"testing"
)

func TestValidateArgs(t *testing.T) {
    // テストデータ
    okValidArgs := []string{"1", "2", "3"}
    ngValidArgs := []string{"1", "abc", "3"}

    // 正常な場合
    err := validateArgs(okValidArgs)
    if err != nil {
        t.Errorf("validateArgs ERROR: %v", err)
    }

    // 異常な場合
    err = validateArgs(ngValidArgs)
    if err == nil {
        t.Errorf("validateArgs ERROR: Expect: err, Result:nil")
    }
}
func TestOptionNDefault(t *testing.T) {
	// テストデータ
	okValidSplitNum := 5
	okValidSplitNumLow := 1
	okValidSplitNumHigh := 676
	ngValidSplitNumLow := 0
	ngValidSplitNumHigh := 677

	// テストケースごとにテスト実行
    tmpFile, err := os.CreateTemp("", "testInputTmp")
    if err != nil {
        t.Fatalf("CreateTemp ERROR: %v", err)
    }
    defer os.Remove(tmpFile.Name())
    defer tmpFile.Close()

    err = option_n_default(okValidSplitNum, tmpFile)
    if err != nil {
        t.Errorf("option_n_default ERROR: %v", err)
    }

    err = option_n_default(okValidSplitNumLow, tmpFile)
    if err != nil {
        t.Errorf("option_n_default ERROR: %v", err)
    }

    err = option_n_default(okValidSplitNumHigh, tmpFile)
    if err != nil {
        t.Errorf("option_n_default ERROR: %v", err)
    }

    err = option_n_default(ngValidSplitNumLow, tmpFile)
    if err == nil {
        t.Errorf("option_n_default ERROR: Expect: err, Result:nil")
    }

    err = option_n_default(ngValidSplitNumHigh, tmpFile)
    if err == nil {
        t.Errorf("option_n_default ERROR: Expect: err, Result:nil")
    }
}
