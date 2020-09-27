package ipc

import (
	"testing"
)

func TestMasterStart(t *testing.T) {
	Accept()

}

func TestSlaveRegister(t *testing.T){
	go register()
}
