package asana_test

import (
	"github.com/HirokazuMiyaji/asana-go"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	endpoint := asana.GetEndpoint()
	if endpoint != "https://app.asana.com/api/1.0" {
		t.Errorf("unexpected result: %s", endpoint)
	}
}

func TestGetApiKey(t *testing.T) {
	d, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}

	defer func() {
		if err := os.RemoveAll(d); err != nil {
			t.Error(err)
		}
	}()

	if err := os.Setenv("HOME", d); err != nil {
		t.Error(err)
	}

	apiKey, err := asana.GetApiKey()
	if err == nil {
		t.Error("asana.GetApiKey not occur error")
	}

	os.Setenv("ASANA_API_KEY", "FcZ23.M4xsMtXmTKmDA4ssLCkEnYi")
	apiKey, err = asana.GetApiKey()
	if err != nil {
		t.Error(err)
	}

	if apiKey != "RmNaMjMuTTR4c010WG1US21EQTRzc0xDa0VuWWk6" {
		t.Errorf("unexpected result: %s", apiKey)
	}

	os.Setenv("ASANA_API_KEY", "")

	data := `---
app.asana.com:
  apikey: FcZ23.M4xsMtXmTKmDA4ssLCkEnYi
`

	if err := ioutil.WriteFile(path.Join(d, ".pit", "default.yaml"), []byte(data), 0644); err != nil {
		t.Error(err)
	}

	apiKey, err = asana.GetApiKey()
	if err != nil {
		t.Error(err)
	}

	if apiKey != "RmNaMjMuTTR4c010WG1US21EQTRzc0xDa0VuWWk6" {
		t.Errorf("unexpected result: %s", apiKey)
	}
}
