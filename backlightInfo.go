package sstat

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// BacklightPath is the directory where the information for
// backlights are located.
const BacklightPath string = "/sys/class/backlight"

// BacklightInfo reports backlight information. Documentation for the
// object methods are taken from [sysfs-class-power].
//
// [sysfs-class-power]: https://www.kernel.org/doc/Documentation/ABI/stable/sysfs-class-backlight
type BacklightInfo struct {
	blPower          int
	brightness       int
	actualBrightness int
	maxBrightness    int
	typ              string
}

// BlPower reports BACKLIGHT power, values are
// compatible with FB_BLANK_* from fb.h
//
// Valid values are:
//   - 0 (FB_BLANK_UNBLANK)   : power on
//   - 4 (FB_BLANK_POWERDOWN) : power off
func (info *BacklightInfo) BlPower() (value int) {
	return info.blPower
}

// Brightness reports the brightness for this <backlight>. Values
// are between 0 and max_brightness. This value will also
// show the brightness level stored in the driver, which
// may not be the actual brightness
// (see [BacklightInfo.ActualBrightness]).
func (info *BacklightInfo) Brightness() (value int) {
	return info.brightness
}

// ActualBrightness reports the actual brightness by querying the hardware.
func (info *BacklightInfo) ActualBrightness() (value int) {
	return info.actualBrightness
}

// MaxBrightness reports the maximum brightness for <backlight>.
func (info *BacklightInfo) MaxBrightness() (value int) {
	return info.maxBrightness
}

// Type reports the type of interface controlled by <backlight>.
// In the general case, when multiple backlight
// interfaces are available for a single device, firmware
// control should be preferred to platform control should
// be preferred to raw control. Using a firmware
// interface reduces the probability of confusion with
// the hardware and the OS independently updating the
// backlight state. Platform interfaces are mostly a
// holdover from pre-standardisation of firmware
// interfaces.
//
// Valid values are:
//   - "firmware": The driver uses a standard firmware interface
//   - "platform": The driver uses a platform-specific interface
//   - "raw": The driver controls hardware registers directly
func (info *BacklightInfo) Type() (value string) {
	return info.typ
}

func (info *BacklightInfo) mapIntPtrs() map[string]*int {
	return map[string]*int{
		"bl_power":          &info.blPower,
		"brightness":        &info.brightness,
		"actual_brightness": &info.actualBrightness,
		"max_brightness":    &info.maxBrightness,
	}
}

func serveBacklight(backlightInfoChans map[string]<-chan *BacklightInfo, errChan chan<- error, basepath string) {
	var (
		backlightInfo, newBacklightInfo *BacklightInfo
		backlightInfoChan               chan *BacklightInfo
		watcher                         *fsnotify.Watcher
		event                           fsnotify.Event
		infoPath, infoName              string
		err                             error
	)

	backlightInfo, err = Backlight(basepath)
	if err != nil {
		errChan <- err

		return
	}

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		errChan <- err

		return
	}

	for _, infoPath = range []string{"bl_power", "brightness", "actual_brightness", "max_brightness", "type"} {
		err = watcher.Add(filepath.Join(BacklightPath, basepath, infoPath))
		if err != nil {
			errChan <- err

			return
		}
	}

	backlightInfoChan = make(chan *BacklightInfo)
	backlightInfoChans[basepath] = backlightInfoChan
	backlightInfoChan <- backlightInfo

	for {
		newBacklightInfo = new(BacklightInfo)
		*newBacklightInfo = *backlightInfo

		select {
		case event = <-watcher.Events:
			if !event.Has(fsnotify.Write) {
				continue
			}
		case err = <-watcher.Errors:
			errChan <- err

			return
		}

		infoName = filepath.Base(event.Name)

		switch infoName {
		case "type":
			newBacklightInfo.typ, err = PathReadStr(filepath.Join(BacklightPath, basepath, "type"))
		default:
			*newBacklightInfo.mapIntPtrs()[infoName], err = PathReadInt(filepath.Join(BacklightPath, basepath, infoName))
		}

		if err != nil {
			errChan <- err

			return
		}

		backlightInfoChan <- newBacklightInfo
		backlightInfo = newBacklightInfo
	}
}

// BacklightInfoChan returns map of channels that sends backlight information
// for each backlight found in [BacklightPath] + glob.
func BacklightInfoChan(glob string) (map[string]<-chan *BacklightInfo, <-chan error, error) {
	var (
		backlightInfoChans map[string]<-chan *BacklightInfo
		errChan            chan error
		backlightPaths     []string
		path               string
		err                error
	)

	backlightInfoChans = make(map[string]<-chan *BacklightInfo)
	errChan = make(chan error)

	backlightPaths, err = filepath.Glob(filepath.Join(BacklightPath, glob))
	if err != nil {
		return nil, nil, err
	}

	for _, path = range backlightPaths {
		go serveBacklight(backlightInfoChans, errChan, filepath.Base(path))
	}

	return backlightInfoChans, errChan, nil
}

// Backlight returns backlight information in
// [BacklightPath] + basepath.
func Backlight(basepath string) (*BacklightInfo, error) {
	var (
		backlightInfo *BacklightInfo
		key           string
		value         *int
		err           error
	)

	backlightInfo = new(BacklightInfo)

	for key, value = range backlightInfo.mapIntPtrs() {
		*value, err = PathReadInt(filepath.Join(BacklightPath, basepath, key))
		if err != nil {
			return nil, err
		}
	}

	backlightInfo.typ, err = PathReadStr(filepath.Join(BacklightPath, basepath, "type"))
	if err != nil {
		return nil, err
	}

	return backlightInfo, err
}

// Backlights returns all backlight information in
// [BacklightPath] + glob.
func Backlights(glob string) ([]*BacklightInfo, error) {
	var (
		backlightPaths []string
		backlightInfos []*BacklightInfo
		idx            int
		err            error
	)

	backlightPaths, err = filepath.Glob(filepath.Join(BacklightPath, glob))
	if err != nil {
		return nil, err
	}

	backlightInfos = make([]*BacklightInfo, len(backlightPaths))

	for idx = range backlightPaths {
		backlightInfos[idx], err = Backlight(filepath.Base(backlightPaths[idx]))
		if err != nil {
			return nil, err
		}
	}

	return backlightInfos, nil
}
