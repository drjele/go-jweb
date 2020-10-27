package jwebconvert

import (
    `strconv`

    jweberror `gitlab.com/drjele-go/jweb/error`
)

func StringToInt(valueToConvert string) int {
    convertedValue, err := strconv.Atoi(valueToConvert)
    jweberror.Panic(err)

    return convertedValue
}
