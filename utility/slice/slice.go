package jwebslice

func StringInSlice(needle string, haystack []string) bool {
    for _, x := range haystack {
        if x == needle {
            return true
        }
    }

    return false
}
