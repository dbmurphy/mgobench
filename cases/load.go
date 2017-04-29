package cases

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type EmptyDoc struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

type FlatT1Doc struct {
	ID    bson.ObjectId `bson:"_id,omitempty`
	StrF  string        `bson:"strf"`
	IntF  int64         `bson:"intf"`
	BoolF bool          `bson:"boolf"`
	TimeF time.Time     `bson:"timef"`
}

type FlatT1InsertTask struct {
}

// EmptyTest Func used to return empty data for test
func EmptyTest() (data *EmptyDoc) {
	return &EmptyDoc{}
}

func FlatT1DocTest() (data *FlatT1Doc) {
	return &FlatT1Doc{}
}
