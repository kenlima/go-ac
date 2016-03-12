package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/Users/jwlee/IdeaProjects/ci-autocomplete/src/main/resources/ciautocomplete/ci_autocomplete_admin.php.template")
	check(err)
	fmt.Println(string(dat))

	log.SetFlags(log.Lshortfile)
	dir := "/Users/jwlee/project/box/www/gagamel_admin"
	ignoreDirs := []string{".git"}
	allowExts := []string{".php"}

	cs := list.New()

	err2 := filepath.Walk(dir, printFile(ignoreDirs, allowExts, cs))
	if err2 != nil {
		log.Fatal(err)
	}
}

func printFile(ignoreDirs []string, allowExts []string, cs *list.List) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}

		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					//	log.Println(dir)
					return filepath.SkipDir
				}
			}
		} else {
			ext := filepath.Ext(path)
			for _, e := range allowExts {
				if e == ext {
					//	fmt.Println(path)
					return nil
				}
			}
		}

		//fmt.Println(path)

		return nil
	}
}
