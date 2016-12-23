package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	pb "github.com/monmaru/github-tren-go/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type ght struct{}

func (g *ght) Fetch(c context.Context, req *pb.Req) (*pb.Res, error) {
	if req.Lang == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "language name is empty.")
	}

	doc, err := goquery.NewDocument("https://github.com/trending?lang=" + req.Lang)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "url scarapping failed")
	}

	log.Printf("[START] Fetch: %s", req.Lang)
	repos := []*pb.Repository{}
	doc.Find(".repo-list li").Each(func(n int, s *goquery.Selection) {
		url, _ := s.Find("h3 a").Attr("href")
		refs := strings.Split(url, "/")

		repos = append(repos, &pb.Repository{
			Rank:        int32(n),
			Owner:       refs[1],
			Title:       refs[2],
			Url:         "https://github.com" + url,
			Description: strings.Trim(s.Find("div[class=py-1] > p").Text(), " "),
			Lang:        strings.Trim(s.Find("span[class=float-right]").Text(), " "),
			Star: &pb.Star{
				Today: parseInt(s.Find("span[class=float-right]").Text()),
				Total: parseInt(s.Find("span[class=float-right]").Text()),
			},
		})
	})

	log.Printf("[END] Fetch: %s", req.Lang)
	return &pb.Res{Repositories: repos}, nil
}

func parseInt(s string) int32 {
	num, _ := strconv.ParseInt(s, 10, 32)
	return int32(num)
}

func main() {
	var addrFlag = flag.String("addr", "localhost:5000", "Address host:post")
	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("tcp connection error")
	}

	s := grpc.NewServer()
	pb.RegisterGithubTrendingServer(s, &ght{})
	s.Serve(lis)
}
