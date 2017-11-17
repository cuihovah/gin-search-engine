package utils

import "github.com/huichen/sego"
import "../../config"

var Segmenter sego.Segmenter

func SegmenterInit() {
	Segmenter.LoadDictionary(config.DictPath)
}
