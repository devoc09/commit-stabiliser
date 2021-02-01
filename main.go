package main

import (
	"fmt"
	"os/exec"
	"strings"
	"unicode/utf8"

	"github.com/manifoldco/promptui"
)

type committype struct {
	Name        string
	Description string
}

func main() {
	var coh string
	cohp := &coh
	ctypes := []committype{
		{Name: "build", Description: "Changes that affect the build system or external dependencies."},
		{Name: "ci", Description: "Changes to our CI configuration files and scripts."},
		{Name: "docs", Description: "Documentation only changes"},
		{Name: "feat", Description: "A new feature"},
		{Name: "fix", Description: "A bug fix"},
		{Name: "perf", Description: "A code change that improves performance."},
		{Name: "refactor", Description: "A code change that neither fixes a bug nor adds a feature."},
		{Name: "style", Description: "Change that do not affect the meaning of the code."},
		{Name: "test", Description: "Adding missing tests or correcting existing tests."},
		{Name: "BREAKING CHANGE", Description: "Introduces a breaking API change."},
	}

	stmp := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U0001F308 {{ .Name | cyan }} ({{ .Description | red }})",
		Inactive: "{{ .Name | cyan }} ({{ .Description | red }})",
		Selected: "{{ .Name | bold }} ({{ .Description | bold }})",
		Details: `
        --------- CommitType ----------
        {{ "Name:" | faint }} {{ .Name }}
        {{ "Description:" | faint }} {{ .Description }}`,
	}

	ptmp := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | cyan }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	searcher := func(input string, index int) bool {
		ctype := ctypes[index]
		name := strings.Replace(strings.ToLower(ctype.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	tp := promptui.Select{
		Label:     "Select the type of change that you're committing.",
		Items:     ctypes,
		Templates: stmp,
		Size:      5,
		Searcher:  searcher,
		IsVimMode: true,
		HideHelp:  true,
	}

	i, _, err := tp.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	scp := promptui.Prompt{
		Label:     "What is the scope of this change (e.g. component or file name):(press enter to skip)  ",
		Templates: ptmp,
	}

	scpr, err := scp.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	smp := promptui.Prompt{
		Label:     "Write a short, imperative tense description of the change (max 84 chars)  ",
		Templates: ptmp,
	}

	smpr, err := smp.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if utf8.RuneCountInString(smpr) >= 85 {
		fmt.Printf("This description is limited to 84 characters(Your description is %d characters).", utf8.RuneCountInString(smpr))
		return
	}

	bop := promptui.Prompt{
		Label:     "Provide a longer description of the change: (press enter to skip)  ",
		Templates: ptmp,
	}

	bopr, err := bop.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch {
	case scpr == "":
		*cohp = fmt.Sprintf("%s: ", ctypes[i].Name) + smpr
		fmt.Println(coh)
	default:
		*cohp = fmt.Sprintf("%s", ctypes[i].Name) + fmt.Sprintf("(%s): ", scpr) + smpr
		fmt.Println(coh)
	}

	switch {
	case bopr == "":
		commit, _ := exec.Command("git", "commit", "-m", coh).CombinedOutput()
		fmt.Println(string(commit))
	default:
		commit, _ := exec.Command("git", "commit", "-m", coh, "-m", bopr).CombinedOutput()
		fmt.Println(string(commit))
	}
}
