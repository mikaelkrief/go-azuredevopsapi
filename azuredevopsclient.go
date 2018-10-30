package azuredevopsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Client for manage azure devops organization
type Client struct {
}

type ProjectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type WorkItemResponse struct {
	ID     int `json:"id"`
	Rev    int `json:"rev"`
	Fields struct {
		SystemAreaPath                     string    `json:"System.AreaPath"`
		SystemTeamProject                  string    `json:"System.TeamProject"`
		SystemIterationPath                string    `json:"System.IterationPath"`
		SystemWorkItemType                 string    `json:"System.WorkItemType"`
		SystemState                        string    `json:"System.State"`
		SystemReason                       string    `json:"System.Reason"`
		SystemCreatedDate                  time.Time `json:"System.CreatedDate"`
		SystemCreatedBy                    string    `json:"System.CreatedBy"`
		SystemChangedDate                  time.Time `json:"System.ChangedDate"`
		SystemChangedBy                    string    `json:"System.ChangedBy"`
		SystemCommentCount                 int       `json:"System.CommentCount"`
		SystemTitle                        string    `json:"System.Title"`
		MicrosoftVSTSCommonStateChangeDate time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonPriority        int       `json:"Microsoft.VSTS.Common.Priority"`
	} `json:"fields"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemHistory struct {
			Href string `json:"href"`
		} `json:"workItemHistory"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}

func CreateWorkItem(pat string, organization string, projectName string, title string) WorkItemResponse {

	var jsonRequest = "[ { \"op\": \"add\", \"path\": \"/fields/System.Title\", \"from\": null,\"value\": \"" + title + "\"}]"

	var jsonStr = []byte(jsonRequest)

	var baseURL = "https://dev.azure.com/" + organization + "/" + projectName + "/_apis/wit/workitems/$Task?api-version=4.1"

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonStr))

	basic := "Basic " + pat

	req.Header.Set("Authorization", basic)
	req.Header.Set("Content-Type", "application/json-patch+json")

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

	data := WorkItemResponse{}
	json.Unmarshal([]byte(responseData), &data)

	return data

}

func CreateProject(pat string, organization string, projectName string) ProjectResponse {

	var jsonFormat = "{ \"name\": \"" + projectName + "\", \"description\": \"Frabrikam travel app for Windows Phone\", \"capabilities\": { \"versioncontrol\": { \"sourceControlType\": \"Git\"}, \"processTemplate\": {  \"templateTypeId\": \"6b724908-ef14-45cf-84f8-768b5384da45\" }}}"

	var jsonStr = []byte(jsonFormat)

	var baseUrl = "https://dev.azure.com/" + organization + "/_apis/projects?api-version=4.1"

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

	data := ProjectResponse{}
	json.Unmarshal([]byte(responseData), &data)

	return data

}
