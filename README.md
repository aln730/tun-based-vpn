# TUN-Based Encrypted VPN

A minimal peer-to-peer VPN in Go using TUN interfaces and UDP encryption.

## Features
- TUN interface for raw IP tunneling
- Encrypted UDP transport (XChaCha20-Poly1305)
- Lightweight and fast

## Usage
### Peer A
```bash
sudo ./vpn --listen :9000 --peer-ip 10.0.0.1
sudo ip addr add 10.0.0.1/24 dev tun0
sudo ip link set tun0 up
```

### Peer B
```bash
sudo ./vpn --connect A_IP:9000 --peer-ip 10.0.0.2
sudo ip addr add 10.0.0.2/24 dev tun0
sudo ip link set tun0 up
```

## Build
```bash
go build -o vpn cmd/vpn.go
```
