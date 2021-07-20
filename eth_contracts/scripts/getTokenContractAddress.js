module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	require("dotenv").config();
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const erc20TokenContract = truffleContract(
		require("../build/contracts/BridgeERC20Token.json")
	);
	const erc721TokenContract = truffleContract(
		require("../build/contracts/BridgeERC721Token.json")
	);

	/*******************************************
	 *** Constants
	 ******************************************/
	const NETWORK_ROPSTEN =
		process.argv[4] === "--network" && process.argv[5] === "ropsten";
	const NETWORK_ETHDEV =
		process.argv[4] === "--network" && process.argv[5] === "ethdev";

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
	erc20TokenContract.setProvider(web3.currentProvider);
	erc721TokenContract.setProvider(web3.currentProvider);

	/*******************************************
	 *** Contract interaction
	 ******************************************/
	const erc20TokenAddress = await erc20TokenContract.deployed().then(function(instance) {
		return instance.address;
	});
	console.log("ERC20Token contract address: ", erc20TokenAddress);

	const erc721TokenAddress = await erc721TokenContract.deployed().then(function(instance) {
		return instance.address;
	});
	console.log("ERC721Token contract address: ", erc721TokenAddress);

	return;
};
