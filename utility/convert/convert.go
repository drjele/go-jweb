package jwebconvert

import (
    `fmt`
    `os`
    `strconv`

    jweberror `gitlab.com/drjele-go/jweb/error`
)

func StringToInt(valueToConvert string) int {
    convertedValue, err := strconv.Atoi(valueToConvert)
    jweberror.Panic(err)

    return convertedValue
}

func InterfaceToMap(valueToConvert interface{}) (convertedValue map[string]interface{}, err error) {
    convertedValue, ok := valueToConvert.(map[string]interface{})
    if ok == false {
        err = jweberror.New(`could not convert interface to map`)
    }

    return
}

func InterfaceToMapString(valueToConvert interface{}) (convertedValue map[string]string, err error) {
    fmt.Println(valueToConvert)
    os.Exit(12)

    convertedValue, ok := valueToConvert.(map[string]string)
    if ok == false {
        err = jweberror.New(`could not convert interface to map`)
    }

    return
}
