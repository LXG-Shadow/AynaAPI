package core

import (
	"encoding/hex"
	"encoding/json"
)

type ProviderMeta struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *ProviderMeta) Load(data string) error {
	bdata, err := hex.DecodeString(data)
	if err != nil {
		return nil
	}
	return json.Unmarshal(bdata, p)
}

func (p *ProviderMeta) Dump() string {
	data, _ := json.Marshal(p)
	return hex.EncodeToString(data)
}

func (p *ProviderMeta) GetCompletionStatus() bool {
	return p.Url != ""
}

type ProviderManager struct {
	ProviderMap map[string]interface{}
}

func (m *ProviderManager) Add(name string, provider interface{}) {
	m.ProviderMap[name] = provider
}

func (m *ProviderManager) AddMap(providers map[string]interface{}) {
	for key, val := range providers {
		m.ProviderMap[key] = val
	}
}

func (m *ProviderManager) GetProviderList() []string {
	plist := make([]string, 0)
	for key, _ := range m.ProviderMap {
		plist = append(plist, key)
	}
	return plist
}

func (m *ProviderManager) IsProviderAvailable(identifier string) bool {
	val, _ := m.ProviderMap[identifier]
	return val != nil
}

func (p *ProviderManager) GetProvider(identifier string) interface{} {
	val, _ := p.ProviderMap[identifier]
	if val == nil {
		return nil
	}
	return val
}
