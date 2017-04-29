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
type Influxdb struct {
	conn client.Client
}

func (client *Influxdb) InsertData(measurement string, tag string, field float64) {
	tags := map[string]string{
		"mongodb": tag,
	}
	fields := map[string]interface{}{
		"count": field,
	}
	createMetrics(client.conn, measurement, tags, fields)
}

func NewInfluxClient() *Influxdb {

	c, err := client.NewUDPClient(client.UDPConfig{Addr: "localhost:8089"})
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return &Influxdb{conn: c}
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
