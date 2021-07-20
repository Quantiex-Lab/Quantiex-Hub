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
		const NETWORK_DEVELOP =
			process.argv[4] === "--network" && process.argv[5] === "develop";
		const NETWORK_ETHDEV =
			process.argv[4] === "--network" && process.argv[5] === "ethdev";

		let tokenType
		let account
		let token
		if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
			tokenType = process.argv[6];
			account = process.argv[7];
			token = (process.argv[8] || 'eth').toString();
		} else {
			tokenType = process.argv[4];
			account = process.argv[5];
			token = (process.argv[6] || 'eth').toString();
		}

		if (tokenType !== "erc20" && tokenType !== "erc721")
		{
			console.log("Please input token type: erc20, erc721");
			return;
		}

		if (!account) {
			console.log("Please provide an Ethereum address to check their balance");
			return;
		}

		/*******************************************
		 *** Web3 provider
		 *** Set contract provider based on --network flag
		 ******************************************/
		let provider;
		let operator;
		if (NETWORK_ROPSTEN) {
			provider = new HDWalletProvider(
				process.env.MNEMONIC,
				"https://ropsten.infura.io/v3/".concat(process.env.INFURA_PROJECT_ID)
			);
			operator = process.env.OPERATOR;
		} else if (NETWORK_ETHDEV) {
			provider = new HDWalletProvider(
				process.env.MNEMONIC,
				process.env.HDWALLET_PROVIDER
			);
			operator = process.env.OPERATOR;
		} else {
			provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
			operator = process.env.LOCAL_OPERATOR;
		}

		const web3 = new Web3(provider);
		erc20TokenContract.setProvider(web3.currentProvider);
		erc721TokenContract.setProvider(web3.currentProvider);

		/*******************************************
		 *** Contract interaction
		 ******************************************/
		if (tokenType === "erc20")
		{
			let balanceWei
			let balanceEth
			if (token === "eth") {
				balanceWei = await web3.eth.getBalance(account)
				balanceEth = web3.utils.fromWei(balanceWei)
				console.log(`BNB balance for ${account} is ${balanceEth} BNB (${balanceWei} Wei)`)
				return
			}

			const tokenInstance = await erc20TokenContract.at(token)
			const name = await tokenInstance.name()
			const symbol = await tokenInstance.symbol()
			const decimals = await tokenInstance.decimals()
			balanceWei = new BigNumber(await tokenInstance.balanceOf(account))
			balanceEth = balanceWei.div(new BigNumber(10).pow(decimals.toNumber()))
			console.log(`Balance of ${name} for ${account} is ${balanceEth.toString(10)} ${symbol} (${balanceWei} ${symbol} with ${decimals} decimals)`)

		} else if (tokenType === "erc721") {

			const tokenInstance = await erc721TokenContract.at(token)
			const name = await tokenInstance.name()
			const symbol = await tokenInstance.symbol()
			const balanceToken = new BigNumber(await tokenInstance.balanceOf(account))
			console.log(`Balance of ${name} for ${account} is ${balanceToken.toString(10)} ${symbol}`)
		}

	} catch (error) {
		console.error({error})
	}

	return;
};
