package zoneParser

import (
	"fmt"
	"os"

	"allypost.net/binder/app/zone"
	"github.com/bwesterb/go-zonefile"
)

func Parse(zoneFilePath string) (*zone.Zone, error) {
	data, err := os.ReadFile(zoneFilePath)
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %s", zoneFilePath, err)
	}

	zf, perr := zonefile.Load(data)
	if perr != nil {
		return nil, fmt.Errorf("\"%s:%d\" %s", zoneFilePath, perr.LineNo(), perr)
	}

	return &zone.Zone{
		Zone: zf,
	}, nil
}
