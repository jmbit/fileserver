package session


import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var SessionStore sessions.FilesystemStore

func InitSessions() {
	SessionStore = sessionStore()

}

func sessionStore() sessions.FilesystemStore {
	var key []byte
	if keystring := viper.GetString("web.session.key"); keystring != "" {
		key = []byte(keystring)
	} else {
		key = securecookie.GenerateRandomKey(32)
	}
	sessionPath := viper.GetString("web.session.path")
	if path.IsAbs(sessionPath) == false {
		workdir, err := os.Getwd()
		if err != nil {
			log.Fatal("Error determining working directory", err)
		}
		sessionPath = path.Join(workdir, sessionPath)
	}

	log.Println("Session store set to ", sessionPath)
	return *sessions.NewFilesystemStore(sessionPath, key)
}

func serializeToJSON(slice []string) (string, error) {
	jsonData, err := json.Marshal(slice)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func deserializeFromJSON(data string) ([]string, error) {
	var slice []string
	err := json.Unmarshal([]byte(data), &slice)
	return slice, err
}
