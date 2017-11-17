package utils

import (
	"gopkg.in/mgo.v2"
    "../../config"
)

var DBSession *mgo.Session

func DBInit() {
	DBSession, _ = mgo.Dial(config.DB)
	DBSession.SetMode(mgo.Monotonic, true)
}
