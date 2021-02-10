package utility

import (
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

func MapInterfaceToString(valueToConvert map[string]interface{}) (convertedValue map[string]string, err error) {
    convertedValue = map[string]string{}

    for key, value := range valueToConvert {
        stringValue, ok := value.(string)

        if ok == false {
            err = jweberror.New(`could not convert interface to string fo "%v"`, key)
        }

        convertedValue[key] = stringValue
    }

    return
}
