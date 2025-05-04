package prompts

import "fmt"

type Prompt struct {
	ID           string `json:"id"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	SystemPrompt string `json:"systemPrompt"`
	UserPrompt   string `json:"userPrompt"`
	Default      bool   `json:"default"`
}

type Manager struct {
	defaultIndex map[string]Prompt
	versionIndex map[string]map[string]Prompt
}

func New(prompts []Prompt) (*Manager, error) {
	m := &Manager{
		defaultIndex: make(map[string]Prompt),
		versionIndex: make(map[string]map[string]Prompt),
	}

	for _, p := range prompts {
		if p.ID == "" || p.Version == "" {
			return nil, fmt.Errorf("prompt id and version is empty")
		}

		if _, ok := m.versionIndex[p.ID][p.Version]; ok {
			return nil, fmt.Errorf("prompt with ID %s and version %s already exists", p.ID, p.Version)
		}

		if p.Default {
			if _, ok := m.defaultIndex[p.ID]; ok {
				return nil, fmt.Errorf("default prompt with ID %s already exists", p.ID)
			}

			m.defaultIndex[p.ID] = p
		}

		if m.versionIndex[p.ID] == nil {
			m.versionIndex[p.ID] = make(map[string]Prompt)
		}

		m.versionIndex[p.ID][p.Version] = p
	}

	for _, p := range prompts {
		if _, ok := m.defaultIndex[p.ID]; !ok {
			return nil, fmt.Errorf("default prompt with ID %s not found", p.ID)
		}
	}

	return m, nil
}

func (m *Manager) Get(id string, version ...string) (*Prompt, error) {
	if len(version) == 0 {
		p, ok := m.defaultIndex[id]
		if !ok {
			return nil, fmt.Errorf("default prompt with ID %s not found", id)
		}

		return &p, nil
	}

	vp, ok := m.versionIndex[id]
	if !ok {
		return nil, fmt.Errorf("prompt with ID %s not found", id)
	}

	p, ok := vp[version[0]]
	if !ok {
		return nil, fmt.Errorf("prompt with ID %s and version %s not found", id, version)
	}

	return &p, nil
}
