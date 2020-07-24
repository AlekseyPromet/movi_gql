// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type NewVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
	URL         string `json:"url"`
}

type Screenshot struct {
	ID      string `json:"id"`
	VideoID string `json:"videoId"`
	URL     string `json:"url"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Video struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	User        *User         `json:"user"`
	URL         string        `json:"url"`
	CreatedAt   time.Time     `json:"createdAt"`
	Screenshots []*Screenshot `json:"screenshots"`
	Related     []*Video      `json:"related"`
}

func MarshalID(id int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

func UnmarshalID(v interface{}) (int, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return int(i), e
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("Timestamp error")
}
