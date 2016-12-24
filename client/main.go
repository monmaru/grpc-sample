package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/monmaru/grpc-sample/proto"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	addrFlag = flag.String("addr", "localhost:5000", "server address host:post")
)

func fetch(lang string) ([]*pb.Repository, error) {
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	defer conn.Close()
	c := pb.NewGithubTrendingClient(conn)

	resp, err := c.Fetch(context.Background(), &pb.Req{Lang: lang})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
		return nil, err
	}

	return resp.Repositories, nil
}

func show(repos []*pb.Repository) {
	if len(repos) == 0 {
		fmt.Println("no repository.")
	}

	for _, r := range repos {
		fmt.Printf("[Rank]: %v\n", r.Rank)
		fmt.Printf("[Lang]: %v\n", r.Lang)
		fmt.Printf("[Title]: %v\n", r.Title)
		fmt.Printf("[Owner]: %v\n", r.Owner)
		fmt.Printf("[Url]: %v\n", r.Url)
		fmt.Printf("[Description]: %v\n", r.Description)
		fmt.Printf("[Today Star]: %v\n", r.Star.Today)
		fmt.Printf("[Total Star]: %v\n", r.Star.Total)
		fmt.Println("-----------------------------------------------")
	}
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

		repos, err := fetch(lang)

		if err != nil {
			return err
		}

		show(repos)
		return nil
	}

	app.Run(os.Args)
}
