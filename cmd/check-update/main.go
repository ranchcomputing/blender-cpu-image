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
	tagName, err := githubLatestTag()
	if err != nil {
		return err
	}
	log.Println("latest remote tag:", tagName)

	ok, err := tagIsLocal(tagName)
	if err != nil {
		return err
	}
	if ok {
		log.Println("already present locally")
		return nil
	}
	fmt.Println(tagName)

	return nil
}

func githubLatestTag() (string, error) {
	req, err := http.NewRequest("GET", baseURL+"tags", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	type githubResponse struct {
		Name string `json:"name"`
	}
	var gr []githubResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		return "", err
	}
	if len(gr) == 0 {
		return "", errors.New("got no tags")
	}
	for _, r := range gr {
		if !strings.HasPrefix(r.Name, "v") {
			continue
		}
		return r.Name, nil
	}
	return "", errors.New("got no usable tag")
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
