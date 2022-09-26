package config

import (
	"encoding/json"
	"testing"
)

func TestInitConfig(t *testing.T) {
	config, err := InitConfig("../")
	if err != nil {
		t.Errorf("InitConfig err:%v", err)
	}
	marshal, err := json.Marshal(config)
	if err != nil {
		t.Errorf("InitConfig Marshal config err:%v", err)
	}
	t.Logf("load config %s", string(marshal))
}

