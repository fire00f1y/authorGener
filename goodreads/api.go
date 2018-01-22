package goodreads

import (
	"strings"
	"net/http"
	"fmt"
	"github.com/fire00f1y/authorGener/model"
	"encoding/xml"
)

const (
	byNameUrl = "https://www.goodreads.com/api/author_url/{name}?key={key}"
	byIdUrl   = "https://www.goodreads.com/author/show/{id}?format=xml&key={key}"
)

func GetAuthorId(name, key string, retry bool) (id string, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("Panic while getting id for author [%s]. Will not be in total count: %v\n", name, r)
		}
	}()
	replacer := strings.NewReplacer(
		"{name}",strings.TrimSpace(name),
		"{key}",key,
	)
	url := replacer.Replace(byNameUrl)

	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		fmt.Printf("[%d] code returned when looking up [%s] at url [%s]\n", resp.StatusCode, name, url)
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r model.NameLookupResponse
	err = xml.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Printf("Failed to get ID for [%s] using url [%s]: %v\n", name, url, err)
		return "",err
	}

	id = r.Author.Id

	return
}

func CorrectedName(name string) string {
	splitName := strings.Split(name, " ")
	if len(splitName) == 3 {
		return strings.TrimSpace(splitName[0]) + " " + strings.TrimSpace(splitName[2])
	}
	excludeList := make([]int, 0)
	for i, n := range splitName {
		if len(n) == 2 || strings.Contains(n, ".") || strings.Contains(n, "(") || strings.Contains(n, ")"){
			excludeList = append(excludeList, i)
		}
	}
	newString := ""
	for i, v := range splitName {
		if listContainsInt(excludeList, i) {
			continue
		}
		newString += v + " "
	}
	return strings.TrimSpace(newString)
}

func listContainsInt(list []int, i int) bool {
	for _,v := range list {
		if v == i {
			return true
		}
	}
	return false
}

func GetAuthorInfo(id, key string) (gender string, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("Panic while getting info for author [%s]. Will be counted as an 'unknown': %v\n", id, r)
		}
	}()
	replacer := strings.NewReplacer(
		"{id}",id,
		"{key}",key,
	)
	url := replacer.Replace(byIdUrl)

	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		fmt.Printf("[%d] code returned when looking up %s\n", resp.StatusCode, id)
	}
	if err != nil {
		fmt.Printf("Failed to get info for [id:%s] using url [%s]: %v\n", id, url, err)
		return "", err
	}
	defer resp.Body.Close()

	var r model.IdLookupResponse
	err = xml.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Printf("Failed to get info for [id:%s] using url [%s]: %v\n", id, url, err)
		return "",err
	}
	gender = r.Author.Gender
	return
}