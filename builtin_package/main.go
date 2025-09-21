package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const APIKEY = "193ef3a"

// 구조체에서 태그를 통해 json변수의 이름과 구조체 변수의 이름을 매치시킬 수 있음
type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbId     string `json:"imdbID"`
}

// http 요청을 날려서 응답의 body를 문자열 형태로 반환하는 함수
func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}

// 이름으로 영화 정보를 검색하는 함수
func searchByName(name string) (*MovieInfo, error) {
	// 쿼리 파라미터 설정
	params := &url.Values{}
	params.Set("apikey", APIKEY)
	params.Set("t", name)
	siteURL := "http://www.omdbapi.com/?" + params.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}
	// 빈 구조체를 반환하는 대신 nil을 반환하기 위해 MovieInfo의 포인터를 사용
	mi := &MovieInfo{}

	// 언마샬한 body를 mi포인터가 가리키는 공간에 작성
	// Unmarshal함수는 error를 리턴값으로 가져서 다음과 같이 작성하면 Unmarshal과정에서 에러가 난 경우 nil, error 반환 가능
	return mi, json.Unmarshal([]byte(body), mi)
}

// id로 영화 정보를 검색하는 함수
func searchById(id string) (*MovieInfo, error) {
	// 쿼리 파라미터 설정
	params := &url.Values{}
	params.Set("apikey", APIKEY)
	params.Set("i", id)
	siteURL := "http://www.omdbapi.com/?" + params.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}
	// 빈 구조체를 반환하는 대신 nil을 반환하기 위해 MovieInfo의 포인터를 사용
	mi := &MovieInfo{}

	// 언마샬한 body를 mi포인터가 가리키는 공간에 작성
	// Unmarshal함수는 error를 리턴값으로 가져서 다음과 같이 작성하면 Unmarshal과정에서 에러가 난 경우 nil, error 반환 가능
	return mi, json.Unmarshal([]byte(body), mi)
}

func main() {
	mi1, _ := searchByName("Game of")
	fmt.Println(mi1.Title)
	mi2, _ := searchById("tt3896198")
	fmt.Println(mi2.Title)
}
