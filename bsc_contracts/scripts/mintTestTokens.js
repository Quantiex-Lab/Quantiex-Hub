module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
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
		// Config values
	const NETWORK_BSCDEV =
		process.argv[4] === "--network" && process.argv[5] === "bscdev";

	/*******************************************
	 *** Web3 provider
	 ******************************************/
	let provider;
	let operator;
	if (NETWORK_BSCDEV) {
		provider = new HDWalletProvider(
			process.env.MNEMONIC,
			process.env.HDWALLET_PROVIDER
		);
		operator = process.env.OPERATOR;
	} else {
		provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
		operator = process.env.LOCAL_OPERATOR;
	}
	// const provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
	const web3 = new Web3(provider);
	erc20TokenContract.setProvider(web3.currentProvider);
	erc721TokenContract.setProvider(web3.currentProvider);

	// Mint ERC20 token
	try {
		// Mint transaction parameters
		const TOKEN_AMOUNT = (1).toString().padEnd(25, "0")
		console.log({TOKEN_AMOUNT})
		// Send mint transaction
		const { logs } = await erc20TokenContract.deployed().then(
			function(instance) {
				return instance.mint(operator, TOKEN_AMOUNT, {
					from: operator,
					value: 0,
					gas: 300000 // 300,000 Gwei
				});
			});

		// Get event logs
		const event = logs.find(e => e.event === "Transfer");

		// Parse event fields
		const transferEvent = {
			from: event.args.from,
			to: event.args.to,
			value: Number(event.args.value)
		};
		console.log(transferEvent);

	} catch (error) {
		console.error({error})
	}

	// Mint ERC721 token
	try {
		for (i = 0; i < 10; i++)
		{
			// Send mint transaction
			const { logs } = await erc721TokenContract.deployed().then(
				function(instance) {
					return instance.mintTo(operator, i, "token" + i.toString(), {
						from: operator,
						value: 0,
						gas: 300000 // 300,000 Gwei
					});
				});

			// Get event logs
			const event = logs.find(e => e.event === "Transfer");

			// Parse event fields
			const transferEvent = {
				from: event.args.from,
				to: event.args.to,
				tokenId: Number(event.args.tokenId)
			};
			console.log(transferEvent);
		}
	} catch (error) {
		console.error({error})
	}

	return;
};
