package utils

import (
	"fmt"
	"os"
)

// GetRequiredEnv は、キーで指定された環境変数の値を取得します。
// 環境変数が設定されていない場合は、エラーを返します。
//
// パラメータ:
//   - key: 取得する環境変数名。
//
// 戻り値:
//   - 環境変数の値。
//   - 環境変数が設定されていない場合はエラー。
func GetRequiredEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable %s is required but not set", key)
	}
	return value, nil
}
