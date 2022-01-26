package main

import (
	"context"
	"fmt"
	"go-grpc/internal/rpc"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:6565"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := rpc.NewMovieServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("===============================FETCH MOVIES===============================================")

	// This is for fetching movies with searchword Batman from page 1-5
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			result, err := c.FetchArticle(ctx, &rpc.FetchRequest{
				Searchword: "Batman",
				Page:       "3",
			})
			if err != nil {
				log.Fatalf("could not getch movie: %v", err)
			}
			log.Printf("List Movie page %v detail: %v", i, result.Movies)
		}()
	}
	wg.Wait()

	fmt.Println("=================================SINGLE MOVIES=============================================")

	// This is for Get Single Movies with IMDBID tt2975590
	result, err := c.GetArticle(ctx, &rpc.SingleRequest{
		Id: "tt2975590",
	})
	if err != nil {
		log.Fatalf("could not create sample: %v", err)
	}
	log.Printf("user detail: %v", result)

}
