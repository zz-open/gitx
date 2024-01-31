package git

import (
	"encoding/json"
	"log"

	"github.com/zz-open/tpl/internal/httpx"
)

const GithubProvideUrl = "https://api.github.com"
const githubDownloadUrl = "https://raw.githubusercontent.com/"

func RequestGithubProvideUrl() {
	b, err := httpx.Get(GithubProvideUrl)
	if err != nil {
		return
	}

	var m = make(map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return
	}
	log.Println(m["authorizations_url"])
}
