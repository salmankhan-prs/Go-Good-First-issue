package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

const (
	goodFirstIssuesURL = "https://api.github.com/search/repositories?q=language:go+good-first-issues:>2+archived:false+is:public&sort=updated"
	readmeFileURL      = "https://api.github.com/repos/salmankhan-prs/Go-Good-First-issue/contents/README.md?ref=master"
)

var (
	client      = http.Client{}
	apiToken    string
	IssueCount  int
	NAME        string
	EMAIl       string
	repoDetails = RepoDetails{}
)

type GoodFirstIssue struct {
	RepoFullName string
	RepoStars    int
	RepoIssue    []GoodIssue
}

type GoodIssue struct {
	IssueHtmlURL string
	IssueName    string
}

type RepoDetails struct {
	Sha string `json:"sha"`
}

type Repository struct {
	FullName string `json:"full_name"`
	Stars    int    `json:"stargazers_count"`
}

type Issue struct {
	HtmlURL string `json:"html_url"`
	Title   string `json:"title"`
}

var issueCount = 0

func getGoodFirstIssue() {
	var goodFirstIssueArray []GoodFirstIssue
	var page = 1

	for issueCount <= IssueCount {
		repos, err := getRepositoriesWithGoodFirstIssues(page)
		page++
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, repo := range repos {
			issues, err := getIssuesWithGoodFirstLabel(repo.FullName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			var issueArray []GoodIssue
			for _, issue := range issues {
				issueArray = append(issueArray, GoodIssue{IssueHtmlURL: issue.IssueHtmlURL, IssueName: issue.IssueName})
			}
			goodFirstIssueArray = append(goodFirstIssueArray, GoodFirstIssue{RepoFullName: repo.FullName, RepoStars: repo.Stars, RepoIssue: issueArray})
		}
	}
	markdown := generateMarkdown(goodFirstIssueArray)

	if err := pushMarkdownToRepo(markdown); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("GO GOOD FIRST ISSUE IS DONE !")
}

func getRepositoriesWithGoodFirstIssues(page int) ([]Repository, error) {
	req, err := http.NewRequest("GET", goodFirstIssuesURL+"&page="+strconv.Itoa(page), nil)
	page = page + 1
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get response: %s", res.Status)
	}
	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var objmap map[string][]Repository
	err = json.Unmarshal(dataBytes, &objmap)
	if err != nil {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("failed to unmarshal json: %s\n", err)
		fmt.Printf("response body: %s\n", body)

	}
	return objmap["items"], nil
}

func getIssuesWithGoodFirstLabel(repoName string) ([]GoodIssue, error) {
	// build query string
	issuesUrl := fmt.Sprintf("https://api.github.com/repos/%s/issues", repoName)
	issueReq, err := http.NewRequest("GET", issuesUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create issue request: %w", err)
	}

	query := issueReq.URL.Query()
	query.Add("state", "open")
	query.Add("state", "open")
	query.Add("labels", "good first issue")
	query.Add("sort", "updated")
	issueReq.URL.RawQuery = query.Encode()

	auth := "Bearer " + apiToken
	issueReq.Header.Set("Authorization", auth)

	issueRes, err := client.Do(issueReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch issues for repo %s: %w", repoName, err)
	}

	defer issueRes.Body.Close()

	issuesDataByte, err := io.ReadAll(issueRes.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read issues response body for repo %s: %w", repoName, err)
	}

	var issueObj []struct {
		HtmlUrl   string    `json:"html_url"`
		Title     string    `json:"title"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	if err := json.Unmarshal(issuesDataByte, &issueObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal issues response body for repo %s: %w", repoName, err)
	}

	var issueArray []GoodIssue
	weekBefore := time.Now().AddDate(0, 0, -30)
	for _, issue := range issueObj {
		if issue.CreatedAt.After(weekBefore) {
			issueCount++
			issueArray = append(issueArray, GoodIssue{IssueHtmlURL: issue.HtmlUrl, IssueName: issue.Title})
		}
	}

	return issueArray, nil
}

func generateMarkdown(goodFirstIssueArray []GoodFirstIssue) string {
	markdown := "<h1 align='center'><span style='font-size:36px;'>Go Good First Issue</span></h1>\n\n"
	markdown += "<div align='center' style='font-size:20px;'>Welcome to Go's good first issue list! These are issues that are great for new contributors to the Go community. It is updated every 24 hours.</div> <br>\n\n\n\n"
	currentTime := time.Now()
	markdown += "<div align='center'>Last updated at " + currentTime.Format("January 2, 2006 15:04 MST") + ".</div>\n\n"

	for _, repoIssue := range goodFirstIssueArray {
		if len(repoIssue.RepoIssue) == 0 {
			continue
		}
		markdown += "\n## " + repoIssue.RepoFullName + " <span style='color:#F1C40F'>(" + beautifyNumber(repoIssue.RepoStars) + " ⭐️)</span>\n\n"
		for _, eachIssue := range repoIssue.RepoIssue {
			markdown += "- [" + eachIssue.IssueName + "](" + eachIssue.IssueHtmlURL + ")\n\n"
		}
	}

	return markdown
}

func beautifyNumber(num int) string {
	if num >= 1000 {
		return fmt.Sprintf("%.1fK", float64(num)/1000)
	}
	return fmt.Sprintf("%d", num)
}

func pushMarkdownToRepo(markdown string) error {
	repoReq, err := http.NewRequest("GET", readmeFileURL, nil)

	if err != nil {
		return err
	}
	repoReq.Header.Set("Authorization", "Bearer "+apiToken)

	repoRes, err := client.Do(repoReq)

	repoResBytes, _ := io.ReadAll(repoRes.Body)

	var repoMap RepoDetails

	err = json.Unmarshal(repoResBytes, &repoMap)
	sha := repoMap.Sha

	encodedMarkdown := base64.StdEncoding.EncodeToString([]byte(markdown))
	payload, err := json.Marshal(map[string]interface{}{
		"message": "Update README.md",
		"content": encodedMarkdown,
		"committer": map[string]interface{}{
			"name":  NAME,
			"email": EMAIl,
		},
		"sha": sha,
	})

	pushReq, err := http.NewRequest("PUT", readmeFileURL, bytes.NewBuffer(payload))

	if err != nil {
		return err
	}
	pushReq.Header.Set("Authorization", "Bearer "+apiToken)

	_, err = client.Do(pushReq)

	return nil

}

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(24).Hour().Do(getGoodFirstIssue)
	s.StartBlocking()
}
func main() {
	godotenv.Load(".env")
	apiToken = os.Getenv("API_TOKEN")
	IssueCount, _ = strconv.Atoi(os.Getenv("ISSUE_COUNT"))
	EMAIl = os.Getenv("EMAIL")
	NAME = os.Getenv("NAME")
	runCronJobs()
}
