package device

import "time"

type Devices struct {
	Dev        *Device   `json:"dev,omitempty"`
	Devs       []*Device `json:"devs"`
	pingSecond time.Duration
}

func NewDevices(pingSecond time.Duration) *Devices {
	return &Devices{
		Devs:       []*Device{},
		pingSecond: pingSecond,
	}

}

func (d *Devices) WithDevice(Buvid string) *Device {
	for _, dev := range d.Devs {
		if dev.Buvid == Buvid {
			return dev
		}
	}
	return nil
}

func (d *Devices) WithDeviceMd5(Md5Buvid string) *Device {
	for _, dev := range d.Devs {
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
		d.Devs = append(d.Devs, NewDevice(buvid, ua))
	}

	var newDevs []*Device
	for _, dev := range d.Devs {
		if dev.LastPing.Add(d.pingSecond).After(time.Now()) {
			dev.LastPing = time.Now()
			newDevs = append(newDevs, dev)
		}
	}
	d.Devs = newDevs

}

func (d *Devices) Split(devs *Devices, buvid string) {
	d.Dev = devs.WithDevice(buvid)
	d.Dev.SplitSync()
	var newDevs []*Device
	for _, dev := range devs.Devs {
		if dev.Buvid != buvid {
			newDevs = append(newDevs, dev)
		}
	}
	d.Devs = newDevs
}
