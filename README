# go-sweet

Go client library for interacting with Sweet’s API

---

## Usage

```go
package main

import (
	"log/slog"
	"os"

	"github.com/sweet-security/go-sweet/sweet"
)

func main() {
	sweetClient := sweet.New(
		"00000000-0000-0000-0000-000000000000",
		"00000000-0000-0000-0000-000000000000",
		sweet.WithEnv("stg"),
		sweet.WithSubenv("main"),
	)

	awsAccounts, err := sweetClient.GetAwsAccounts()
	if err != nil {
		slog.Error("error getting aws accounts", "err", err)
		os.Exit(1)
	}
	for _, awsAccount := range *awsAccounts {
		slog.Info("got aws account", "aws account", awsAccount)
	}

	awsAccount, err := sweetClient.AddAwsAccount(&sweet.AwsAccount{
		AccountId: "123456789012",
		RoleArn:   "arn:aws:iam::123456789012:role/example",
	})
	if err != nil {
		slog.Error("error adding aws account", "err", err)
		os.Exit(1)
	}
	slog.Info("added aws account", "aws account", awsAccount)

	awsAccount, err = sweetClient.GetAwsAccount("123456789012")
	if err != nil {
		slog.Error("error getting aws account", "err", err)
		os.Exit(1)
	}
	slog.Info("got aws account", "aws account", awsAccount)

	awsAccount, err = sweetClient.UpdateAwsAccount(&sweet.AwsAccount{
		AccountId:  "123456789012",
		RoleArn:    "arn:aws:iam::123456789012:role/example",
		ExternalId: "testing",
	})
	if err != nil {
		slog.Error("error updating aws account", "err", err)
		os.Exit(1)
	}
	slog.Info("updated aws account", "aws account", awsAccount)

	err = sweetClient.DeleteAwsAccount(awsAccount.AccountId)
	if err != nil {
		slog.Error("error deleting aws account", "err", err)
		os.Exit(1)
	}
	slog.Info("deleted aws account", "aws account", awsAccount.AccountId)

	awsOrganizations, err := sweetClient.GetAwsOrganizations()
	if err != nil {
		slog.Error("error getting aws organizations", "err", err)
		os.Exit(1)
	}
	slog.Info("got aws organization", "aws organization", awsOrganizations)

	organization, err := sweetClient.AddAwsOrganization(&sweet.AwsOrganization{
		AccountId:            "123456789012",
		RoleArn:              "arn:aws:iam::123456789012:role/example",
		RoleNameParameterArn: "arn:aws:ssm:us-east-2:123456789012:parameter/example",
	})
	if err != nil {
		slog.Error("error adding aws organization", "err", err)
		os.Exit(1)
	}
	slog.Info("add aws organization", "aws organization", organization)

	organization, err = sweetClient.UpdateAwsOrganization(&sweet.AwsOrganization{
		AccountId:            "123456789012",
		RoleArn:              "arn:aws:iam::123456789012:role/example",
		RoleNameParameterArn: "arn:aws:ssm:us-east-2:123456789012:parameter/example",
		ExternalId:           "testing",
	})
	if err != nil {
		slog.Error("error updating aws organization", "err", err)
		os.Exit(1)
	}
	slog.Info("update aws organization", "aws organization", organization)

	err = sweetClient.DeleteAwsOrganization(organization.AccountId)
	if err != nil {
		slog.Error("error deleting aws organization", "err", err)
		os.Exit(1)
	}
	slog.Info("delete aws organization", "aws organization", organization.AccountId)
}

```
Obtain tenant api key & secret from [Sweet platform](https://app.sweet.security) inside settings -> API Tokens

---

## Install

```bash
go get github.com/sweet-security/go-sweet
