//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestResultsList.json
func ExampleTestResultsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestResultsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		armtestbase.OsUpdateTypeSecurityUpdate,
		&armtestbase.TestResultsListOptions{Filter: to.StringPtr("<filter>")})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("TestResultResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestResultGet.json
func ExampleTestResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestResultsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		"<test-result-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TestResultResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestResultGetDownloadURL.json
func ExampleTestResultsClient_GetDownloadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestResultsClient("<subscription-id>", cred, nil)
	_, err = client.GetDownloadURL(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		"<test-result-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestResultGetVideoDownloadURL.json
func ExampleTestResultsClient_GetVideoDownloadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestResultsClient("<subscription-id>", cred, nil)
	_, err = client.GetVideoDownloadURL(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		"<test-result-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
