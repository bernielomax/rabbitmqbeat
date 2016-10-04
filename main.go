package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/metricbeat/beater"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/bernielomax/rabbitmqbeat/include"
	// Uncomment the following line to include all official metricbeat module and metricsets
	//_ "github.com/elastic/beats/metricbeat/include"
)

var Name = "rabbitmqbeat"

func main() {
	if err := beat.Run(Name, "", beater.New); err != nil {
		os.Exit(1)
	}
}
