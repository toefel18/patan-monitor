package remote

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/toefel18/patan-monitor/remote/patan-1.0.0"
    "github.com/toefel18/patan-monitor/remote/common"
    "fmt"
    "github.com/toefel18/patan-monitor/remote/patan-3.0.0"
    "errors"
)

//use get bytes and try all models
func GetJson(url string, targetDataModel interface{}) error {
	resp, err := http.Get(url)
    if err != nil {
		return err
	}
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    if err = json.Unmarshal(body, targetDataModel); err != nil {
        return err
    }
    return nil
}

func GetBytes(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return []byte{}, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return []byte{}, err
    }
    return body, nil
}

// Retrieves the statistics URL and tries unmarshall it into each of the patan models, and stops at the first match
func GetSnapshot(url string) (common.Snapshot, error) {
    bytes, err := GetBytes(url)
    if err != nil {
        return nil, err
    }
    if snapshot1, err := tryUnmarshall(bytes, &patan_1_0_0.Snapshot{}); err == nil {
        return snapshot1, nil;
    }
    if snapshot3, err := tryUnmarshall(bytes, &patan_3_0_0.Snapshot{}); err == nil {
        return snapshot3, nil;
    }
    return nil, errors.New("could not unmarshal to any of the configured models")
}

func tryUnmarshall(bytes []byte, target common.Snapshot) (common.Snapshot, error) {
    err := json.Unmarshal(bytes, target)
    if err == nil {
        fmt.Println("received model version 3.0.0")
        return target, nil
    }
    return nil, err
}
