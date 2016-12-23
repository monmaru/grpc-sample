package main

import (
	"flag"
	"log"
	"os"

	"fmt"

	pb "github.com/monmaru/github-tren-go/proto"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	addrFlag = flag.String("addr", "localhost:5000", "server address host:post")
)

func fetch(lang string) error {
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	defer conn.Close()
	c := pb.NewGithubTrendingClient(conn)

	resp, err := c.Fetch(context.Background(), &pb.Req{Lang: lang})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
		return err
	}
	log.Printf("Repositories: %s", resp.Repositories)
	return nil
}

func main() {
	usage := "Search for trends in the programming language. Ex: `-l Go`"
	var lang string
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       usage,
			Destination: &lang,
		},
	}

	app.Action = func(c *cli.Context) error {
		if lang == "" {
			fmt.Println(usage)
			return nil
		}
		return fetch(lang)
	}

	app.Run(os.Args)
}
