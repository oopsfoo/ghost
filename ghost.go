package main

import (
	"context"
	"fmt"
	"ghost/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/google/go-github/github"
	"log"
	"net"
	"os"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}

func GetGithub() {
	//ctx := context.Background()
	//ts := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: "... your access token ..."},
	//)
	//tc := oauth2.NewClient(ctx, ts)
	//client := github.NewClient(tc)
	client := github.NewClient(nil)
	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(context.Background(), "oopsfoo", nil)
	if err != nil {
		log.Fatal(err)
	}

	for i, repo := range repos {
		fmt.Printf("[%d] %s %s\n", i, *repo.Name, *repo.CloneURL)

	}
}

func CloneFromGithub() {
	//Info("git clone https://github.com/oopsfoo/ghost_oopsfoo.git")
	repoURL := "https://github.com/oopsfoo/ghost_oopsfoo.git"
	repoUser := "oopsfoo@gmail.com"
	repoToken := "c74ae2834ad5702515a1d9a330d75a2f72b0bf74"
	options := &git.CloneOptions{
		Auth:          &http.BasicAuth{repoUser, repoToken},
		URL:           repoURL,
		Depth:         500,
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
		SingleBranch:  true,
		Tags:          git.NoTags,
		Progress:      os.Stdout,
	}

	//r, err := git.PlainClone("/tmp/foo2", false, options)
	r, err := git.Clone(memory.NewStorage(), nil, options)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	w.Checkout(nil)

	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ref.String())
}

func main() {
	m := util.NewHostMap(GetHostname(), GetOutboundIP())
	println(m.String())
}
