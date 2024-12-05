package handler

import (
	"net/http"
	"strconv"

	"github.com/mjcarpent/calc-lib"
)

func NewHTTPRouter() http.Handler {

	router := http.NewServeMux()
	router.Handle("GET /add", NewHTTPHandler(calc.Addition{}))
	router.Handle("GET /sub", NewHTTPHandler(calc.Subtraction{}))
	router.Handle("GET /mul", NewHTTPHandler(calc.Multiplication{}))
	router.Handle("GET /div", NewHTTPHandler(calc.Division{}))
	router.Handle("GET /mod", NewHTTPHandler(calc.Modulus{}))

	return router
}

type HTTPHandler struct {
	calculator Calculator
}

func NewHTTPHandler(calculator Calculator) http.Handler {
	return &HTTPHandler{calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	param1, err := strconv.Atoi(req.URL.Query().Get("param1"))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(operandMustBeInt.Error()))
		return
	}

	param2, err := strconv.Atoi(req.URL.Query().Get("param2"))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(operandMustBeInt.Error()))
		return
	}

	answer := this.calculator.Calculate(param1, param2)

	//endpoint := string(req.URL.Path[1:])
	//opSymbol := operators[endpoint]

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(strconv.Itoa(answer)))
	//_, err := res.Write([]byte(strconv.Itoa(answer)))
	//if err != nil {
	//	log.Fatalf("There's an error with the server")
	//}
}
