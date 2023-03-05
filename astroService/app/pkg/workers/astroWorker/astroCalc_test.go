package astroWorker

import (
	"testing"
	"time"
)

func Test_CalculateAstro_WithDefautTimestamp_ReturnsStringSign(t *testing.T) {

	dateBirth := time.Now()

	sign := CalculateSign(dateBirth)

	t.Error(sign == "")

}
