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
		const contract = truffleContract(
			require("../build/contracts/BridgeERC20Bank.json")
		);

		/*******************************************
		 *** Constants
		 ******************************************/
		const NETWORK_BSCDEV =
			process.argv[4] === "--network" && process.argv[5] === "bscdev";

		/*******************************************
		 *** Web3 provider
		 *** Set contract provider based on --network flag
		 ******************************************/
		let provider;
		if (NETWORK_BSCDEV) {
			provider = new HDWalletProvider(
				process.env.MNEMONIC,
				process.env.HDWALLET_PROVIDER
			);
		} else {
			provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
		}

		const web3 = new Web3(provider);
		contract.setProvider(web3.currentProvider);
		/*******************************************
		 *** Contract interaction
		 ******************************************/
		const address = await contract.deployed().then(function(instance) {
			return instance.address;
		});

		return console.log("BridgeERC20Bank deployed contract address: ", address);
	} catch (error) {
		console.error({error})
	}
};
