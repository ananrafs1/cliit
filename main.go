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

		pluginSelected, err := SelectPlugin()
		if err != nil {
			log.Fatalf(err.Error())
		}

		action, params := GenerateParameter(pluginSelected.GetActionMetadata())

		if action == "quit" {
			Exec()
			return
		}

		msg := pluginSelected.Execute(action, params)

		spinner.New().Title("Execute ...").Action(func() {
			for m := range msg {
				fmt.Println(m)
			}
		}).Run()

		time.Sleep(2 * time.Second)

		retryForm := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Are you want to try Again ?").
					Value(&retry),
			),
		)

		err = retryForm.Run()
		if err != nil {
			log.Fatalf("error on Retry form")
		}

	}

}

func SelectPlugin() (plugin.Executor, error) {
	pluginList := internal.GetPluginList()
	pluginOptionList := []huh.Option[string]{}
	for _, v := range pluginList {
		metadata := v.GetMetadata()
		pluginOptionList = append(pluginOptionList,
			huh.NewOption(fmt.Sprintf("[%s] %s", metadata.Title, metadata.Subtitle), metadata.Title),
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

	pluginSelected, ok := pluginList[*pluginSelectedTitle]
	if !ok {
		return nil, fmt.Errorf("err no plugin found")
	}

	return pluginSelected, nil
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
