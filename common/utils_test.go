package common

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAttachmentFilename(t *testing.T) {
	str := "attachment; filename=zzopen-mysqldoc-e3cbfdf.tar.gz"
	filename, err := ParseAttachmentFilename(str)
	assert.Equal(t, nil, err)
	assert.Equal(t, "zzopen-mysqldoc-e3cbfdf.tar.gz", filename)
}

func TestRandomDelay(t *testing.T) {
	log.Println("=== START ===")
	RandomDelay(5)
	log.Println("=== END ===")
}
