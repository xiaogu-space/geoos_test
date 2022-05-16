package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spatial-go/geoos/geojson"
	"github.com/spatial-go/geoos/planar"
)

func main() {
	jsonFile, err := os.Open("data/河北天津.geojson")
	if err != nil {
		fmt.Println("报错")
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	geoStr := string(byteValue)
	// fmt.Println(geoStr)

	fc, err := geojson.UnmarshalFeatureCollection([]byte(geoStr))
	f0 := fc.Features[0]
	f1 := fc.Features[1]
	geo0 := f0.Geometry.Coordinates
	geo1 := f1.Geometry.Coordinates

	G := planar.GEOAlgorithm{}

	time0 := time.Now()

	got, err := G.Union(geo0, geo1)

	data, _ := geojson.NewFeature(*geojson.NewGeometry(got)).MarshalJSON()
	// geoOut := string(data)

	// 将 JSON 格式数据写入当前目录下的d books 文件（文件不存在会自动创建）
	err = os.WriteFile("data/union.geojson", data, 0644)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(time0)
	fmt.Println("该函数执行完成耗时：", elapsed)

	// fmt.Println(geoStr)
}
