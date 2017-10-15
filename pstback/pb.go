package pstback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
)

//BackUp backup
func BackUp(cfgpath string) error {

	d, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		return errors.Wrap(err, "")
	}

	var cfg Config
	err = json.Unmarshal(d, &cfg)
	if err != nil {
		return errors.Wrap(err, "")
	}

	t := time.Now()
	for _, item := range cfg.Items {
		err = backup(item, t)
		if err != nil {
			return errors.Wrap(err, "")
		}
	}

	return nil
}

func backup(item Item, t time.Time) error {
	var buff bytes.Buffer
	err := zip.Archive(item.Src, &buff, progress)
	if err != nil {
		return errors.Wrap(err, "")
	}

	name := outputPath(item, t)
	err = ioutil.WriteFile(name, buff.Bytes(), 0666)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func outputPath(item Item, t time.Time) string {
	ts := t.Format("20060102_150405")
	name := fmt.Sprintf("%s_%s.zip", item.SrcBase(), ts)
	return filepath.Join(item.Dest, name)
}

func progress(archivePath string) {
	fmt.Println(archivePath)
}
