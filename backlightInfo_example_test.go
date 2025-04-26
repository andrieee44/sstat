package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Print the brightness percentage of intel_backlight.
func ExampleBacklight() {
	var (
		backlightInfo *sstat.BacklightInfo
		perc          float64
		err           error
	)

	backlightInfo, err = sstat.Backlight("intel_backlight")
	if err != nil {
		panic(err)
	}

	perc = float64(backlightInfo.Brightness()) / float64(backlightInfo.MaxBrightness()) * 100
	fmt.Printf("Brightness: %g%%\n", perc)
}

// Print the brightness percentage of each backlight.
func ExampleBacklights() {
	var (
		backlightInfos []*sstat.BacklightInfo
		perc           float64
		idx            int
		err            error
	)

	backlightInfos, err = sstat.Backlights("*")
	if err != nil {
		panic(err)
	}

	for idx = range backlightInfos {
		perc = float64(backlightInfos[idx].Brightness()) / float64(backlightInfos[idx].MaxBrightness()) * 100
		fmt.Printf("%s Brightness: %g%%\n", backlightInfos[idx].Name(), perc)
	}
}

// Print the brightness percentage of intel_backlight at every change.
func ExampleBacklightChans() {
	var (
		backlightChans map[string]<-chan *sstat.BacklightInfo
		backlightInfo  *sstat.BacklightInfo
		errChan        <-chan error
		perc           float64
		err            error
	)

	backlightChans, errChan, err = sstat.BacklightChans("*")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case backlightInfo = <-backlightChans["intel_backlight"]:
			perc = float64(backlightInfo.Brightness()) / float64(backlightInfo.MaxBrightness()) * 100
			fmt.Printf("Brightness: %g%%\n", perc)
		case err = <-errChan:
			panic(err)
		}
	}
}
