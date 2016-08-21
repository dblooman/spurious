package command

// Endpoint for Docker
var Endpoint = "unix:///var/run/docker.sock"

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
