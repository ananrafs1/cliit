package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ananrafs1/cliit/internal"
	"github.com/ananrafs1/cliit/plugin"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

func main() {
	Exec()
}

func Exec() {
	var (
		retry = true
	)
	for {
		if !retry {
			break
		}

		func(retry *bool) {
			pluginSelected, closer, err := SelectPlugin()
			if err != nil {
				log.Fatalf(err.Error())
			}
			defer closer()

			action, params := GenerateParameter(pluginSelected.GetActionMetadata())
			if action == "quit" {
				return
			}

			spinner.New().Title("Execute ...").Action(func() {
				ListenResponse(pluginSelected.Execute(action, params))
			}).Run()

			time.Sleep(2 * time.Second)

			createRetryForm(retry)

		}(&retry)

	}

}

func createRetryForm(retry *bool) {
	retryForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Are you want to try Again ?").
				Value(retry),
		),
	)

	err := retryForm.Run()
	if err != nil {
		log.Fatalf("error on Retry form")
	}
}

func SelectPlugin() (plugin.Executor, func(), error) {
	pluginList := internal.GetPluginList()
	pluginOptionList := []huh.Option[string]{}
	for _, fileName := range pluginList {
		pluginOptionList = append(pluginOptionList,
			huh.NewOption(fileName, fileName),
		)
	}

	pluginSelectedTitle := new(string)
	pluginForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick Plugin").
				Options(pluginOptionList...).
				Value(pluginSelectedTitle),
		),
	)

	err := pluginForm.Run()
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	return internal.AccessPlugin(*pluginSelectedTitle)
}

func ListenResponse(response <-chan string) {
	for m := range response {
		fmt.Println(m)
	}
}

func GenerateParameter(actionParamMap map[string]map[string]string) (string, map[string]string) {
	paramFormGroup := []*huh.Group{}
	formGroup := []*huh.Group{}

	var (
		selectedAction = new(string)
		paramsOnAction = map[string]map[string]*string{}
	)

	actionSelectList := []huh.Option[string]{}
	for k, _params := range actionParamMap {
		currentForm := k
		actionSelectList = append(actionSelectList,
			huh.NewOption(k, k),
		)
		paramForm := []huh.Field{}
		for val, q := range _params {
			if _, ok := paramsOnAction[k]; !ok {
				paramsOnAction[k] = map[string]*string{}
			}

			if _, ok := paramsOnAction[k][val]; !ok {
				paramsOnAction[k][val] = new(string)
			}

			paramForm = append(paramForm,
				huh.NewInput().
					Title(q).
					Value(paramsOnAction[k][val]).
					Description((fmt.Sprintf("[%s]", currentForm))),
			)
		}

		paramFormGroup = append(paramFormGroup,
			huh.NewGroup(
				paramForm...,
			).WithHideFunc(func() bool {
				return currentForm != *selectedAction
			}))
	}

	actionSelectList = append(actionSelectList,
		huh.NewOption[string]("Quit", "quit"),
	)

	formGroup = append(formGroup,
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Action to Execute").
				Options(actionSelectList...).
				Value(selectedAction),
		),
	)
	formGroup = append(formGroup,
		paramFormGroup...,
	)

	form := huh.NewForm(
		formGroup...,
	)

	err := form.Run()
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	__param := map[string]string{}
	for k, val := range paramsOnAction {
		if k == *selectedAction {
			for param_key, param_val := range val {
				__param[param_key] = *param_val
			}
		}
	}

	return *selectedAction, __param
}
