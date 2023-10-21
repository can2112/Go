// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Sizer interface{
// 	area() float64
// }

// type  Circle struct{
// 	Radius float64
// }
//  func (c Circle) area() float64 {
// 	return c.Radius * c.Radius* math.Pi
//  }
// type  Square struct{
// 	Side float64
// }
//  func (c Square) area() float64 {
// 	return c.Side * c.Side
//  }

//  func Less(a Sizer,b Sizer) Sizer{
// 	if(a.area()<b.area()){
// 		return a
// 	}
// 	return b
//  }

//  func main( ){
// fmt.Println("hello world!")
// c1,c2:=Circle{Radius:2.5},Circle{Radius:5.0}
// sq1,sq2:=Square{Side:3.0},Square{Side:4.0}
// fmt.Println(Less(c1,c2))
// fmt.Println(Less(sq1,sq2))
//  }

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type MatchService interface {
	Add(a int, b int) int
	Mul(a int, b int) int
}

type matchsvc struct{}

func (m matchsvc) Add(a, b float64) float64 {
	return a + b
}
func (m matchsvc) Mul(a, b int) int {
	return a * b
}

// type EndPoint func(ctx context.Context, request interface{})(response interface{},err error)

type AddReq struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}
type AddRes struct {
	Response float64 `json:"response"`
}

func makeAddEndPoint(svc matchsvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddReq)
		res := svc.Add(req.A, req.B)
		return AddRes{
			Response: res,
		}, nil

	}

}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := AddReq{}
	err := json.NewDecoder(r.Body).Decode((&request))
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode((response))
}

func makeLogginMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			startTime := time.Now()
			fmt.Println("Request started")
			defer fmt.Println("Requesst ended, time:" + time.Since(startTime).String())
			return next(ctx, request)

		}
	}
}
func main1() {
	svc := matchsvc{}
	addEndPoint := makeAddEndPoint(svc)
	logEndPoint := makeLogginMiddleware()(addEndPoint)

	upperCaseHandler := httptransport.NewServer(logEndPoint, decodeAddRequest, encodeResponse)

	http.Handle("/add", upperCaseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
