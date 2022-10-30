package zone

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bwesterb/go-zonefile"
)

type Zone struct {
	Zone *zonefile.Zonefile
}

func (zf *Zone) findSoaEntry() (*zonefile.Entry, error) {
	for _, e := range zf.Zone.Entries() {
		if !bytes.Equal(e.Type(), []byte("SOA")) {
			continue
		}

		vs := e.Values()
		if len(vs) != 7 {
			return nil, errors.New("wrong number of parameters to SOA line")
		}

		return &e, nil
	}

	return nil, errors.New("could not find SOA entry")
}

func (zf *Zone) IncrementSoaRecord() error {
	entry, err := zf.findSoaEntry()

	if err != nil {
		return err
	}

	serialIndex := 2

	oldSerial, err := strconv.Atoi(string(entry.Values()[serialIndex]))
	if err != nil {
		return fmt.Errorf("could not parse existing serial: %+v", err)
	}

	newSerial, err := strconv.Atoi(time.Now().Format("20060102") + "00")
	if err != nil {
		return fmt.Errorf("could not generate new serial: %+v", err)
	}

	if newSerial <= oldSerial {
		newSerial = oldSerial + 1
	}

	err = entry.SetValue(serialIndex, []byte(strconv.Itoa(newSerial)))
	if err != nil {
		return fmt.Errorf("could not set value: %+v", err)
	}

	return nil
}

func (zf *Zone) Save(path string) error {
	zoneFile, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("could not open zone file: %s", err)
	}

	_, err = zoneFile.Write(zf.Zone.Save())
	if err != nil {
		return fmt.Errorf("could not write zone file: %s", err)
	}

	return nil
}
