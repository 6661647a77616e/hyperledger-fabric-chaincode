package main


import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
    chaincode, err := contractapi.NewChaincode(&parcel_tracking_chaincode.ParcelTrackingChaincode{})
    if err != nil {
        panic(err)
    }

    if err := chaincode.Start(); err != nil {
        panic(err)
    }
}

