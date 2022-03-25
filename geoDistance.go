package main

import (
	"fmt"
	. "math"
	"strconv"
	"strings"

	"github.com/pierrre/geohash"
)

func main() {
	// Point 2
	location1 := "-8.5797245,116.0987025"
	locationSplited1 := strings.Split(location1, ",")
	latitude1, _ := strconv.ParseFloat(locationSplited1[0], 64)
	longitude1, _ := strconv.ParseFloat(locationSplited1[1], 64)
	fmt.Println(latitude1)
	fmt.Println(longitude1)

	// Point 2
	location2 := "-8.5876678,116.1175865"
	locationSplited2 := strings.Split(location2, ",")
	latitude2, _ := strconv.ParseFloat(locationSplited2[0], 64)
	longitude2, _ := strconv.ParseFloat(locationSplited2[1], 64)
	fmt.Println(latitude2)
	fmt.Println(longitude2)

	// Distance 2 Points
	distanceInMeter := 6371000 * Acos(Sin(latitude1*Pi/180)*Sin(latitude2*Pi/180)+Cos(latitude1*Pi/180)*Cos(latitude2*Pi/180)*Cos(longitude2*Pi/180-longitude1*Pi/180)) // Hasil dalam meter
	distanceInKiloMeter := float64(int(distanceInMeter*1)/100) / 10 // Hasil dalam kilometer
	fmt.Println("Distance is :", distanceInKiloMeter)

	// Encode
	geohashEncoded := geohash.EncodeAuto(latitude2, longitude2)
	fmt.Println("geohash encoded", geohashEncoded)

	// Decode
	geohashBox, _ := geohash.Decode(geohashEncoded)
	geohashDecoded := geohashBox.Round()
	latitudeDecoded := geohashDecoded.Lat
	longitudeDecoded := geohashDecoded.Lon
	fmt.Println("latitude decoded", latitudeDecoded)
	fmt.Println("longitude decoded", longitudeDecoded)
}
