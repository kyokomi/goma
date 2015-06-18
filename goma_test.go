package goma

import (
	"fmt"
	"testing"
	"time"
)

type HogeType string

const (
	OKType HogeType = "OK"
)

func TestGenerateQuery(t *testing.T) {
	argsMap := map[string]interface{}{
		"description":             "hogehoge",
		"image_url":               "http://hgehoge.com",
		"pickup_link_type":        OKType,
		"manga_id":                "aaasddsd",
		"content_id":              "dsddasdaa",
		"link_dest_url":           "http://hgehoge.com",
		"published_at":            time.Now(),
		"publish_finish_datetime": time.Now(),
		"publishing":              false,
		"pickup_order":            1,
		"created_at":              time.Now(),
		"updated_at":              time.Now(),
	}

	query, args, err := GenerateQuery(`insert into pickup(
description
, image_url
, pickup_link_type
, manga_id
, content_id
, link_dest_url
, published_at
, publish_finish_datetime
, publishing
, pickup_order
, created_at
, updated_at
) values(

  /* :description */'1111'
, :image_url
, :pickup_link_type
, :manga_id
, :content_id
, :link_dest_url
, :published_at
, :publish_finish_datetime
, :publishing
, :pickup_order
, :created_at
, :updated_at
)
`, argsMap)
	fmt.Println(query, args, err)
}
