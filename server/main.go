package main

import (
	"flag"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"
	pb "github.com/monmaru/grpc-sample/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// QS ...
type QS struct {
	Lang string `url:"l"`
}

var reg = regexp.MustCompile("(.*) stars today")

type ght struct{}

func (g *ght) Fetch(c context.Context, req *pb.Req) (*pb.Res, error) {
	if req.Lang == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "language name is empty.")
	}

	v, _ := query.Values(QS{Lang: req.Lang})
	url := "https://github.com/trending?" + v.Encode()
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "scarapping failed")
	}

	log.Printf("[START] Fetch: %s", req.Lang)
	repos := []*pb.Repository{}
	doc.Find(".repo-list li").Each(func(idx int, s *goquery.Selection) {
		repos = append(repos, parse(idx, s))
	})

	log.Printf("[END] Fetch: %s", req.Lang)
	return &pb.Res{Repositories: repos}, nil
}

func parse(idx int, s *goquery.Selection) *pb.Repository {
	href, _ := s.Find("h3 a").Attr("href")
	refs := strings.Split(href, "/")
	description := sanitize(s.Find("div[class=py-1] > p").Text())
	lang := sanitize(s.Find("span[itemprop=programmingLanguage]").Text())
	todayStar := sanitize(s.Find("span[class=float-right]").Text())
	if todayStar == "" {
		todayStar = "-"
	} else {
		todayStar = reg.FindStringSubmatch(todayStar)[1]
	}
	totalStar := sanitize(s.Find("a[aria-label=Stargazers]").Text())

	return &pb.Repository{
		Rank:        int32(idx + 1),
		Owner:       refs[1],
		Title:       refs[2],
		Url:         "https://github.com" + href,
		Description: description,
		Lang:        lang,
		Star: &pb.Star{
			Today: todayStar,
			Total: totalStar,
		},
	}
}

func sanitize(s string) string {
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, " ")
	s = strings.Trim(s, "\n")
	return strings.Trim(s, " ")
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
