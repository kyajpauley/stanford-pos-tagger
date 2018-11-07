package main

import (
	"fmt"
	"github.com/kyajpauley/stanford-pos-tagger"
)

func main() {
	var (
		tagger *pos_tagger.Tagger
		res    []*pos_tagger.Result
		//err error
	)

	path := "/home/kya/Documents/stanford-postagger-full-2018-10-16/"
	modelPath := fmt.Sprintf("%s/models/chinese-distsim.tagger", path)
	taggerPath := fmt.Sprintf("%s/stanford-postagger.jar", path)

	tagger, _ = pos_tagger.NewTagger(modelPath, taggerPath)

	res, _ = tagger.Tag("我来 到 北京 清华大学")

	for _, r := range res {
		fmt.Println(r.Word, r.TAG)
	}
}
