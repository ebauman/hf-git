package main

import (
	"github.com/rancher/wrangler/pkg/cleanup"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cleanup.Cleanup("./apis/"); err != nil {
		logrus.Errorf("error during cleanup: %v", err)
	}

	if err := cleanup.Cleanup("./generated"); err != nil {
		logrus.Errorf("error during cleanup: %v", err)
	}
}