package helpers

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ochom/focus/domain"
)

//Search findRxcuiByString for more information:
//see https://rxnav.nlm.nih.gov/api-RxNorm.findRxcuiByString.html
func Search(param string) {
	fmt.Printf("Searching: %v...\n", param)
	payload := domain.HttpRequestParam{
		Method: "GET",
		URL:    domain.BaseRoute + fmt.Sprintf("/REST/rxcui.json?name=%s&search=1", url.QueryEscape(param)),
		Headers: map[string]string{
			"Authorization": "Bearer ",
		},
	}

	resp, err := MakeHttpRequest(payload)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	data := make(map[string]interface{})

	err = json.Unmarshal(resp, &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}

	//pretty print data
	m, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	fmt.Printf("%v\n", string(m))
}

//Describe getNDCProperties for more information
//see https://rxnav.nlm.nih.gov/api-RxNorm.getNDCProperties.html
func Describe(param string) {
	fmt.Printf("Describing: %v...\n", param)
	payload := domain.HttpRequestParam{
		Method: "GET",
		URL:    domain.BaseRoute + fmt.Sprintf("/REST/ndcproperties.json?id=%s", param),
		Headers: map[string]string{
			"Authorization": "Bearer ",
		},
	}

	resp, err := MakeHttpRequest(payload)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	data := make(map[string]interface{})

	err = json.Unmarshal(resp, &data)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}

	//pretty print data
	m, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	fmt.Printf("%v\n", string(m))
}
