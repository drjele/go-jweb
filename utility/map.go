package utility

import (
    jweberror `gitlab.com/drjele-go/jweb/error`
)

func CheckKeysMatch(keys []string, mapToCheck map[string]interface{}) error {
    mapToCheckKeys := map[string]interface{}{}

    for key, _ := range mapToCheck {
        mapToCheckKeys[key] = true
    }

    for _, key := range keys {
        _, ok := mapToCheckKeys[key]

        if ok == false {
            return jweberror.New(`the key "%v" is missing`, key)
        }

        delete(mapToCheckKeys, key)
    }

    if len(mapToCheckKeys) > 0 {
        return jweberror.New(`extra keys were found: %v`, GetKeys(mapToCheckKeys))
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
