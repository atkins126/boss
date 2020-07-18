package boss

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Masterminds/semver"
	"github.com/hashload/boss/internal/pkg/utils"
	"github.com/hashload/boss/internal/version"
	"github.com/snakeice/gogress"
)

const latestRelease = "https://api.github.com/repos/HashLoad/boss/releases/latest"
const releaseTag = "https://api.github.com/repos/HashLoad/boss/releases/tags/%s"
const tags = "https://api.github.com/repos/HashLoad/boss/tags"

func next(bar *gogress.Progress, message string) {
	bar.Inc()
	bar.Prefix(message)
	time.Sleep(500 * time.Millisecond)
}

func Upgrade(beta bool) {
	var link string
	var size float64
	var newVersion string

	fmt.Println("Finding new version...")
	if !beta {
		link, size, newVersion = getInfo(latestRelease)
	} else {
		tag := getLastBeta()
		link, size, newVersion = getInfo(fmt.Sprintf(releaseTag, tag))
	}

	if !checkVersion(newVersion, beta) {
		return
	}

	pool := gogress.NewPool()
	pool.RefreshRate = time.Second / 30
	pool.Start()

	root := pool.NewBar(5)

	ex, err := os.Executable()
	utils.CheckError(err)

	exePath, _ := filepath.Abs(ex)

	next(root, "Downloading new version...")
	downloadBar := pool.NewBar64(int64(size))
	downloadFile(exePath+".new", link, downloadBar)
	pool.RemoveBar(downloadBar)

	next(root, "Removing old backup...")
	_ = os.Remove(exePath + ".old")

	next(root, "Moving current to backup file...")
	if err := os.Rename(exePath, exePath+".old"); err != nil {
		fmt.Printf("[WARN] Failed on rename " + exePath + " to " + exePath + ".old")
	}

	next(root, "Replacing current version...")
	if err := os.Rename(exePath+".new", exePath); err != nil {
		utils.CheckError(fmt.Errorf("Failed on rename "+exePath+".new"+" to "+exePath, err.Error()))
	}
	next(root, "Done...")
	pool.FinishAll()
	fmt.Println("Run `boss version` to see new version")
	time.Sleep(1 * time.Second)

}

func getLastBeta() string {
	resp := doGet(tags)

	contents, err := ioutil.ReadAll(resp.Body)
	utils.CheckError(err)

	defer utils.CheckError(resp.Body.Close())

	var obj []interface{}
	err = json.Unmarshal(contents, &obj)
	utils.CheckError(err)

	tagObj := obj[0].(map[string]interface{})
	return tagObj["name"].(string)
}

func getInfo(url string) (string, float64, string) {
	resp := doGet(url)
	contents, err := ioutil.ReadAll(resp.Body)
	utils.CheckError(err)
	var obj interface{}

	err = json.Unmarshal(contents, &obj)
	utils.CheckError(err)

	defer utils.CheckError(resp.Body.Close())

	latest := obj.(map[string]interface{})
	version := latest["name"].(string)

	link := ""
	size := 0.0
	assets := latest["assets"].([]interface{})
	for _, assetRaw := range assets {
		asset := assetRaw.(map[string]interface{})
		if asset["name"].(string) == "boss.exe" {
			link = asset["browser_download_url"].(string)
			size = asset["size"].(float64)
		}
	}

	return link, size, version
}

func doGet(url string) *http.Response {
	resp, err := http.Get(url)
	utils.CheckError(err)

	if resp.StatusCode != http.StatusOK {
		utils.CheckError(fmt.Errorf("bad status: %s", resp.Status))
	}
	return resp
}

func downloadFile(filepath string, url string, bar *gogress.Progress) {
	_ = os.Remove(filepath)
	out, err := os.Create(filepath)
	if err != nil {
		utils.CheckError(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		utils.CheckError(err)
	}

	if resp.StatusCode != http.StatusOK {
		utils.CheckError(fmt.Errorf("bad status: %s", resp.Status))
	}

	proxyReader := bar.NewProxyReader(resp.Body)
	defer proxyReader.Close()

	_, err = io.Copy(out, proxyReader)
	utils.CheckError(err)

	utils.CheckError(out.Close())
	utils.CheckError(resp.Body.Close())
}

func checkVersion(newVersionString string, beta bool) bool {

	newVersion, _ := semver.NewVersion(newVersionString)
	current, _ := semver.NewVersion(version.Get().Version)

	fmt.Printf("Current: %s\n", current)
	fmt.Printf("Remote: %s\n", newVersion)

	needUpdate := newVersion.GreaterThan(current)

	if !needUpdate && beta {
		needUpdate = current.Prerelease() == "" && newVersion.Prerelease() != ""
	} else if !needUpdate && !beta {
		needUpdate = current.Prerelease() != "" && newVersion.Prerelease() == ""
	}

	if needUpdate {
		fmt.Printf("Updating %s to %s\n", current.String(), newVersion.String())
	} else {
		fmt.Println("already up to date!")
	}
	return needUpdate
}
