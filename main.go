package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/awslabs/goformation/v5/cloudformation"
	"github.com/awslabs/goformation/v5/cloudformation/sns"
)

func main() {

	// delare a string flag used to decide output format: json or yaml
	format := flag.String("o", "json", "output format - json or yaml")

	// Create a new CloudFormation template
	template := cloudformation.NewTemplate()

	// Create an Amazon SNS topic, with a unique name based off the current timestamp
	template.Resources["MyTopic"] = &sns.Topic{
		TopicName: "my-topic-" + strconv.FormatInt(time.Now().Unix(), 10),
	}

	// Create a subscription, connected to our topic, that forwards notifications to an email address
	template.Resources["MyTopicSubscription"] = &sns.Subscription{
		TopicArn: cloudformation.Ref("MyTopic"),
		Protocol: "email",
		Endpoint: "some.email@example.com",
	}

	flag.Parse()
	if *format == "json" {

		// Let's see the JSON AWS CloudFormation template
		j, err := template.JSON()
		if err != nil {
			fmt.Printf("Failed to generate JSON: %s\n", err)
		} else {
			fmt.Printf("%s\n", string(j))
		}
	} else {

		// and also the YAML AWS CloudFormation template
		y, err := template.YAML()
		if err != nil {
			fmt.Printf("Failed to generate YAML: %s\n", err)
		} else {
			fmt.Printf("%s\n", string(y))
		}
	}

}
