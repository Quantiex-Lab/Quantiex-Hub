module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	require("dotenv").config();
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	try {
		// Contract abstraction
		const truffleContract = require("truffle-contract");
		const erc20BankContract = truffleContract(
			require("../build/contracts/BridgeERC20Bank.json")
		);
		const erc721BankContract = truffleContract(
			require("../build/contracts/BridgeERC721Bank.json")
		);

		/*******************************************
		 *** Constants
		 ******************************************/
		const NETWORK_ROPSTEN =
			process.argv[4] === "--network" && process.argv[5] === "ropsten";

		/*******************************************
		 *** Web3 provider
		 *** Set contract provider based on --network flag
		 ******************************************/
		let provider;
		if (NETWORK_ROPSTEN) {
			provider = new HDWalletProvider(
				process.env.MNEMONIC,
				"https://ropsten.infura.io/v3/".concat(process.env.INFURA_PROJECT_ID)
			);
		} else if (NETWORK_ETHDEV) {
			provider = new HDWalletProvider(
				process.env.MNEMONIC,
				process.env.HDWALLET_PROVIDER
			);
		} else {
			provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
		}

		const web3 = new Web3(provider);
		erc20BankContract.setProvider(web3.currentProvider);
		erc721BankContract.setProvider(web3.currentProvider);

		/*******************************************
		 *** Contract interaction
		 ******************************************/
		const erc20BankAddress = await erc20BankContract.deployed().then(function(instance) {
			return instance.address;
		});
		console.log("BridgeERC20Bank deployed contract address: ", erc20BankAddress);

		const erc721BankAddress = await erc721BankContract.deployed().then(function(instance) {
			return instance.address;
		});
		console.log("BridgeERC721Bank deployed contract address: ", erc721BankAddress);

	} catch (error) {
		console.error({error})
	}

	return;
};
