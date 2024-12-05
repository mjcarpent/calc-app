package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestFixture(t *testing.T) {
	gunit.Run(new(HTTPTestFixture), t)
}

type HTTPTestFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.

	// Declare useful state here (probably the stuff being tested, any fakes, etc...).
}

func (this *HTTPTestFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method (because it starts with "Setup")
}
func (this *HTTPTestFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}

func (this *HTTPTestFixture) TestAdd() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/add?param1=3&param2=4", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusOK)
	this.So(response.Body.String(), should.Equal, "7")
}

func (this *HTTPTestFixture) TestSubtract() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/sub?param1=3&param2=4", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusOK)
	this.So(response.Body.String(), should.Equal, "-1")
}

func (this *HTTPTestFixture) TestMultiply() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/mul?param1=3&param2=4", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusOK)
	this.So(response.Body.String(), should.Equal, "12")
}

func (this *HTTPTestFixture) TestDivide() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/div?param1=17&param2=3", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusOK)
	this.So(response.Body.String(), should.Equal, "5")
}

func (this *HTTPTestFixture) TestMod() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/mod?param1=17&param2=3", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusOK)
	this.So(response.Body.String(), should.Equal, "2")
}

func (this *HTTPTestFixture) TestBadParam1() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/add?param1=NaN&param2=4", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusBadRequest)
	this.So(response.Body.String(), should.Equal, "The given operand must be an integer")
}

func (this *HTTPTestFixture) TestBadParam2() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/add?param1=1&param2=NaN", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusBadRequest)
	this.So(response.Body.String(), should.Equal, "The given operand must be an integer")
}

func (this *HTTPTestFixture) TestURLJunk() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/asdf", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusNotFound)
	this.So(response.Body.String(), should.Equal, "404 page not found\n")
}

func (this *HTTPTestFixture) TestBadVerb() {

	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/add?param1=3&param2=4", nil)

	router.ServeHTTP(response, request)

	this.So(response.Code, should.Equal, http.StatusMethodNotAllowed)
	this.So(response.Body.String(), should.Equal, "Method Not Allowed\n")
}
