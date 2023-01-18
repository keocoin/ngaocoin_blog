package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gosimple/slug"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func loadDataFromMarkdown(filePath string) (metadata PostMetadata, html string) {
	path, err := os.Open(filePath)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	defer path.Close()
	source, _ := ioutil.ReadAll(path)

	var buf bytes.Buffer
	context := parser.NewContext()

	if err := markdown.Convert([]byte(source), &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}

	metaContext := meta.Get(context)

	jsonStr, err := json.Marshal(metaContext)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(jsonStr, &metadata); err != nil {
		fmt.Println(err)
	}

	var categorySlug string
	var titleSlug string

	if metadata.Author == "" {
		metadata.Author = "Editor"
	}

	if metadata.PublishDate == "" {
		metadata.PublishDate = "2023-01-17"
	}

	if metadata.Category != "" {
		categorySlug = slug.Make(metadata.Category)
	} else {
		categorySlug = slug.Make("posts")
		metadata.Category = "Posts"
	}

	if metadata.Slug != "" {
		titleSlug = slug.Make(metadata.Slug)
	} else {
		titleSlug = slug.Make(metadata.Title)
	}

	metadata.Slug = fmt.Sprintf("%s/%s.html", categorySlug, titleSlug)
	return metadata, buf.String()
}
