package bitwarden

import (
	"fmt"

	"github.com/sganon/env-secrets/common"
	log "github.com/sirupsen/logrus"
)

const (
	bwBinary = "bw"
)

// BW wraps utility methods to interacts wit bw CLI
type BW struct {
	foldersIDs []FolderID
	values     map[string]string
}

// SetFoldersIDs fetch and set folder id by getting it by name
// If more than one folder match with folderName the first one will be choosed
func (bw *BW) SetFoldersIDs(foldersNames []string) error {
	for _, name := range foldersNames {
		var results []FolderSearch
		err := common.ExecCLI(bwBinary, []string{
			"list", "folders", "--search", name,
		}, &results)
		if err != nil {
			return fmt.Errorf("env-secrets error: an error occured while setting folderIDs: %v", err)
		}
		if len(results) == 0 {
			return fmt.Errorf("env-secrets error: unable to match folder with name %s", name)
		}
		bw.foldersIDs = append(bw.foldersIDs, results[0].ID)
		log.Debugf("Folder ID found: %s\n", results[0].ID)
	}
	return nil
}

// FetchItems get all items from folder and generates according Values
// SetFolderIF should have been called before end
func (bw *BW) FetchItems() error {
	if len(bw.foldersIDs) == 0 {
		panic("unable to fetch from unknow folder, please set foldersIDs before")
	}
	for _, id := range bw.foldersIDs {
		var results []Item
		err := common.ExecCLI(bwBinary, []string{"list", "items", "--folderid", string(id)}, &results)
		if err != nil {
			return fmt.Errorf("env-secrets error: an error occured fetching items in folder with id %s: %v", id, err)
		}
		if bw.values == nil {
			bw.values = make(map[string]string)
		}
		for _, item := range results {
			bw.values[item.Name] = item.Notes
		}
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
