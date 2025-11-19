package lib

import (
	"strings"

	"github.com/google/uuid"
)

func UUID(hypen bool) string {
	uuidWithHypen := uuid.New()
	uuid := strings.Replace(uuidWithHypen.String(), "-", "", -1)
	if hypen {
		return uuidWithHypen.String()
	} else {
		return uuid
	}
}
