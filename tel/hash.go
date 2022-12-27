package tel

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Hash(tel string) string {
	salt := tel[:6]

	w := md5.New()
	_, _ = io.WriteString(w, tel+salt)
	return fmt.Sprintf("%x", w.Sum(nil))
}
