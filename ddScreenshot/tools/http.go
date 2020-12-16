package tools

import (
	"bytes"
	"io/ioutil"
	"encodig/json"
	"log"
	"net/http"
	"time"

	model "github.com/TRQ1/bot-toy-project/ddScreenshot/core/model"
)

func curlDatadog(dd model.Graphs, ky model.Key) string {

	req, err := http.NewRequest("GET", "https://api.datadoghq.com/api/v1/graph/snapshot", nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("DD-API-KEY", ky.Dd_api)
	req.Header.Add("DD-APPLICATION-KEY", ky.Dd_app_api)

	q := req.URL.Query()
	q.Add("metric_name", dd.Metric_name)
	q.Add("query", dd.Query)
	q.Add("start", strconv.Itoa(dd.Start))
	q.Add("end", strconv.Itoa(dd.End))
	
	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s\n", body)
	
	return body.snapshot_url

}
