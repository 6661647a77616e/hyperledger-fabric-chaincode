package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ParcelTrackingChaincode implements the Chaincode interface
type ParcelTrackingChaincode struct {
	contractapi.Contract
}

// Asset represents a simple asset (parcel)
type Asset struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	Owner     string `json:"owner"`
	MatricNum string `json:"matricNum"`
}

// InitLedger initializes the ledger with some example parcels
func (p *ParcelTrackingChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	parcels := []Asset{
		{ID: "parcel1", Value: "100", Owner: "Mohamad Fadzwan Ashriq Bin Zainurul Hakim", MatricNum: "2116337"},
		{ID: "parcel2", Value: "200", Owner: "Bob Ali", MatricNum: "2116338"},
		{ID: "parcel3", Value: "300", Owner: "Charlie", MatricNum: "2116339"},
	}

	for _, parcel := range parcels {
		parcelJSON, err := json.Marshal(parcel)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(parcel.ID, parcelJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset creates a new parcel with given details
func (p *ParcelTrackingChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, id, value, owner, matricNum string) error {
	exists, err := p.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID:        id,
		Value:     value,
		Owner:     owner,
		MatricNum: matricNum,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// TransferAsset transfers ownership of a parcel to a new owner
func (p *ParcelTrackingChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, id, newOwner string) error {
	asset, err := p.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset retrieves an asset from the ledger
func (p *ParcelTrackingChaincode) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	parcelJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if parcelJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(parcelJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// AssetExists checks if an asset exists in the ledger
func (p *ParcelTrackingChaincode) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	parcelJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return parcelJSON != nil, nil
}

// Main function is not included in chaincode.go to avoid redeclaration error
// This file should be compiled and deployed as a chaincode on Hyperledger Fabric network


