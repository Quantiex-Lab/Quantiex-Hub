module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	require("dotenv").config();
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const contract = truffleContract(
		require("../build/contracts/BridgeRegistry.json")
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
	contract.setProvider(web3.currentProvider);
	try {
		/*******************************************
		 *** Contract interaction
		 ******************************************/
		const address = await contract.deployed().then(function(instance) {
			return instance.address;
		});

		return console.log("BridgeRegistry deployed contract address: ", address);
	} catch (error) {
		console.error({error})
	}
};
