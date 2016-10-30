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

func GetSnapshot(url string) (common.Snapshot, error) {
    bytes, err := GetBytes(url)
    if err != nil {
        return nil, err
    }
    patan1 := &patan_1_0_0.Snapshot{}
    err = json.Unmarshal(bytes, patan1)
    if err == nil {
        fmt.Println("received model version 1.0.0")
        return patan1, nil
    }

    patan3 := &patan_3_0_0.Snapshot{}
    err = json.Unmarshal(bytes, patan3)
    if err == nil {
        fmt.Println("received model version 3.0.0")
        return patan3, nil
    }

    return nil, errors.New("could not unmarshal to model 1.0.0 or 3.0.0")
}

