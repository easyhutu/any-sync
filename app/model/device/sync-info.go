package device

import (
	"strings"
	"time"
)

const (
	SyncPending    = SyncStatus(1)
	SyncOK         = SyncStatus(2)
	SyncTypeText   = SyncType("text")
	SyncTypeFile   = SyncType("file")
	SyncTypeIgnore = SyncType("ignore")
)

type SyncStatus int
type SyncType string

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
	ShowTI   string        `json:"show_ti,omitempty"`
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

func (s *SyncInfo) Generate() {
	s.ShowTI = s.SyncTi.Format("2006-01-02 15:04:05")
}

func WithSyncType(val string) SyncType {
	switch val {
	case string(SyncTypeFile):
		return SyncTypeFile
	case string(SyncTypeText):
		return SyncTypeText
	default:
		return SyncTypeIgnore

	}
}
