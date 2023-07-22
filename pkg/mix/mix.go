package vite

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

type Entry struct {
	File string `json:"file"`
	Src  string `json:"src"`

	Assets []string `json:"assets"`
	CSS    []string `json:"css"`

	IsEntry bool `json:"isEntry"`
}
type Manifest map[string]Entry

// Mix type.
type Mix struct {
	url         string
	publicPath  string
	hotProxyURL string
	manifests   map[string]Manifest
}

// New function.
func xNew(url, publicPath, hotProxyURL string) *Mix {
	m := new(Mix)
	m.url = url
	m.publicPath = publicPath
	m.hotProxyURL = hotProxyURL
	m.manifests = make(map[string]Manifest)

	return m
}

// Mix function.
func (m *Mix) Mix(path, manifestDirectory string) (string, error) {
	path = m.pathPrefix(path)
	rPath := strings.TrimPrefix(path, "/")
	manifestDirectory = m.pathPrefix(manifestDirectory)
	log.Printf("path %v", m.publicPath+"/hot")
	_, err := os.Stat(m.publicPath + "/hot")
	log.Printf("errrsh %v", err)
	if err == nil {
		log.Printf("errr %v", err)

		if m.hotProxyURL != "" {
			return m.hotProxyURL + path, nil
		}

		content, err := os.ReadFile(m.publicPath + "/hot")
		if err != nil {
			return "", err
		}
		log.Printf("content %v", content)
		url := strings.TrimSpace(string(content))

		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			return url[strings.Index(url, ":")+1:] + path, nil
		}

		return "//localhost:8080" + path, nil
	}

	manifestPath := m.publicPath + m.manifestPath(manifestDirectory)
	log.Printf("manifestPath %v", manifestPath)
	if _, ok := m.manifests[manifestPath]; !ok {
		_, err := os.Stat(manifestPath)
		if os.IsNotExist(err) {
			return "", ErrManifestNotExist
		}

		content, err := os.ReadFile(manifestPath)
		if err != nil {
			return "", err
		}

		// var data map[string]string

		var data Manifest
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Printf("err %v", err.Error())
			return "", err
		}

		m.manifests[manifestPath] = data
		// b, err := json.MarshalIndent(m.manifests[manifestPath]["src/main.ts"], " ", "  ")
		// log.Printf("manifests %v", string(b))
		// log.Printf("file %v", m.manifests[manifestPath]["src/main.ts"].File)
		// log.Printf("filePath %v", m.manifests[manifestPath][rPath].File)

	}

	manifest := m.manifests[manifestPath]
	// b, err := json.MarshalIndent(manifest[rPath], " ", "  ")
	// log.Printf("manifestOut %v", string(b))

	if _, ok := manifest[rPath]; !ok {
		log.Printf("mix: unable to locate mix file: %v", path)
		return "", fmt.Errorf("mix: unable to locate mix file: %v", path)
	}
	log.Printf("manifestFile %v", m.url+manifestDirectory+m.pathPrefix(manifest[rPath].File))

	return m.url + manifestDirectory + m.pathPrefix(manifest[rPath].File), nil
}

// Hash function.
func (m *Mix) Hash(manifestDirectory string) (string, error) {
	manifestPath := m.publicPath + m.manifestPath(m.pathPrefix(manifestDirectory))

	_, err := os.Stat(manifestPath)
	if os.IsNotExist(err) {
		return "", ErrManifestNotExist
	}
	file, err := os.Open(manifestPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	log.Println("mix file")
	log.Println(m.hashFromFile(file))
	return m.hashFromFile(file)
}

// HashFromFS function.
func (m *Mix) HashFromFS(manifestDirectory string, staticFS fs.FS) (string, error) {
	log.Println(strings.TrimPrefix(m.manifestPath(manifestDirectory), "/"))
	file, err := staticFS.Open(strings.TrimPrefix(m.manifestPath(manifestDirectory), "/"))
	if err != nil {
		return "", err
	}

	defer file.Close()

	return m.hashFromFile(file)
}

func (m *Mix) hashFromFile(file io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (m *Mix) manifestPath(manifestDirectory string) string {
	return manifestDirectory + "/manifest.json"
}

func (m *Mix) pathPrefix(path string) string {
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return path
}
