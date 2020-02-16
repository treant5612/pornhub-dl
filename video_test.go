package pornhub_dl

import (
	"log"
	"testing"
)

func TestGetVideoInfoByKey(t *testing.T) {
	video,err := GetVideoInfoByKey("ph5d35a4f825603")

	if err!=nil{
		t.Fatal(err)
	}
	log.Println(video)
}
