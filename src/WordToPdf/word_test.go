package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_aa(t *testing.T) {
	now := time.Now()
	wordConvert("F:/wordToPdf/word.docx", "F:/wordToPdf/1.pdf", 17)
	//FileInsertsAndPdf("F:/wordToPdf/aa.doc", "F:/wordToPdf/word.docx","F:/wordToPdf/p.pdf","1","2","3","4","5","6","7","8","9","10","11","12","13")
	fmt.Println("END:   ", time.Since(now))
}
