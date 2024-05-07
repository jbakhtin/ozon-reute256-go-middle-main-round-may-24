package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type JSON any

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanLines)

	in.Scan()
	t, _ := strconv.Atoi(in.Text())

	var myJSONList []JSON

	for i := 0; i < t; i++ {

		//if flag {
		//	out.WriteString(",")
		//}

		in.Scan()
		rowsCount, _ := strconv.Atoi(in.Text())
		var rawJSON strings.Builder

		var myJSON JSON

		for j := 0; j < rowsCount; j++ {
			in.Scan()
			rawJSON.WriteString(in.Text())
		}

		err := json.Unmarshal([]byte(rawJSON.String()), &myJSON)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		myPreparedJSON := prettyfy(myJSON)

		myJSONList = append(myJSONList, myPreparedJSON)

		//flag = true
	}

	bytes, err := json.Marshal(myJSONList)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Fprint(out, string(bytes))
}

func prettyfy(myJSON JSON) JSON {

	switch myConvertedJSON := myJSON.(type) {
	case nil:
		return nil
	case string:
		return myConvertedJSON
	case map[string]any:
		if len(myConvertedJSON) == 0 {
			return nil
		}

		for key, value := range myConvertedJSON {
			prettyJSON := prettyfy(value)
			if prettyJSON == nil {
				delete(myConvertedJSON, key)
			}
		}

		if len(myConvertedJSON) == 0 {
			return nil
		}

		return myConvertedJSON
	case []any:
		for i := len(myConvertedJSON) - 1; i >= 0; i-- {
			prettyJSON := prettyfy(myConvertedJSON[i])
			if prettyJSON == nil {
				myConvertedJSON = append(myConvertedJSON[:i], myConvertedJSON[i+1:]...)
			}
		}

		if len(myConvertedJSON) == 0 {
			return nil
		}

		return myConvertedJSON
	case map[int]any:
		if len(myConvertedJSON) == 0 {
			return nil
		}

		for key, value := range myConvertedJSON {
			prettyJSON := prettyfy(value)
			if prettyJSON == nil {
				delete(myConvertedJSON, key)
			}
		}

		if len(myConvertedJSON) == 0 {
			return nil
		}

		return myConvertedJSON
	}

	return nil
}
