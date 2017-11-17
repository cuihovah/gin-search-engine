package v1

import "github.com/gin-gonic/gin"
import "net/http"
import (
	"../utils"
    "../../schema"
	"github.com/huichen/sego"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sort"
	"strconv"
)

type Query schema.Query
type String []string
type Querys []Query

func (qs *Querys) intersection() String {

	nr := len(*qs)
	wc := make(map[string]int)
	for _, v := range *qs {
		index := v.Index
		for _, value := range index {

			if _, ok := wc[value]; !ok {
				wc[value] = 0
			}

			wc[value] = wc[value] + 1
		}
	}

	result := []string{}
	for key, value := range wc {
		if value >= nr {
			result = append(result, key)
		}
	}

	return result
}

func GetOne(ctx *gin.Context) {

	qs := ctx.Query("q")
	_offset := ctx.DefaultQuery("offset", "0")
	_limit := ctx.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(_offset)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(_limit)
	if err != nil {
		limit = 10
	}

	c := utils.DBSession.DB("search_engine").C("index")
	d := utils.DBSession.DB("search_engine").C("document")

	query := Querys{}

	segments := utils.Segmenter.Segment([]byte(qs))
	output := sego.SegmentsToSlice(segments, false)

	err = c.Find(bson.M{"token": bson.M{"$in": output}}).All(&query)

	if err != nil {
		log.Fatal(err)
	}

	h := query.intersection()

	sort.Strings(h)

	total := len(h)
	max := total

	if offset+limit < max {
		max = offset + limit
	}

	if offset > max {
		h = []string{}
	} else {
		h = h[offset:max]
	}

	_ids := []bson.ObjectId{}

	for _, v := range h {
		_ids = append(_ids, bson.ObjectIdHex(v))
	}

	result := []schema.Document{}

	err = d.Find(bson.M{"_id": bson.M{"$in": _ids}}).All(&result)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": gin.H{
			"total_number": total,
			"items":        result,
		},
	})
}
