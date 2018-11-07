package main

import (
	"fmt"
	"github.com/kyajpauley/stanford-pos-tagger"
)

func main() {
	var (
		tagger *pos_tagger.Tagger
		res    []*pos_tagger.Result
		err    error
	)

	//needs to be the path relative to the pos_tagger.go file, or the absolute path on your system
	modelPath := "examples/chinese-distsim.tagger"
	taggerPath := "examples/stanford-postagger.jar"

	tagger, err = pos_tagger.NewTagger(modelPath, taggerPath)
	if err != nil {
		fmt.Println(err)
	}

	res, err = tagger.Tag("我来 到 北京 清华大学")
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range res {
		fmt.Println(r.Word, r.TAG)
	}
}
