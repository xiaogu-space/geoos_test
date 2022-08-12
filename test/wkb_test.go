package test

import (
	"bytes"
	"log"
	"testing"

	"github.com/spatial-go/geoos/geoencoding"
)

func TestWkbToGeom(t *testing.T) {
	var wkbHex = `0101000020E610000000000020D8135D400000004072054440`
	buf := new(bytes.Buffer)
	buf.Write([]byte(wkbHex))
	got, err := geoencoding.Read(buf, geoencoding.WKB)
	if err != nil {
		log.Println(err)
	}
	log.Println(got.Geom())
}
