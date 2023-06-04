package connection

import (
	"encoding/json"
	"fmt"
	"github.com/Thrashy190/info-center-api/pkg/utils"
	"io/ioutil"
	"os"
)

func OpenJson() {
	jsonFile, err := os.Open("pkg/connection/data/careers.json")

	if err != nil {
		fmt.Println("Error opening seed file")
		fmt.Println(err)
	}
	utils.Success("File opened successfully")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var careers map[string]interface{}

	err = json.Unmarshal(byteValue, &careers)

	fmt.Println(careers["careers"])

}

func GenerateSeed() {
	OpenJson()
}
