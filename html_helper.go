package main

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func genHTMLFile(template_name string, data any, outputPath string) {
	var outputHtml bytes.Buffer
	templates := loadTemplate()
	templates.ExecuteTemplate(&outputHtml, template_name, data)

	os.MkdirAll(filepath.Dir(outputPath), 0770)
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(outputHtml.String())
	w.Flush()
}

func listFilesInDir(TemplateDir string) (files []string, err error) {
	err = filepath.Walk(TemplateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func loadTemplate() (templates *template.Template) {
	allPaths, err := listFilesInDir(TemplateDir)

	if err != nil {
		log.Fatal(err)
	}

	templates = template.Must(template.New("").ParseFiles(allPaths...))
	return templates
}
