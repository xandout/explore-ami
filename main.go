package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func getSessionForRegion(region string) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(region)},
	}))
}

func main() {

	region := *flag.String("region", "us-east-1", "Which region to query")
	ami := *flag.String("ami", "ami-a4c7edb2", "Which AMI to display")
	full := *flag.Bool("full", true, "Display full output from AWS")

	flag.Parse()

	svc := ec2.New(getSessionForRegion(region))
	input := &ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("image-id"),
				Values: []*string{aws.String(ami)},
			},
		},
	}

	result, err := svc.DescribeImages(input)

	if err != nil {
		log.Fatal(err)
	}

	if full == true {
		fmt.Println(result)
	} else {
		data := [][]string{
			[]string{*result.Images[0].Description, *result.Images[0].VirtualizationType, *result.Images[0].RootDeviceName, *result.Images[0].RootDeviceType},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Description", "Virtualization Type", "Root Dev Name", "Root Dev Type"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output
	}

}
