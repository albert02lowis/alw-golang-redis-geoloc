package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//RecordEntityLocation records the latest location of the entity with id to the datastore
func RecordEntityLocation(id string, lat float64, lon float64) error {
	err := client.GeoAdd(locationKey, &redis.GeoLocation{Longitude: lon, Latitude: lat, Name: id}).Err()
	return err
}

//CheckLatLon checks if latitude and longitude values falls in the correct range and returns error otherwise
func CheckLatLon(lat float64, lon float64) error {
	if lat < latitudeMinimum || lat > latitudeMaximum {
		return &InvalidEntityStoreInput{fmt.Sprintf(errorlatitudeExceed+" %v", latitudeMaximum)}
	}
	if lon < longitudeMinimum || lon > longitudeMaximum {
		return &InvalidEntityStoreInput{fmt.Sprintf(errorlongitudeExceed+" %v", longitudeMaximum)}
	}
	return nil
}

func (e *InvalidEntityStoreInput) Error() string {
	return e.Errors
}
