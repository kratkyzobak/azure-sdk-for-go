//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_ListByReplicationProtectionContainers.json
func ExampleReplicationMigrationItemsClient_ListByReplicationProtectionContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	pager := client.ListByReplicationProtectionContainers("<fabric-name>",
		"<protection-container-name>",
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientListByReplicationProtectionContainersOptions{SkipToken: nil,
			TakeToken: nil,
			Filter:    nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Get.json
func ExampleReplicationMigrationItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientGetResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Create.json
func ExampleReplicationMigrationItemsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.EnableMigrationInput{
			Properties: &armrecoveryservicessiterecovery.EnableMigrationInputProperties{
				PolicyID: to.StringPtr("<policy-id>"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtEnableMigrationInput{
					InstanceType:            to.StringPtr("<instance-type>"),
					DataMoverRunAsAccountID: to.StringPtr("<data-mover-run-as-account-id>"),
					DisksToInclude: []*armrecoveryservicessiterecovery.VMwareCbtDiskInput{
						{
							DiskID:                         to.StringPtr("<disk-id>"),
							IsOSDisk:                       to.StringPtr("<is-osdisk>"),
							LogStorageAccountID:            to.StringPtr("<log-storage-account-id>"),
							LogStorageAccountSasSecretName: to.StringPtr("<log-storage-account-sas-secret-name>"),
						}},
					SnapshotRunAsAccountID: to.StringPtr("<snapshot-run-as-account-id>"),
					TargetNetworkID:        to.StringPtr("<target-network-id>"),
					TargetResourceGroupID:  to.StringPtr("<target-resource-group-id>"),
					VmwareMachineID:        to.StringPtr("<vmware-machine-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientCreateResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Delete.json
func ExampleReplicationMigrationItemsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginDeleteOptions{DeleteOption: nil})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Update.json
func ExampleReplicationMigrationItemsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.UpdateMigrationItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateMigrationItemInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtUpdateMigrationItemInput{
					InstanceType: to.StringPtr("<instance-type>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientUpdateResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Migrate.json
func ExampleReplicationMigrationItemsClient_BeginMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginMigrate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.MigrateInput{
			Properties: &armrecoveryservicessiterecovery.MigrateInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtMigrateInput{
					InstanceType:    to.StringPtr("<instance-type>"),
					PerformShutdown: to.StringPtr("<perform-shutdown>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientMigrateResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Resync.json
func ExampleReplicationMigrationItemsClient_BeginResync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginResync(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.ResyncInput{
			Properties: &armrecoveryservicessiterecovery.ResyncInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtResyncInput{
					InstanceType: to.StringPtr("<instance-type>"),
					SkipCbtReset: to.StringPtr("<skip-cbt-reset>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientResyncResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_TestMigrate.json
func ExampleReplicationMigrationItemsClient_BeginTestMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginTestMigrate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.TestMigrateInput{
			Properties: &armrecoveryservicessiterecovery.TestMigrateInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtTestMigrateInput{
					InstanceType:    to.StringPtr("<instance-type>"),
					NetworkID:       to.StringPtr("<network-id>"),
					RecoveryPointID: to.StringPtr("<recovery-point-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientTestMigrateResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_TestMigrateCleanup.json
func ExampleReplicationMigrationItemsClient_BeginTestMigrateCleanup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginTestMigrateCleanup(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.TestMigrateCleanupInput{
			Properties: &armrecoveryservicessiterecovery.TestMigrateCleanupInputProperties{
				Comments: to.StringPtr("<comments>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientTestMigrateCleanupResult)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_List.json
func ExampleReplicationMigrationItemsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	pager := client.List(&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientListOptions{SkipToken: nil,
		TakeToken: nil,
		Filter:    nil,
	})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
