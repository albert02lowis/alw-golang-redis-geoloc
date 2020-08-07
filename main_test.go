package main

import (
	"net/http"
	"testing"

	"gopkg.in/gavv/httpexpect.v1"
)

func TestPutDrivers(t *testing.T) {
	e := fastHTTPTester(t)

	PutDriversNormalCases(t, e)
}

func PutDriversNormalCases(t *testing.T, e *httpexpect.Expect) {

	jsonbody := map[string]interface{}{
		userLatitudeKey:  12.97161923,
		userLongitudeKey: 77.59463452,
		userAccuracyKey:  accuracyMinimum,
	}

	e.PUT("/drivers/12/location").
		WithJSON(jsonbody).
		Expect().
		Status(http.StatusOK).Body().Equal(emptyBody)

	jsonbody = map[string]interface{}{
		userLatitudeKey:  80.97161923,
		userLongitudeKey: 3.59463452,
		userAccuracyKey:  accuracyMinimum,
	}

	e.PUT("/drivers/13/location").
		WithJSON(jsonbody).
		Expect().
		Status(http.StatusOK).Body().Equal(emptyBody)

	jsonbody = map[string]interface{}{
		userLatitudeKey:  55.0,
		userLongitudeKey: 147.0,
		userAccuracyKey:  accuracyMinimum,
	}

	e.PUT("/drivers/151/location").
		WithJSON(jsonbody).
		Expect().
		Status(http.StatusOK).Body().Equal(emptyBody)
}

// fastHTTPTester returns a new Expect instance to test FastHTTPHandler().
func fastHTTPTester(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		// Pass requests directly to FastHTTPHandler.
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(CreateHandlers()),
			Jar:       httpexpect.NewJar(),
		},
		// Report errors using testify.
		Reporter: httpexpect.NewAssertReporter(t),
	})
}
