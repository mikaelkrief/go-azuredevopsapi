package azuredevopsapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Client for manage azure devops organization
type Client struct {
}

func createProject(pat string, project string) string {

	var jsonFormat = "{ \"name\": \"alextestBuenas\", \"description\": \"Frabrikam travel app for Windows Phone\", \"capabilities\": { \"versioncontrol\": { \"sourceControlType\": \"Git\"}, \"processTemplate\": {  \"templateTypeId\": \"6b724908-ef14-45cf-84f8-768b5384da45\" }}}"

	var jsonStr = []byte(jsonFormat)

	var baseUrl = "https://dev.azure.com/" + project + "/_apis/projects?api-version=4.1"

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonStr))

	basic := "Basic " + pat

	req.Header.Set("Authorization", basic)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

}
