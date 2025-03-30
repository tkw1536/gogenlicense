package gogenlicense

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/google/go-licenses/v2/licenses"
	"go.opencensus.io/plugin/ochttp"
)

// This entire file is a huge hack and shouldn't be needed.
// Please bump https://github.com/google/go-licenses/issues/256.

var libraryFunc = reflect.ValueOf((*licenses.Library).FileURL)
var sourceClientTyp = libraryFunc.Type().In(2).Elem() // reflect.TypeFor[source.Client]

// like (*licenses.Library).FileURL, but not taking an internal parameter
func goLicensesLibraryFileURL(lib *licenses.Library, context context.Context, timeout time.Duration, file string) (string, error) {
	result := libraryFunc.Call([]reflect.Value{
		reflect.ValueOf(lib),
		reflect.ValueOf(context),
		newSourceClient(timeout),
		reflect.ValueOf(file),
	})

	str := result[0].String()
	err, _ := result[1].Interface().(error)
	return str, err
}

// like source.NewClient, but not internal
func newSourceClient(timeout time.Duration) reflect.Value {
	client := &http.Client{
		Transport: &ochttp.Transport{},
		Timeout:   timeout,
	}

	sourceClient := reflect.New(sourceClientTyp)

	// get the httpClient field
	// and forget that it was retrieved by means of an unexported field
	httpClientField := sourceClient.Elem().FieldByName("httpClient")
	httpClientField = reflect.NewAt(httpClientField.Type(), httpClientField.Addr().UnsafePointer()).Elem()
	httpClientField.Set(reflect.ValueOf(client))

	return sourceClient
}
