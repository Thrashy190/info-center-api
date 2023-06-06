package connection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Thrashy190/info-center-api/pkg/utils"
)

type Data struct {
	Name string `json:"name"`
}

func pullData(path string) []Data {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening seed file")
		fmt.Println(err)
	}
	utils.Success("File opened successfully")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data []Data

	err = json.Unmarshal(byteValue, &data)

	if err != nil {
		fmt.Printf("Error")
	}

	return data
}

func GenerateSeed() {
	pullData("pkg/connection/data/careers.json")
	pullData("pkg/connection/data/departments.json")
}
