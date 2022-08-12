package test

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/geoencoding"
	"github.com/spatial-go/geoos/space"
)

func TestGeomToGeoJson(t *testing.T) {
	geom := space.MultiPoint{{-1, 0}, {-1, 2}, {-1, 3}, {-1, 4}, {-1, 7}, {0, 1}, {0, 3}}
	buf := new(bytes.Buffer)
	err := geoencoding.Write(buf, geom.Geom(), geoencoding.GeoJSON)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(buf.String())
}
func TestGeoJsonToGeom(t *testing.T) {
	var json = `{"type":"MultiPoint","coordinates":[[-1,0],[-1,2],[-1,3],[-1,4],[-1,7],[0,1],[0,3]]}`
	buf := new(bytes.Buffer)
	buf.Write([]byte(json))
	got, err := geoencoding.Read(buf, geoencoding.GeoJSON)
	if err != nil {
		log.Println(err)
	}
	log.Println(reflect.TypeOf(got))
}
