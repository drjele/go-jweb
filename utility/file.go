package utility

import (
    `bytes`
    `encoding/csv`
    `os`

    jweberror `gitlab.com/drjele-go/jweb/error`
)

func Exists(filename string) bool {
    info, err := os.Stat(filename)

    if os.IsNotExist(err) {
        return false
    }

    return !info.IsDir()
}

func WriteCsv(rows [][]string) bytes.Buffer {
    buffer := bytes.Buffer{}
    writer := csv.NewWriter(&buffer)

    for _, row := range rows {
        err := writer.Write(row)

        jweberror.Panic(err)
    }

    writer.Flush()
    if err := writer.Error(); err != nil {
        jweberror.Panic(err)
    }

    return buffer
}
