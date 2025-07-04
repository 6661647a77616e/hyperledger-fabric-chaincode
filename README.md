
---
# fabric newbie
```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git curl wget make jq docker.io docker-compose
```

**Enable Docker Desktop integration for WSL:**
- Go to **Settings > Resources > Integrations**
- Enable integration with your Kali Linux WSL distribution

* setup Golang

```bash
cd ~
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Set environment variables
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPROXY=https://proxy.golang.org,direct' >> ~/.bashrc
source ~/.bashrc

# test installation
go version
```

* install hyperledger fabric

```bash
mkdir -p $HOME/go/src/github.com/hyperledger
cd $HOME/go/src/github.com/hyperledger
git clone https://github.com/hyperledger/fabric-samples.git
cd fabric-samples

curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh

./install-fabric.sh --fabric-version 2.5.12 binary
```

* ensure TLS files and Proper ENV
```
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

* create Channel

```bash
cd test-network
./network.sh up createChannel
```

* test peer tls connection
```
peer channel list
```
* run cli command with tls
```
peer channel getinfo -c mychannel
```

* deploy smart contracts manually
```
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go
```

---

## reset  
```
rm -rf ~/go
sudo rm -rf /usr/local/go
docker system prune -af
```

----

# Setup Projects
* create folder in this sturcture use the code above
```
fabric-samples/
└── test-network/
    └── chaincode/
        └── parcel-tracking/
            ├── chaincode.go
            ├── main.go
            └── go.mod
```

* create `go.mod`
```
cd fabric-samples/test-network/chaincode/parcel-tracking
go mod init parcel-tracking
go mod tidy
```

* package the chaincode
```

cd ~/go/src/github.com/hyperledger/fabric-samples/bin

/peer lifecycle chaincode package parcel.tar.gz \
  --path ../test-network/chaincode/parcel-tracking \
  --lang golang \
  --label parcel_1
```

* install the chain code on both peers
- you must installl on each peer , peer0.org1 and peero0.org22
```
/peer lifecycle chaincode install parcel.tar.gz
```

```
/peer lifecycle chaincode install parcel.tar.gz
2025-07-05 02:11:19.366 +08 0001 INFO [cli.lifecycle.chaincode] submitInstallProposal -> Installed remotely: response:<status:200 payload:"\nIparcel_1:6d5ac45f8bbb7debe24d6c5f837f8f50e5b5ed54cd6ff92c48d9b5b52db8e6b4\022\010parcel_1" >
2025-07-05 02:11:19.366 +08 0002 INFO [cli.lifecycle.chaincode] submitInstallProposal -> Chaincode code package identifier: parcel_1:6d5ac45f8bbb7debe24d6c5f837f8f50e5b5ed54cd6ff92c48d9b5b52db8e6b4
```
- after that , confirm installation with 
```
./peer lifecycle chaincode queryinstalled
```
```
┌──(fadzwan㉿GreenQIQI)-[~/go/src/github.com/hyperledger/fabric-samples/bin]
└─$ ./peer lifecycle chaincode queryinstalled
Installed chaincodes on peer:
Package ID: parcel_1:6d5ac45f8bbb7debe24d6c5f837f8f50e5b5ed54cd6ff92c48d9b5b52db8e6b4, Label: parcel_1
```

* approve chaincode definition for org
template
```
./peer lifecycle chaincode approveformyorg \
  --orderer localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls \
  --cafile "$PWD/../test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
  --channelID mychannel \
  --name parcel \
  --version 1.0 \
  --package-id parcel_1:<your-package-id> \
  --sequence 1
```
example output
```
./peer lifecycle chaincode approveformyorg \
  --orderer localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls \
  --cafile "$PWD/../test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
  --channelID mychannel \
  --name parcel \
  --version 1.0 \
  --package-id parcel_1:6d5ac45f8bbb7debe24d6c5f837f8f50e5b5ed54cd6ff92c48d9b5b52db8e6b4 \
  --sequence 1
2025-07-05 02:15:34.816 +08 0001 INFO [chaincodeCmd] ClientWait -> txid [3e2ca2c5ea45ea20fb6f50c1012bf2797b2949031ca4b35f48d6cae141aa932a] committed with status (VALID) at localhost:7051
```
* check commmit readiness
```
./peer lifecycle chaincode checkcommitreadiness \
  --channelID mychannel \
  --name parcel \
  --version 1.0 \
  --sequence 1 \
  --output json
```

```
┌──(fadzwan㉿GreenQIQI)-[~/go/src/github.com/hyperledger/fabric-samples/bin]
└─$ ./peer lifecycle chaincode checkcommitreadiness \
  --channelID mychannel \
  --name parcel \
  --version 1.0 \
  --sequence 1 \
  --output json
{
        "approvals": {
                "Org1MSP": true,
                "Org2MSP": false
        }
}
```

then do for peer2
```
┌──(fadzwan㉿GreenQIQI)-[~/go/src/github.com/hyperledger/fabric-samples/bin]
└─$ export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$PWD/../test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=$PWD/../test-network/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
┌──(fadzwan㉿GreenQIQI)-[~/go/src/github.com/hyperledger/fabric-samples/bin]
└─$ ./peer lifecycle chaincode approveformyorg \
  --channelID mychannel \
  --name parcel \
  --version 1.0 \
  --package-id parcel_1:1f2e9b01234abcde5678... \
  --sequence 1 \
  --tls \
  --cafile $PWD/../test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  --orderer localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com
2025-07-05 02:19:52.464 +08 0001 INFO [chaincodeCmd] ClientWait -> txid [6d8aecb0184bf82c5dda5e69a5082c47aa02335afed2b6a07151de41f4a32a22] committed with status (VALID) at localhost:9051

┌──(fadzwan㉿GreenQIQI)-[~/go/src/github.com/hyperledger/fabric-samples/bin]
└─$ ./peer lifecycle chaincode checkcommitreadiness \
  --channelID mychannel \
  --name parcel \
  --version 1.0 \
  --sequence 1 \
  --output json
{
        "approvals": {
                "Org1MSP": true,
                "Org2MSP": true
        }
}
```

uh later cont