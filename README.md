# MasterThesis-Lightning-Channel
[PaymentChannel-BasedProcessMonitoring.pdf](https://github.com/user-attachments/files/16553184/PaymentChannel-BasedProcessMonitoring.pdf)

## Overview
This repository contains the codebase for a Command-Line Interface (CLI) tool designed to facilitate secure peer-to-peer messaging on the blockchain using the Lightning Network. This project is part of a master's thesis that explores integrating the Lightning Network into Business Process Management (BPM) systems to enhance transaction efficiency, reduce costs, and maintain security.
## Introduction
Blockchain technology offers significant potential for enhancing security and transparency in various applications. However, traditional blockchain implementations face issues such as high transaction costs, latency, and scalability. This project addresses these issues by leveraging the Lightning Network, a second-layer solution that enables off-chain transactions, thus offering faster, cheaper, and scalable interactions.
## Installation

To set up the project, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/DogukanGun/MasterThesis-Lightning-Channel.git
   cd MasterThesis-Lightning-Channel

2. **Install Go** Ensure you have Go installed on your machine. You can download it from Go's official site.
3. **Install dependencies**
   ```bash
   go mod tidy

4. **Set up Bitcoin and Lightning Network nodes** You need running instances of Bitcoin and Lightning Network nodes. Follow the setup instructions for Bitcoin Core and lnd.

## Creating Channels
To create a payment channel with a peer, use the following command:
```bash
lnmsg create --peer <PEER_ADDRESS> --amount <AMOUNT>
```
**peer:** The address of the peer with whom you want to open the channel. \n
**amount:** The amount of Bitcoin to fund the channel.
This command sets up a payment channel with the specified peer and funds it with the specified amount of Bitcoin.

## Sending Messages
To send a message through an established Lightning Network channel:
```bash
lnmsg send --peer <PEER_ADDRESS> --message "Your message here"
```
**peer** The address of the peer to whom you want to send the message.
**message** The content of the message you want to send.
This sends an encrypted message to the specified peer using the payment channel.


## Receiving Messages
To listen for incoming messages:
```bash
lnmsg receive
```
This command continuously listens for messages from peers and displays them as they are received.

## Closing Channels
To close an active payment channel:
```bash
lnmsg close --peer <PEER_ADDRESS>
```
**peer** The address of the peer with whom you want to close the channel.
This closes the specified payment channel and settles the final balances on the Bitcoin blockchain. The command ensures that all transactions are recorded and verified.


