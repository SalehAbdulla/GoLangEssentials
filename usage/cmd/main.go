package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 22
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", val)
	fmt.Println("Took:", time.Since(start))
}


type Response struct {
	value int
	err error
}


func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond * 200)
	defer cancel()
	
	responseChan := make(chan Response)
	
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		responseChan <- Response{
			value: val,
			err: err,
		}
	}()

	for {
		select{
		case <- ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took too long")
		case <- responseChan:
			return responseChan.val, responseChan.err

		}

	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 666, nil
}
