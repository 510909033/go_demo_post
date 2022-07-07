package discovery

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := GetLocalIP()
	assert.Nil(t, err)
	log.Println(ip, err)
	t.Log("GetLocalIP=", ip)
}
