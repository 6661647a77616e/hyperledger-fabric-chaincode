
---

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git curl wget make jq docker.io docker-compose
```

**Enable Docker Desktop integration for WSL:**
- Go to **Settings > Resources > Integrations**
- Enable integration with your Kali Linux WSL distribution

### Setup Golang

```bash
cd ~
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Set environment variables
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPROXY=https://proxy.golang.org,direct' >> ~/.bashrc
source ~/.bashrc

# Test installation
go version
```

### Install Hyperledger Fabric

```bash
mkdir -p $HOME/go/src/github.com/hyperledger
cd $HOME/go/src/github.com/hyperledger
git clone https://github.com/hyperledger/fabric-samples.git
cd fabric-samples

curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh

./install-fabric.sh --fabric-version 2.5.12 binary
```

### Create Channel

```bash
./network.sh up createChannel
```

---

## reset  
```
rm -rf ~/go
sudo rm -rf /usr/local/go
```
