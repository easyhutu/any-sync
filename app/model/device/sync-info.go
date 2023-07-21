package device

import (
	"strings"
	"time"
)

var (
	PreviewExt = []string{
		".png",
		".jpg",
		".jpeg",
		".mp4",
		".js",
		".html",
		".webp",
		".txt",
	}
)

type SyncInfo struct {
	SyncType SyncType      `json:"sync_type,omitempty"`
	From     string        `json:"from,omitempty"`
	FromMd5  string        `json:"from_md5,omitempty"`
	Details  []*SyncDetail `json:"details,omitempty"`
	Status   SyncStatus    `json:"status,omitempty"`
	SyncTi   time.Time     `json:"sync_ti,omitempty"`
}

type SyncDetail struct {
	// filename
	Desc string `json:"desc,omitempty"`
	// filepath
	Content    string `json:"content,omitempty"`
	FileExt    string `json:"file_ext,omitempty"`
	CanPreview bool   `json:"can_preview,omitempty"`
	Md5        string `json:"md5,omitempty"`
	FileSize   int64  `json:"file_size,omitempty"`
	SizeShow   string `json:"size_show,omitempty"`
}

func (sd *SyncDetail) CheckCanPreview() {
	for _, pe := range PreviewExt {
		if strings.ToLower(sd.FileExt) == pe {
			sd.CanPreview = true
		}
	}
}
