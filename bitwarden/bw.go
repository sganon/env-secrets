package bitwarden

import (
	"encoding/json"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const (
	bwBinary = "bw"
)

// BW wraps utility methods to interacts wit bw CLI
type BW struct {
	folderID FolderID
	values   map[string]string
}

// SetFolderID fetch and set folder id by getting it by name
// If more than one folder match with folderName the first one will be choosed
func (bw *BW) SetFolderID(folderName string) error {
	cmd := exec.Command(bwBinary, "list", "folders", "--search", folderName)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	var results []FolderSearch
	if err := json.NewDecoder(stdout).Decode(&results); err != nil {
		return err
	}
	if len(results) == 0 {
		return fmt.Errorf("env-secrets error: unable to match folder with name %s", folderName)
	}
	bw.folderID = results[0].ID
	log.Debugf("Folder ID found: %s\n", bw.folderID)
	return nil
}

// FetchItems get all items from folder and generates according Values
// SetFolderIF should have been called before end
func (bw *BW) FetchItems() error {
	if bw.folderID == "" {
		panic("unable to fetch from unknow folder, please set folderID before")
	}
	var results []Item
	err := execBW([]string{"list", "items", "--folderid", string(bw.folderID)}, &results)
	if err != nil {
		return fmt.Errorf("env-secrets error: an error occured fetching items in folder with id %s: %v", bw.folderID, err)
	}
	if bw.values == nil {
		bw.values = make(map[string]string)
	}
	for _, item := range results {
		bw.values[item.Name] = item.Notes
	}
	return nil
}

func execBW(args []string, output interface{}) error {
	cmd := exec.Command(bwBinary, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	if err := json.NewDecoder(stdout).Decode(&output); err != nil {
		return err
	}
	return nil
}

// GenerateEnv implements common.EnvGenerator
func (bw BW) GenerateEnv() string {
	str := ""
	for key, val := range bw.values {
		str += fmt.Sprintf("export %s=%s\n", key, val)
	}
	return str
}
