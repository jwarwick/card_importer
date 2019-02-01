package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
)

func main() {
	tokenFlag := flag.String("token", "", "Github OAuth Token")
	repoFlag := flag.String("repo", "", "Github repo '<owner>/<repo>'")
	flag.Parse()

	if "" == *tokenFlag || "" == *repoFlag || 1 != len(flag.Args()) {
		fmt.Println("Usage: import -token <OAuth Token> -repo <owner/repo> cards.csv")
	} else {
		issues := parseFile(flag.Args()[0])
		createIssues(issues, tokenFlag, repoFlag)
	}
}

type issue struct {
	Name        string   `json:"title"`
	Description string   `json:"body"`
	Label       []string `json:"labels"`
}

func createIssues(issues []issue, token *string, repo *string) {
	const issueUrl = "https://api.github.com/repos"
	u, err := url.Parse(issueUrl)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, *repo)
	u.Path = path.Join(u.Path, "issues")

	client := &http.Client{}
	count := 0
	for _, issue := range issues {
		issueJson, err := json.Marshal(&issue)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(issueJson))
		if err != nil {
			panic(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("token %s", *token))
		req.Header.Add("Accept", "application/vnd.github.v3+json")
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		count++
	}

	fmt.Println("Created", count, "issues")
}

func parseFile(filename string) []issue {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = ';'

	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	issues := make([]issue, len(lines)-1)
	const titleField int = 17
	const textField int = 16
	skipLine := true
	for idx, line := range lines {
		if skipLine {
			skipLine = false
			continue
		}
		issues[idx-1] = issue{line[titleField], line[textField], []string{"New Card"}}
	}

	return issues
}
