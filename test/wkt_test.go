package test

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/spatial-go/geoos/geoencoding"
	"github.com/spatial-go/geoos/space"
)

func TestWktToGeom(t *testing.T) {
	var polygonWktStr = "POLYGON ((30 10, 10 20, 20 40, 40 40, 30 10))"
	buf := new(bytes.Buffer)
	buf.Write([]byte(polygonWktStr))
	got, err := geoencoding.Read(buf, geoencoding.WKT)
	if err != nil {
		log.Println(err)
	}
	log.Println(got)
}
func TestGeomToWkt(t *testing.T) {
	multiPoint := space.MultiPoint{{-1, 0}, {-1, 2}, {-1, 3}, {-1, 4}, {-1, 7}, {0, 1}, {0, 3}}
	buf := new(bytes.Buffer)
	err := geoencoding.Write(buf, multiPoint, geoencoding.WKT)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(buf.String())
}
