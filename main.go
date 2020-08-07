package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const userIDKey string = "id"

var emptyBodyRaw = json.RawMessage(`{}`)

func main() {
	defer MyLogFile.Close()
	log.Println("Listening at :8080")
	ServeMain(":8080")
}

//ServeMain starts the web server on the specified address
func ServeMain(addr string) {
	log.Println(fasthttp.ListenAndServe(addr, CreateHandlers()))
}

//CreateHandlers creates and return the necessary request handlers
func CreateHandlers() fasthttp.RequestHandler {
	myrouter := fasthttprouter.New()
	myrouter.PUT("/entity/:"+userIDKey, StoreEntity)

	return myrouter.Handler
}

//StoreEntity is a function to update location data
func StoreEntity(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(emptyBodyRaw)
	reply := ""
	defer fmt.Fprintf(ctx, reply)

	//check request body parsing
	rbody := &StoreEntityBody{}
	err := json.Unmarshal(ctx.Request.Body(), rbody)
	if err != nil {
		GenerateInvalidErrorBody(ctx, err, fasthttp.StatusUnprocessableEntity)
		return
	}

	//check latitude and longitude
	if CheckLatLonWithResponse(ctx, rbody.Latitude, rbody.Longitude) != nil {
		return
	}

	//check if entityID is valid
	_, err = ParseIntWithResponse(ctx, userIDKey)
	if err != nil {
		return
	}

	err = RecordEntityLocation(ctx.UserValue(userIDKey).(string), rbody.Latitude, rbody.Longitude)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
}

//ParseIntWithResponse checks validity of int strings and generate response accordingly
func ParseIntWithResponse(ctx *fasthttp.RequestCtx, key string) (int, error) {
	val, err := strconv.Atoi(ctx.UserValue(key).(string))
	if err != nil {
		GenerateInvalidErrorBody(ctx, &InvalidInputError{errorUnableParseValue + ": " + key}, fasthttp.StatusUnprocessableEntity)
		return 0, err
	}
	return val, err
}

//CheckLatLonWithResponse checks latitude and longitude validity and generate response accordingly
func CheckLatLonWithResponse(ctx *fasthttp.RequestCtx, lat, lon float64) error {
	err := CheckLatLon(lat, lon)
	if err != nil {
		GenerateInvalidErrorBody(ctx, err, fasthttp.StatusBadRequest)
	}
	return err
}

//GenerateInvalidErrorBody creates JSON response to report any errors
func GenerateInvalidErrorBody(ctx *fasthttp.RequestCtx, uerr error, statusCode int) {
	ctx.SetStatusCode(statusCode)
	b, merr := json.Marshal(uerr)
	if merr == nil {
		ctx.SetContentType("application/json")
		ctx.SetBody(b)
	}
}

func (e *InvalidInputError) Error() string {
	return e.Errors
}
