package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/AndrefHub/ppdiff-overlay/token"
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

	client := token.TokenConfig.Client(context.Background())
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// log.Println(string(body))

	var data BasicUserData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func AddOsuUserDataToChan(userid string, c chan *BasicUserData, wg *sync.WaitGroup) {
	defer wg.Done()
	if user, err := GetOsuUserData(userid); err != nil {
		log.Println(err)
	} else {
		c <- user
	}
}

// Returns encoded JSON
func GetUsersData(users ...string) []byte {
	var wg sync.WaitGroup
	c := make(chan *BasicUserData, 3)
	var arr []BasicUserData

	for _, user := range users {
		wg.Add(1)
		AddOsuUserDataToChan(string(user), c, &wg)
		time.Sleep(250 * time.Millisecond)
	}
	wg.Wait()
	close(c)

	for user := range c {
		arr = append(arr, *user)
	}
	if len(arr) == 0 {
		log.Println("FUCKING HELL WHY IM HERE")
		return []byte(`{"error": "Failed to get user's data, check your credentials or internet connection"}`)
	}
	jsonResponse, err := json.Marshal(arr)
	if err != nil {
		log.Println(err)
	}
	return jsonResponse
}
