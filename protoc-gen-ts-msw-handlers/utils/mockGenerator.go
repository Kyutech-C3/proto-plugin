package utils

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jaswdr/faker"
)

func MockGenerator(key string, propType string) string {
	if strings.Contains(key, "url") {
		return string("https://placehold.jp/150x150.png")
	}
	if strings.Contains(key, "name") {
		return string(faker.New().Person().Name())
	}
	if strings.Contains(key, "id") {
		return string(faker.New().UUID().V4())
	}
	if strings.Contains(key, "created_at") {
		return "2021-09-01T00:00:00Z"
	}
	if strings.Contains(key, "updated_at") {
		return "2021-09-01T00:00:00Z"
	}
	if strings.Contains(key, "color") {
		return string(faker.New().Color().Hex())
	}
	if strings.Contains(key, "title") {
		return string(faker.New().Lorem().Word())
	}
	if strings.Contains(key, "token") {
		return string(faker.New().Internet().Password())
	}
	if strings.Contains(key, "status") {
		return "ok"
	}
	if strings.Contains(key, "visibility") {
		return []string{"public", "private", "draft"}[rand.Intn(3)]
	}
	if strings.Contains(key, "description") || strings.Contains(key, "profile") {
		return string(faker.New().Lorem().Sentence(40))
	}
	if strings.Contains(key, "count") {
		return fmt.Sprint(rand.Intn(100))
	}

	return ""
}
