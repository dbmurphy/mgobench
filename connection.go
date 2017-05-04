package mgobench

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"
)

// IsBlankString return if string is only space or empty / zero length
func IsBlankString(s string) bool {
	return strings.TrimSpace(s) == ""
}

type CollectionBindFunc func(s *mgo.Session) (*mgo.Collection, error)

func NewCollectionBindFunc(db string, coll string) CollectionBindFunc {
	fmt.Println(db, "                 ", coll)
	if IsBlankString(db) || IsBlankString(coll) {
		return nil
	}
	return func(s *mgo.Session) (*mgo.Collection, error) {
		if s == nil {
			return nil, errors.New("nil session")
		}
		sc := s.Copy()
		// fmt.Printf("Database ---- %s \n Collection ---- %s \n", db, coll)
		return sc.DB(db).C(coll), nil
	}
}

type MgoManager struct {
	Session *mgo.Session
	CFn     CollectionBindFunc
}

func (mc MgoManager) Coll() (*mgo.Collection, error) {
	if mc.Session == nil {
		panic("nil mgo session")
	}
	if mc.CFn == nil {
		return nil, errors.New("CollectionBindFunc 'CFn' is nil")
	}

	return mc.CFn(mc.Session)
}

func NewMgoManagerWithDefaultBinder(s *mgo.Session, db string, coll string) *MgoManager {
	cfn := NewCollectionBindFunc(db, coll)
	if cfn == nil {
		return nil
	}
	return &MgoManager{
		Session: s,
		CFn:     cfn,
	}
}
