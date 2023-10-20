package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)

		os.Exit(1)
	}
}

const baseURL = "https://api.github.com/repos/blender/blender/"

func run() error {
	log.Println("checking:", baseURL)
	tagNames, err := githubLatestTags()
	if err != nil {
		return err
	}

	tagNames = latestMajorTags(tagNames)
	log.Println("latest remote tags:", tagNames)

	tagName, err := firstNonLocalTag(tagNames)
	if err != nil {
		return err
	}
	if tagName == "" {
		log.Println("already present locally")
		return nil
	}
	fmt.Println(tagName)

	return nil
}

func githubLatestTags() ([]string, error) {
	req, err := http.NewRequest("GET", baseURL+"tags", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type githubResponse struct {
		Name string `json:"name"`
	}
	var gr []githubResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		return nil, err
	}
	tags := make([]string, 0, len(gr))
	for _, r := range gr {
		if !strings.HasPrefix(r.Name, "v") {
			continue
		}
		tags = append(tags, r.Name)
	}
	if len(tags) == 0 {
		return nil, errors.New("got no usable tag")
	}
	return tags, nil
}

func latestMajorTags(sortedTags []string) []string {
	seen := make(map[string]bool)
	tags := make([]string, 0)
	for _, t := range sortedTags {
		major := extractMajor(t)
		if seen[major] {
			continue
		}
		seen[major] = true
		tags = append(tags, t)
	}
	return tags
}

func extractMajor(tag string) string {
	major, tail, found := strings.Cut(tag, ".")
	if !found {
		panic(fmt.Sprintf("could not find first dot in %s", tag))
	}
	minor, _, found := strings.Cut(tail, ".")
	if !found {
		panic(fmt.Sprintf("could not find second dot in %s", tag))
	}
	return major + "." + minor
}

func firstNonLocalTag(tagNames []string) (string, error) {
	for _, tagName := range tagNames {
		ok, err := tagIsLocal(tagName)
		if err != nil {
			return "", err
		}
		if ok {
			continue
		}

		return tagName, nil
	}
	// all tags are local, return an empty string
	return "", nil
}

func tagIsLocal(tagName string) (bool, error) {
	cmd := exec.Command("git", "tag", "-l", tagName)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(out)) == tagName, nil
}
