package command

import "runtime"

// GetEndpoint for docker
func GetEndpoint() string {
	if runtime.GOOS == "windows" {
		return "tcp://127.0.0.1:2375"
	}
	return "unix:///var/run/docker.sock"
}

// Images to use
var Images = []string{"spurious/sqs", "spurious/s3", "smaj/memcached", "spurious/dynamodb", "spurious/browser", "spurious/elasticache"}

// SpuriousContainers label
var SpuriousContainers = map[string][]string{
	"label": []string{
		"uk.co.spurious",
	},
}

var stoppedSpuriousContainers = map[string][]string{
	"label": []string{
		"uk.co.spurious",
	},
	"status": []string{
		"exited",
	},
}
