package main

import (
	"fmt"
	"os"

	"allypost.net/binder/app/version"
	"allypost.net/binder/app/zone"
	"allypost.net/binder/app/zone/zoneParser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s v%s\n\n", version.BuildProgramName(), version.BuildVersion())
		fmt.Println("Usage:\n\t", version.BuildProgramName(), "./path/to/zonefile1", "./path/to/zonefile2", "...")
		os.Exit(1)
	}

	zoneFilePaths := os.Args[1:]
	newZones := make(map[string]*zone.Zone)
	zoneErrors := make(map[string]error)

	for _, zoneFilePath := range zoneFilePaths {
		zone, err := zoneParser.Parse(zoneFilePath)
		if err != nil {
			zoneErrors[zoneFilePath] = err
			continue
		}

		err = zone.IncrementSoaRecord()
		if err != nil {
			zoneErrors[zoneFilePath] = err
			continue
		}

		newZones[zoneFilePath] = zone
	}

	if len(zoneErrors) > 0 {
		fmt.Println("Errors encounterd while processing zones.")
		fmt.Println("No zones were updated to save integrity.")
		fmt.Println("Errors:")
		for zoneFilePath, err := range zoneErrors {
			fmt.Printf("\t%s: %s\n", zoneFilePath, err)
		}
		os.Exit(1)
	}

	hadErrors := false
	for zoneFilePath, zone := range newZones {
		err := zone.Save(zoneFilePath)
		if err != nil {
			hadErrors = true
			fmt.Printf("Error saving zone `%s`: %s\n\n", zoneFilePath, err)
			continue
		}
	}

	if hadErrors {
		os.Exit(1)
	} else {
		fmt.Println("All zones updated successfully.")
	}
}
