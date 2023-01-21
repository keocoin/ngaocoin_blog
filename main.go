package main

import (
	"html/template"
	"log"
	"path/filepath"
	"sort"
	"fmt"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/renderer/html"
)

var PostsDir = "./posts"
var OutputDir = "./output"
var TemplateDir = "./templates"

var markdown = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithXHTML(),
	),
)

type PostMetadata struct {
	Title       string `json:"Title"`
	Slug        string `json:"Slug"`
	Thumb       string `json:"Thumb"`
	Description string `json:"Description"`
	Category    string `json:"Category"`
	Author      string `json:"Author"`
	PublishDate string `Json:"PublishDate"`
}

type Post struct {
	Metadata PostMetadata
	Html     template.HTML
}

type DetailData struct {
	SinglePost   Post
	Others       []PostMetadata
	SiteMetadata map[string]string
}

type HomeData struct {
	Posts        []PostMetadata
	SiteMetadata map[string]string
}

var siteMetadata = map[string]string{
	"SiteName": "Ngáo Coin",
	"Title":    "Đồng hành cùng coin thủ Việt",
	"Domain":   "ngaocoin.com",
	"Slogan":   "Ngáo không phải là tội, Ngáo là một loại bệnh",
	"Description":   "chia sẻ thông tin kiến thức, research về blockchain, cryptocurrency",
	"Type" : "website",
	"Image" : "https://i.ibb.co/VLMNBBx/FVGsa-Y-X0-AEewjw.jpg",
	"URL" : "https://ngaocoin.com",
}

func loadPostData() []*Post {
	files, err := listFilesInDir(PostsDir)

	if err != nil {
		log.Fatal(err)
	}

	var data []*Post

	for _, filePath := range files {
		metaData, html := loadDataFromMarkdown(filePath)

		post := &Post{
			Metadata: metaData,
			Html:     template.HTML(html),
		}
		data = append(data, post)
	}

	return data
}

func genDetailpage(post *Post, others []PostMetadata) {
	detailMetadata := siteMetadata
	detailMetadata["Title"] = post.Metadata.Title
	detailMetadata["Type"] = "article"
	detailMetadata["URL"] = fmt.Sprintf("%s/%s", "https://ngaocoin.com", post.Metadata.Slug)
	if(post.Metadata.Thumb != ""){
		detailMetadata["Image"] = post.Metadata.Thumb
	}

	data := &DetailData{
		SinglePost:   *post,
		Others:       others,
		SiteMetadata: detailMetadata,
	}

	outputPath := filepath.Join(OutputDir, post.Metadata.Slug)
	genHTMLFile("single_post", data, outputPath)
}

func genHomepage(posts []PostMetadata) {

	var data = &HomeData{
		Posts:        posts,
		SiteMetadata: siteMetadata,
	}

	outputPath := filepath.Join(OutputDir, "index.html")
	genHTMLFile("home_page", data, outputPath)
}

func main() {
	posts := loadPostData()
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Metadata.PublishDate > posts[j].Metadata.PublishDate
	})

	var postsMetadata []PostMetadata

	for _, post := range posts {
		postsMetadata = append(postsMetadata, post.Metadata)
	}

	genHomepage(postsMetadata)

	for _, post := range posts {
		if len(postsMetadata) > 3 {
			genDetailpage(post, postsMetadata[:3])
		} else {
			genDetailpage(post, postsMetadata)
		}
	}
}
