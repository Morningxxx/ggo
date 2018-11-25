package strUtils

import "bytes"

func ConnectStrs(strings ...string) string {
    buffer := new(bytes.Buffer)
    for _, str := range strings {
        buffer.WriteString(str)
    }
    return buffer.String()
}

