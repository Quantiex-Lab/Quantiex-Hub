module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	require("dotenv").config();
	const Web3 = require("web3");
	const BigNumber = require("bignumber.js")
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	try {
		// Contract abstraction
		const truffleContract = require("truffle-contract");
		const erc721TokenContract = truffleContract(
			require("../build/contracts/BridgeERC721Token.json")
		);

		/*******************************************
		 *** Constants
		 ******************************************/
		const NETWORK_ROPSTEN =
			process.argv[4] === "--network" && process.argv[5] === "ropsten";
		const NETWORK_DEVELOP =
			process.argv[4] === "--network" && process.argv[5] === "develop";
		const NETWORK_ETHDEV =
			process.argv[4] === "--network" && process.argv[5] === "ethdev";

		let tokenType
		let account
		let token
		if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
			account = process.argv[6];
			token = process.argv[7];
		} else {
			account = process.argv[4];
			token = process.argv[5];
		}

		if (!account) {
			console.log("Please provide an Ethereum address to check their balance")
			return
		}

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
		erc721TokenContract.setProvider(web3.currentProvider);

		const tokenInstance = await erc721TokenContract.at(token);
		const tokens = await tokenInstance.listTokens(account);

		if (tokens.length === 0)
		{
			console.log("No tokens.");
		} else {
			console.log("Token Ids:");
			for (var i = 0; i < tokens.length; i++)
			{
				console.log(tokens[i].toString());
			}
			console.log("");
		}

	} catch (error) {
		console.error({error})
	}

	return;
};
