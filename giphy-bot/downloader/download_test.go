package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImageWithProxy(t *testing.T) {
	err := DownloadFileWithProxy("temp.gif", "", "https://i.giphy.com/media/cZ7rmKfFYOvYI/giphy.gif")
	if err != nil {
		t.Error(err)
	}
	assert.FileExists(t, "cZ7rmKfFYOvYI.gif")
}

func TestGetImageParrelly(t *testing.T) {
	err := DownloadParallelly("temp.gif", "https://i.giphy.com/media/cZ7rmKfFYOvYI/giphy.gif")
	if err != nil {
		t.Error(err)
	}
	assert.FileExists(t, "temp.gif")
}
