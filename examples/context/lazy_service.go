package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const requestIDKey = "rid"

//func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
//	reqID := req.Header.Get("X-Request-ID")
//	if reqID == "" {
//		reqID = "0"
//	}
//	return context.WithValue(ctx, requestIDKey, reqID)
//}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqID := req.Header.Get("X-Request-ID")
		if reqID == "" {
			rand.Seed(time.Now().UnixNano())
			reqID = strconv.Itoa(rand.Intn(2))
		}
		ctx := context.WithValue(req.Context(), requestIDKey, reqID)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	reqID := req.Context().Value(requestIDKey).(string)
	ranNum, _ := strconv.Atoi(reqID)
	if ranNum == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "slow response, %d\n", ranNum)
		fmt.Printf("slow response, %d\n", ranNum)
		return
	}
	fmt.Fprintf(w, "quick response, %d\n", ranNum)
	fmt.Printf("quick response, %d\n", ranNum)
	return
}

func main() {
	lh := http.HandlerFunc(lazyHandler)
	http.Handle("/", middleware(lh))
	http.ListenAndServe(":9200", nil)
}
