package godingtalk

import (
	"io"
	"net/url"
)

//MediaResponse is
type MediaResponse struct {
	OAPIResponse
	Type    string
	MediaID string `json:"media_id"`
}

//UploadMedia is to upload media file to DingTalk
func (c *DingTalkClient) UploadMedia(mediaType string, filename string, reader io.Reader) (media MediaResponse, err error) {
	upload := UploadFile{
		FieldName: "media",
		FileName:  filename,
		Reader:    reader,
	}
	params := url.Values{}
	params.Add("type", mediaType)
	err = c.httpRPC("media/upload", params, upload, &media)
	return media, err
}