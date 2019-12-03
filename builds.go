package teamcity_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c *Client) SearchBuilds(locators map[string]string, count int) (b BuildList, e error) {
	path := "/builds?"
	for k, v := range locators {
		path = path + fmt.Sprintf("%s:%s,", k, v)
	}
	js := &BuildList{}
	res, err := c.doRequest(fmt.Sprintf("%s&count=%d", path, count))
	if err != nil {
		return BuildList{}, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, js)
	return *js, nil
}

type BuildList struct {
	Count  int     `json:"count"`
	HRef   string  `json:"href"`
	Builds []Build `json:"build"`
}

type Build struct {
	ID          int64  `json:"id"`
	BuildTypeId string `json:"buildTypeId"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	State       string `json:"state"`
	BranchName  string `json:"branchName"`
	HRef        string `json:"href"`
	WebUrl      string `json:"webUrl"`
}