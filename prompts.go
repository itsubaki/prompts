package prompts

import (
	"bytes"
	"fmt"
	"text/template"
)

type Prompt struct {
	ID           string `json:"id"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	SystemPrompt string `json:"systemPrompt"`
	UserPrompt   string `json:"userPrompt"`
	IsDefault    bool   `json:"is_default"`
	Template     *Template
}

func (p Prompt) With(template *Template) Prompt {
	return Prompt{
		ID:           p.ID,
		Version:      p.Version,
		Description:  p.Description,
		SystemPrompt: p.SystemPrompt,
		UserPrompt:   p.UserPrompt,
		IsDefault:    p.IsDefault,
		Template:     template,
	}
}

type Template struct {
	SystemPrompt *template.Template
	UserPrompt   *template.Template
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

	parsed := make([]Prompt, 0, len(prompts))
	for _, p := range prompts {
		if p.ID == "" || p.Version == "" {
			return nil, fmt.Errorf("prompt id or version is empty")
		}

		system, err := template.New("system_prompt").Parse(p.SystemPrompt)
		if err != nil {
			return nil, fmt.Errorf("new system prompt template: %w", err)
		}

		user, err := template.New("user_prompt").Parse(p.UserPrompt)
		if err != nil {
			return nil, fmt.Errorf("new user prompt template: %w", err)
		}

		parsed = append(parsed, p.With(&Template{
			SystemPrompt: system,
			UserPrompt:   user,
		}))
	}

	for _, p := range parsed {
		if m.versionIndex[p.ID] == nil {
			m.versionIndex[p.ID] = make(map[string]Prompt)
		}

		if _, ok := m.versionIndex[p.ID][p.Version]; ok {
			return nil, fmt.Errorf("prompt with ID %s and version %s already exists", p.ID, p.Version)
		}

		if p.IsDefault {
			if _, ok := m.defaultIndex[p.ID]; ok {
				return nil, fmt.Errorf("default prompt with ID %s already exists", p.ID)
			}

			m.defaultIndex[p.ID] = p
		}

		m.versionIndex[p.ID][p.Version] = p
	}

	for _, p := range parsed {
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

func Render(p *Prompt, data any) (*Prompt, error) {
	render := func(tmpl *template.Template) (string, error) {
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return "", fmt.Errorf("execute template: %w", err)
		}

		return buf.String(), nil
	}

	system, err := render(p.Template.SystemPrompt)
	if err != nil {
		return nil, fmt.Errorf("render system prompt: %w", err)
	}

	user, err := render(p.Template.UserPrompt)
	if err != nil {
		return nil, fmt.Errorf("render user prompt: %w", err)
	}

	return &Prompt{
		ID:           p.ID,
		Version:      p.Version,
		Description:  p.Description,
		SystemPrompt: system,
		UserPrompt:   user,
		IsDefault:    p.IsDefault,
		Template:     p.Template,
	}, nil
}
