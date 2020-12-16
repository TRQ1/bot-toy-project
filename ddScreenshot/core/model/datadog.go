package model

type Key struct {
	Dd_api		string
	Dd_app_api	string
}

type Graphs struct {
	Start      	int
	End		   	int
	Metric_name	string
	Query		string
}
