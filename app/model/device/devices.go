package device

import (
	"log"
	"time"
)

type Devices struct {
	Dev        *Device   `json:"dev,omitempty"`
	Devs       []*Device `json:"devs"`
	devs       map[string]*Device
	pingSecond time.Duration
}

func NewDevices(pingSecond time.Duration) *Devices {
	return &Devices{
		Devs:       []*Device{},
		devs:       map[string]*Device{},
		pingSecond: pingSecond,
	}

}

func (d *Devices) WithDevice(Buvid string) *Device {
	for _, dev := range d.devs {
		if dev.Buvid == Buvid {
			return dev
		}
	}
	return nil
}

func (d *Devices) WithDeviceMd5(Md5Buvid string) *Device {
	for _, dev := range d.devs {
		if dev.Md5 == Md5Buvid {
			return dev
		}
	}
	return nil
}

func (d *Devices) Check(buvid, ua string) {
	dev := d.WithDevice(buvid)
	if dev != nil {
		dev.LastPing = time.Now()
	} else {
		d.devs[buvid] = NewDevice(buvid, ua)
	}

	var (
		delKey []string
	)

	for _, dev := range d.devs {
		if dev.LastPing.Add(d.pingSecond).Before(time.Now()) {
			delKey = append(delKey, dev.Buvid)
		}
	}
	for _, k := range delKey {
		log.Printf("delete dev %s", k)
		delete(d.devs, k)
	}
}

func (d *Devices) Split(devs *Devices, buvid string) {
	d.Dev = devs.WithDevice(buvid)
	d.Dev.SplitSync()
	var newDevs []*Device
	for _, dev := range devs.devs {
		if dev.Buvid != buvid {
			newDevs = append(newDevs, dev)
		}
	}
	d.Devs = newDevs
}
