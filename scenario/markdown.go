package scenario

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"github.com/go-git/go-billy/v5"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"strings"
)

// process markdown files and turn them into scenario steps

func RenderContent(fs billy.Filesystem) (*hfv1.Scenario, error) {
	// first get and parse the scenario file
	scenarioFile, err := fs.Open("scenario.yaml")
	if err != nil {
		return nil, err
	}
	scenario, err := readScenario(scenarioFile)
	if err != nil {
		return nil, err
	}

	// from here, open each markdown file, parse it into a step, and add that to the scenario
	// then return the scenario
	files, err := fs.ReadDir("")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !strings.Contains(".md", f.Name()) {
			continue // not a markdown file
		}

		stepFile, err := fs.Open(f.Name())
		if err != nil {
			return nil, err
		}

		step, err := renderStep(stepFile)
		if err != nil {
			return nil, err
		}

		scenario.Spec.Steps = append(scenario.Spec.Steps, *step)
	}

	return scenario, nil
}

func renderStep(f billy.File) (*hfv1.ScenarioStep, error) {
	var step = &hfv1.ScenarioStep{}
	reader := bufio.NewReader(f)

	// step name will be the first line of the file
	stepName, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	step.Title = base64.StdEncoding.EncodeToString([]byte(stepName))

	// step content will be everything after the first line of the file

	var stepContent []byte
	_, err = reader.Read(stepContent)
	if err != nil {
		return nil, err
	}

	step.Content = base64.StdEncoding.EncodeToString([]byte(stepContent))

	return step, nil
}

func readScenario(f billy.File) (*hfv1.Scenario, error) {
	decoder := json.NewDecoder(f)

	scenario := &hfv1.Scenario{}

	if err := decoder.Decode(scenario); err != nil {
		return nil, err
	}

	return scenario, nil
}
