package motuavb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type FloatValue struct {
	Value float32 `json:"value"`
}

func (c *Client) SendFloat32(request string, value float64) {

	f := fmt.Sprint(value)

	newR := `json={"` + request + `":` + f + `}`
	fmt.Println(newR)
	body := strings.NewReader(newR)
	req, err := http.NewRequest("POST", "http://"+c.ip+"/datastore", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func (c *Client) GetFloat32Value(request string) float32 {

	//I HATE MYSELF.

	resp, err := http.Get("http://" + c.ip + "/datastore/" + request)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var r FloatValue
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		// handle err
	}

	return r.Value

}

func (c *Client) GetMeters(mixer string) []float64 {

	//I HATE YOU ASWELL.

	resp, err := http.Get("http://" + c.ip + "/meters?meters=mix/level")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var m interface{}
	err = json.Unmarshal(bytes, &m)
	ma := m.(map[string]interface{})

	var ms []float64
	for _, e := range ma[mixer].([]interface{}) {
		ms = append(ms, e.(float64))
	}
	return ms
}
