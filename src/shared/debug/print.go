package debug

import (
	"encoding/json"
	"fmt"
)

func PrintStr(v string) {
	fmt.Println("===================")
	fmt.Println(v)
	fmt.Println("===================")
}

func PrintMap(data map[string]interface{}) {
	// マップを整形されたJSONに変換
	prettyData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// JSONを出力
	fmt.Println(string(prettyData))
}
