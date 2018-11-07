# stanford-pos-tagger
Golang wrapper for stanford pos tagger, with support for Chinese

This package extends the functionality of the [go-stanford-wrapper by kamildrazkiewicz](https://github.com/kamildrazkiewicz/go-stanford-nlp).
I added support for Chinese part of speech tagging by updating the delimiters and the POS descriptions. 

## install
`go get github.com/kyajpauley/stanford-pos-tagger`

## usage
Input strings for tagging need to be pre-tokenized and delimited by whitespace.  

## license 
MIT