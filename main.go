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

var (
	burger       string
	toppings     []string
	sauceLevel   int
	name         string
	instructions string
	discount     bool
)

func main() {
	Exec()
}

func Serve() {
	// form := huh.NewForm(
	// 	huh.NewGroup(
	// 		// Ask the user for a base burger and toppings.
	// 		huh.NewSelect[string]().
	// 			Title("Choose your burger").
	// 			Options(
	// 				huh.NewOption("Charmburger Classic", "classic"),
	// 				huh.NewOption("Chickwich", "chickwich"),
	// 				huh.NewOption("Fishburger", "fishburger"),
	// 				huh.NewOption("Charmpossible™ Burger", "charmpossible"),
	// 			).
	// 			Value(&burger), // store the chosen option in the "burger" variable

	// 		// Let the user select multiple toppings.
	// 		huh.NewMultiSelect[string]().
	// 			Title("Toppings").
	// 			Options(
	// 				huh.NewOption("Lettuce", "lettuce").Selected(true),
	// 				huh.NewOption("Tomatoes", "tomatoes").Selected(true),
	// 				huh.NewOption("Jalapeños", "jalapeños"),
	// 				huh.NewOption("Cheese", "cheese"),
	// 				huh.NewOption("Vegan Cheese", "vegan cheese"),
	// 				huh.NewOption("Nutella", "nutella"),
	// 			).
	// 			Limit(4). // there’s a 4 topping limit!
	// 			Value(&toppings),

	// 		// Option values in selects and multi selects can be any type you
	// 		// want. We’ve been recording strings above, but here we’ll store
	// 		// answers as integers. Note the generic "[int]" directive below.
	// 		huh.NewSelect[int]().
	// 			Title("How much Charm Sauce do you want?").
	// 			Options(
	// 				huh.NewOption("None", 0),
	// 				huh.NewOption("A little", 1),
	// 				huh.NewOption("A lot", 2),
	// 			).
	// 			Value(&sauceLevel),
	// 	),

	// 	// Gather some final details about the order.
	// 	huh.NewGroup(
	// 		huh.NewInput().
	// 			Title("What's your name?").
	// 			Value(&name).
	// 			// Validating fields is easy. The form will mark erroneous fields
	// 			// and display error messages accordingly.
	// 			Validate(func(str string) error {
	// 				if str == "Frank" {
	// 					return errors.New("Sorry, we don’t serve customers named Frank.")
	// 				}
	// 				return nil
	// 			}),
	// 		huh.NewText().
	// 			Title("Special Instructions").
	// 			CharLimit(400).
	// 			Value(&instructions),

	// 		huh.NewConfirm().
	// 			Title("Would you like 15% off?").
	// 			Value(&discount),
	// 	),
	// )

	// err := form.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if !discount {
	// 	fmt.Println("What? You didn’t take the discount?!")
	// }

	// fmt.Printf("Burger selected: %s \n", burger)
	// fmt.Printf("Toppings selected: %s \n", toppings)
	// fmt.Printf("Sauce Level selected: %d \n", sauceLevel)
	// fmt.Printf("Instructions selected: %s \n", instructions)

	fieldMp := map[string]string{
		"nama":       "What is your name ?",
		"alamat":     "What is your alamat ?",
		"usia":       "how old are you ?",
		"keterangan": "keterangan :",
	}
	fieldVal := map[string]*string{
		"nama":       new(string),
		"alamat":     new(string),
		"usia":       new(string),
		"keterangan": new(string),
	}
	var hFields []huh.Field
	for k, q := range fieldMp {
		hFields = append(hFields,
			huh.NewInput().
				Title(q).
				Value(fieldVal[k]),
		)
	}

	form3 := huh.NewForm(huh.NewGroup(hFields...))
	err := form3.Run()
	if err != nil {
		log.Fatal(err)
	}
	for k, q := range fieldVal {
		fmt.Printf("Instructions selected: %s on %s \n", k, *q)
	}

	var hFields2 []huh.Field
	for k, q := range fieldMp {
		hFields2 = append(hFields,
			huh.NewInput().
				Title(q).
				Value(fieldVal[k]),
		)
	}

	form4 := huh.NewForm(huh.NewGroup(hFields2...))
	err = form4.Run()
	if err != nil {
		log.Fatal(err)
	}
	for k, q := range fieldVal {
		fmt.Printf("Instructions selected: %s on %s \n", k, *q)
	}
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
