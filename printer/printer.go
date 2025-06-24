package printer

import (
	"encoding/json"
	"fmt"
)

func PrintInJSONFormat(data any) {
	dataJson, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(dataJson))
}
