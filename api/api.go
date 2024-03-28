package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "ppdiff-overlay/token"
	// "bufio"
	// "encoding/binary"
	// "errors"
	// "fmt"
	// "io"
	// "log"
	// "os"
	// "path/filepath"
	// "runtime/debug"
	// "strings"
	// "time"
	// "github.com/l3lackShark/gosumemory/memory"
	// "github.com/k0kubun/pp"
)

type BasicUserData struct {
	AvatarURL   string `json:"avatar_url"`
	CountryCode string `json:"country_code"`
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Cover       struct {
		URL string `json:"url"`
	} `json:"cover"`
	Statistics struct {
		GlobalRank  int     `json:"global_rank"`
		Pp          float64 `json:"pp"`
		RankedScore int64   `json:"ranked_score"`
		HitAccuracy float64 `json:"hit_accuracy"`
		CountryRank int     `json:"country_rank"`
	} `json:"statistics"`
	SupportLevel int `json:"support_level"`
}

func GetOsuUserData(userid string) (*BasicUserData, error) {

	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/users/%s", userid)
	method := "GET"

	client := &token.Client
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	var data BasicUserData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
