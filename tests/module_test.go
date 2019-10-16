package main

import (
    "testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/google/uuid"
)

func Test_module(t *testing.T) {
    expectedName := uuid.New().String()
    region := "us-west-2"

    terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

        // No Color
        NoColor: true,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"key_name": expectedName,
		},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": region,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

    var fingerprint = terraform.Output(t, terraformOptions, "fingerprint")

    t.Log("The template has been applied")

    t.Log("Checking to see if the key exists")
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })

    if err != nil {
        t.Error(err)
    }

    svc := ec2.New(sess)

    result, err := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
        KeyNames: []*string{
            aws.String(expectedName),
        },
    })

    if err != nil {
        t.Error(err)
    }

    var actual *ec2.KeyPairInfo

    for _, item := range result.KeyPairs {
        if *item.KeyName == expectedName {
            actual = item
        }
    }

    if actual == nil {
        t.Errorf("Unable to find the expected key pair: %s", expectedName)
    }
    if *actual.KeyFingerprint != fingerprint {
        t.Errorf("Fingerprints don't match: %s != %s", *actual.KeyFingerprint, fingerprint)
    }
}

func Test_module_2048_rsabits(t *testing.T) {
    expectedName := uuid.New().String()
    expectedBits := "2048"
    region := "us-west-2"

    terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

        // No Color
        NoColor: true,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"key_name": expectedName,
            "rsa_bits": expectedBits,
		},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": region,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

    var fingerprint = terraform.Output(t, terraformOptions, "fingerprint")

    t.Log("The template has been applied")

    t.Log("Checking to see if the key exists")
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })

    if err != nil {
        t.Error(err)
    }

    svc := ec2.New(sess)

    result, err := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
        KeyNames: []*string{
            aws.String(expectedName),
        },
    })

    if err != nil {
        t.Error(err)
    }

    var actual *ec2.KeyPairInfo

    for _, item := range result.KeyPairs {
        if *item.KeyName == expectedName {
            actual = item
        }
    }

    if actual == nil {
        t.Errorf("Unable to find the expected key pair: %s", expectedName)
    }
    if *actual.KeyFingerprint != fingerprint {
        t.Errorf("Fingerprints don't match: %s != %s", *actual.KeyFingerprint, fingerprint)
    }
}
