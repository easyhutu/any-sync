package device

import (
	"github.com/easyhutu/any-sync/app/utils"
	"path"
	"strings"
	"time"
)

type Device struct {
	Buvid    string    `json:"buvid,omitempty"`
	Ua       string    `json:"ua,omitempty"`
	Md5      string    `json:"md_5,omitempty"`
	Show     string    `json:"show,omitempty"`
	LastPing time.Time `json:"last_ping,omitempty"`
	sync     []*SyncInfo
	SyncText []*SyncInfo   `json:"sync_text"`
	SyncFile []*SyncInfo   `json:"sync_file"`
	HasFiles []*SyncDetail `json:"has_files,omitempty"`
}

func NewDevice(buvid, ua string) *Device {
	dev := &Device{
		Buvid:    buvid,
		Ua:       ua,
		Md5:      buvid,
		LastPing: time.Now(),
		sync:     []*SyncInfo{},
		SyncFile: []*SyncInfo{},
		SyncText: []*SyncInfo{},
	}
	dev.Show = strings.Split(dev.Ua, ") ")[0] + ")"
	return dev
}

func (d *Device) GetSync(syncType SyncType) []*SyncInfo {
	var si []*SyncInfo
	for _, s := range d.sync {
		if s.SyncType == syncType {
			si = append(si, s)
		}
	}
	return si
}

func (d *Device) AddSync(info SyncInfo) {
	is_exist := false
	for _, s := range d.sync {
		if s.FromMd5 == info.FromMd5 && s.SyncType == info.SyncType {
			is_exist = true
			s.SyncTi = info.SyncTi
			if s.SyncType == SyncTypeText {
				s.Details = info.Details
			} else {
				s.Details = append(s.Details, info.Details[0])
			}
		}
	}
	if !is_exist {
		d.sync = append(d.sync, &info)
	}

}

func (d *Device) SplitSync() {
	d.SyncFile = []*SyncInfo{}
	d.SyncText = []*SyncInfo{}
	for _, s := range d.sync {
		switch s.SyncType {
		case SyncTypeText:
			d.SyncText = append(d.SyncText, s)
		case SyncTypeFile:
			d.SyncFile = append(d.SyncFile, s)
		}

	}
}

func (d *Device) AddFile(filename string, filepath string, size int64) *SyncDetail {
	for _, f := range d.HasFiles {
		if f.Content == filepath {
			return f
		}
	}
	sd := &SyncDetail{
		FileSize: size,
		Desc:     filename,
		FileExt:  strings.ToLower(path.Ext(filename)),
		Content:  filepath,
		SizeShow: utils.FileSizeFormat(size),
	}
	sd.CheckCanPreview()
	d.HasFiles = append(d.HasFiles, sd)
	return sd
}

func (d *Device) FilterFile(filename string) *SyncDetail {
	for _, f := range d.HasFiles {
		if f.Desc == filename {
			return f
		}
	}
	return nil
}
