package prompts_test

import (
	"fmt"

	"github.com/itsubaki/prompts"
)

func Example() {
	list := []prompts.Prompt{
		{
			ID:           "quantum_agent",
			Version:      "0.0.1",
			Description:  "Agent for Quantum Computation.",
			SystemPrompt: "You are a helpful agent who can answer user questions about the Quantum Computation.",
			UserPrompt:   "What is Quantum Computation?",
		},
		{
			ID:           "quantum_agent",
			Version:      "0.0.2",
			Description:  "Agent for Quantum Computation and Quantum Information.",
			SystemPrompt: "You are a helpful agent who can answer user questions about the Quantum Computation and Quantum Information.",
			UserPrompt:   "What is Quantum Computation and Quantum Information?",
			Default:      true,
		},
		{
			ID:           "google_saerch",
			Version:      "0.0.1",
			Description:  "Agent to answer questions using Google Search.",
			SystemPrompt: "You are a specialist in Google Search",
			UserPrompt:   "Search for Shor's algorithm and explain it to me.",
			Default:      true,
		},
	}

	manager, err := prompts.New(list)
	if err != nil {
		panic(err)
	}

	prompt, err := manager.Get("quantum_agent")
	if err != nil {
		panic(err)
	}

	fmt.Println(prompt.ID)
	fmt.Println(prompt.Version)
	fmt.Println(prompt.Description)
	fmt.Println(prompt.SystemPrompt)
	fmt.Println(prompt.UserPrompt)
	fmt.Println(prompt.Default)

	// Output:
	// quantum_agent
	// 0.0.2
	// Agent for Quantum Computation and Quantum Information.
	// You are a helpful agent who can answer user questions about the Quantum Computation and Quantum Information.
	// What is Quantum Computation and Quantum Information?
	// true
}

func ExampleManager_GetWith() {
	list := []prompts.Prompt{
		{
			ID:           "quantum_agent",
			Version:      "0.0.1",
			Description:  "Agent for Quantum Computation.",
			SystemPrompt: "You are a helpful agent who can answer user questions about the Quantum Computation.",
			UserPrompt:   "What is Quantum Computation?",
		},
		{
			ID:           "quantum_agent",
			Version:      "0.0.2",
			Description:  "Agent for Quantum Computation and Quantum Information.",
			SystemPrompt: "You are a helpful agent who can answer user questions about the Quantum Computation and Quantum Information.",
			UserPrompt:   "What is Quantum Computation and Quantum Information?",
			Default:      true,
		},
		{
			ID:           "google_saerch",
			Version:      "0.0.1",
			Description:  "Agent to answer questions using Google Search.",
			SystemPrompt: "You are a specialist in Google Search",
			UserPrompt:   "Search for Shor's algorithm and explain it to me.",
			Default:      true,
		},
	}

	manager, err := prompts.New(list)
	if err != nil {
		panic(err)
	}

	prompt, err := manager.Get("quantum_agent", "0.0.1")
	if err != nil {
		panic(err)
	}

	fmt.Println(prompt.ID)
	fmt.Println(prompt.Version)
	fmt.Println(prompt.Description)
	fmt.Println(prompt.SystemPrompt)
	fmt.Println(prompt.UserPrompt)
	fmt.Println(prompt.Default)

	// Output:
	// quantum_agent
	// 0.0.1
	// Agent for Quantum Computation.
	// You are a helpful agent who can answer user questions about the Quantum Computation.
	// What is Quantum Computation?
	// false
}
