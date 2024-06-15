package parcel_tracking_chaincode


import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "github.com/your-organization/parcel-tracking-chaincode"
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

