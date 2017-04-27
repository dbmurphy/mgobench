package mgobench

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	database = "mongo"
	username = ""
	password = ""
)

type tags map[string]string
type fields map[string]interface{}

func InsertData(measurement string, tag string, field float64) {
	tags := map[string]string{
		"mongodb": tag,
	}
	fields := map[string]interface{}{
		"count": field,
	}
	c := influxDBClient()
	createMetrics(c, measurement, tags, fields)
}

func influxDBClient() client.Client {

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

func createMetrics(c client.Client, measurement string, tags tags, fields fields) {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	point, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp.AddPoint(point)

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
}
