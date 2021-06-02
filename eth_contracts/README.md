# Unidirectional Peggy Project Specification

This project specification focuses on the role of 'Peggy', a Smart Contract system deployed to the Ethereum network as part of the Ethereum Binance Bridge project, and is meant to contextualize its role within the bridge. Specifications detailing structure and process of the non-Ethereum components (Relayer service, EthOracle module, Oracle module) will soon be available in the binance-ethereum-bridge repository linked below.

## Project Summary

Unidirectional Peggy is the starting point for cross chain value transfers from the Ethereum blockchain to Binance based blockchains as part of the Ethereum Binance Bridge project. The smart contract system accepts incoming transfers of Ethereum and ERC20 tokens, locking them while the transaction is validated and equitable funds issued to the intended recipient on the Binance bridge chain.

## Project Background

We are hoping to create a closed system for intra network transfers of cryptocurrency between blockchains, spearheaded by a proof-of-concept which enables secured transactions between Ethereum and Binance.

## Smart Contract Scope

### Goals of the Smart Contracts

1. Securely implement core functionality of the system such as asset locking and event emission without endangering any user funds. As such, this prototype does not permanently lock value and allows the original sender full access to their funds at any time.
2. Interface with the Relayer service, which is used by validators to listen for contract events which are signed and submitted to the Binance network as proof of transaction.
3. Successfully end-to-end test the Binance Ethereum Bridge, sending Ethereum and ERC20 tokens from Ethereum to Binance.

### Non-Goals of the Smart Contracts

1. Creating a production-grade system for cross-chain value transfers which enforces strict permissions and limits access to locked funds.
2. Implementing a validator set which enables observers to submit proof of fund locking transactions on Binance to Peggy. These features are not required for unidirectional transfers from Ethereum to Binance and will be re-integrated during phase two of the project, which aims to send funds from Binance back to Ethereum.
3. Fully gas optimize and streamline operational functionality; ease and clarity of testing has been favored over some gas management and architectural best practices.

## Ethereum Binance Bridge Architecture

Unidirectional Peggy focuses on core features for unidirectional transfers. This prototype includes functionality to safely lock and unlock Ethereum and ERC20 tokens, emitting associated events which are witnessed by validators using the Relayer service. The Relayer is a service which interfaces with both blockchains, allowing validators to attest on the Binance blockchain that specific events on the Ethereum blockchain have occurred. The Relayer listens for `LogLock` events, parses information associated with the Ethereum transaction, uses it to build unsigned Binance transactions, and enables validators to sign and send the transactions to the Oracle module on Binance. Through the Relayer service, validators witness the events and submit proof in the form of signed hashes to the Binance based modules, which are responsible for aggregating and tallying the Validators’ signatures and their respective signing power. The system is managed by the contract's deployer, designated internally as the provider, a trusted third-party which can unlock funds and return them their original sender. If the contract’s balances under threat, the provider can pause the system, temporarily preventing users from depositing additional funds.

The Peggy Smart Contract is deployed on the Ropsten testnet at address: 0x05d9758cb6b9d9761ecb8b2b48be7873efae15c0

### Architecture Diagram

![peggyarchitecturediagram](https://user-images.githubusercontent.com/15370712/58388886-632c7700-7fd9-11e9-962e-4e5e9d92c275.png)

### System Process:

1. Users lock Ethereum or ERC20 tokens on the Peggy contract, resulting in the emission of an event containing the created item's original sender's Ethereum address, the intended recipient's Binance address, the type of token, the amount locked, and the item's unique nonce.
2. Validators on the Binance chain witness these lock events via a Relayer service and sign a hash containing the unique item's information, which is sent as a Binance transaction to Oracle module.
3. Once the Oracle module has verified that the validators' aggregated signing power is greater than the specified threshold, it mints the appropriate amount of tokens and forwards them to the intended recipient.

The Relayer service and Oracle module are under development here: https://github.com/binance/binance-ethereum-bridge.

## Installation

Install Truffle: `$ npm install -g truffle`  
Install dependencies: `$ npm install`

Note: This project currently uses solc@0.5.0, make sure that this version of the Solidity compiler is being used to compile the contracts and does not conflict with other versions that may be installed on your machine.

## Testing

Run commands from the appropriate directory: `$ cd eth_contracts`  
Start the truffle environment: `$ truffle develop`  
In another tab, run tests: `$ truffle test`  
Run individual tests: `$ truffle test test/<test_name.js>`

Expected output of the test suite:
![peggytestsuite](https://user-images.githubusercontent.com/15370712/58388940-34fb6700-7fda-11e9-9aef-6ae7b2442a55.png)

## Security, Privacy, Risks

Disclaimer: These contracts are for testing purposes only and are NOT intended for production. In order to prevent any loss of user funds, locked Ethereum and ERC20 tokens can be withdrawn directly by the original sender at any time. However, these contracts have not undergone external audits and should not be trusted with mainnet funds. Any use of Peggy is at the user’s own risk.

## Other Considerations

We decided to temporarily remove the validator set from this version of the Smart Contracts, our reasoning being that system transparency and clarity should take precedence over the inclusion of future project features. The validator set is not required on Ethereum for unidirectional transfers and will be reimplemented once it is needed in the bidrectional version to validate transactions that have occured on Binance.

## Ongoing work

The Ethereum Oracle module and Oracle modules are completed, with the Relayer service currently being actively integrated in order to interface between the smart contracts and Oracles. Once Ethereum -> Binance transfers have been successfully prototyped, functionality for bidirectional transfers (such as validator sets, signature validation, and secured token unlocking procedures) will be integrated into the contracts. Previous work in these areas is a valuable resource that will be leveraged once the complete system is ready for bidirectional transfers.

Thanks to @adrianbrink, @mossid, and @sunnya97 for contributions to the original Peggy repository.
