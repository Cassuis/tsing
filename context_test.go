package tsing

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGET(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	app := New()
	app.Event.Handler = func(event Event) {
		log.Println(event.Message)
	}
	app.Router.GET("/", func(ctx Context) error {
		t.Error(ctx.QueryValue("id").Int())
		return nil
	})

	r, err := http.NewRequest("GET", "/?id=abc", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}

func TestPOST(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	app := New()
	app.Event.Handler = func(event Event) {
		log.Println(event.Message)
	}
	app.Router.POST("/", func(ctx Context) error {
		t.Error(ctx.FormValue("id").Int())
		return nil
	})

	v := url.Values{}
	v.Add("id", "abc")
	r, err := http.NewRequest("POST", "/", strings.NewReader(v.Encode()))
	if err != nil {
		t.Error(err.Error())
		return
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(httptest.NewRecorder(), r)
}

func TestRoute(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	app := New()
	app.Event.Handler = func(event Event) {
		log.Println(event.Message)
	}
	app.Router.GET("/:classID/:id", func(ctx Context) error {
		t.Error(ctx.RouteValue("classID").Int())
		t.Error(ctx.RouteValue("id").Int())
		return nil
	})

	r, err := http.NewRequest("GET", "/1/2", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}
