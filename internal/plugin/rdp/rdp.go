package rdp

import (
	"fmt"
	"github.com/chainreactors/zombie/pkg"
	rdpclient "github.com/tomatome/grdp/client"
)

type RdpPlugin struct {
	*pkg.Task
	conn *rdpclient.Client
}

func (s *RdpPlugin) Unauth() (bool, error) {
	return false, pkg.NotImplUnauthorized
}

func (s *RdpPlugin) Login() error {
	client := rdpclient.NewClient(s.Address(), s.Username, s.Password, rdpclient.TC_RDP, nil)
	if client == nil {
		return fmt.Errorf("init error")
	}
	err := client.Login()
	if err != nil {
		return err
	}
	readyChan := make(chan bool)
	defer close(readyChan)
	client.OnReady(func() {
		readyChan <- true
	})
	<-readyChan
	s.conn = client
	return nil
}

func (s *RdpPlugin) Close() error {
	return nil
}

func (s *RdpPlugin) Name() string {
	return s.Service
}

func (s *RdpPlugin) GetResult() *pkg.Result {
	// todo list dbs
	return &pkg.Result{Task: s.Task, OK: true}
}
