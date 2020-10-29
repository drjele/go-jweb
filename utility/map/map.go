package jwebmap

import (
    jweberror `gitlab.com/drjele-go/jweb/error`
)

func CheckKeysMatch(keys []string, mapToCheck map[string]interface{}) error {
    for _, key := range keys {
        _, ok := mapToCheck[key]

        if ok == false {
            return jweberror.New(`the key "%v" is missing`, key)
        }

        delete(mapToCheck, key)
    }

    if len(mapToCheck) > 0 {
        return jweberror.New(`extra keys were found: %v`, GetKeys(mapToCheck))
    }

    return nil
}

func GetKeys(m map[string]interface{}) []string {
    keys := make([]string, 0, len(m))

    for k := range m {
        keys = append(keys, k)
    }

    return keys
}
