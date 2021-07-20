module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const bankContract = truffleContract(
		require("../build/contracts/BridgeERC721Bank.json")
	);
	const tokenContract = truffleContract(
		require("../build/contracts/BridgeERC721Token.json")
	);

	/*******************************************
	 *** Constants
	 ******************************************/
	// Config values
	const NETWORK_ROPSTEN =
		process.argv[4] === "--network" && process.argv[5] === "ropsten";
	const NETWORK_DEVELOP =
		process.argv[4] === "--network" && process.argv[5] === "develop";
	const NETWORK_ETHDEV =
		process.argv[4] === "--network" && process.argv[5] === "ethdev";
	const NUM_ARGS = process.argv.length - 4;

	/*******************************************
	 *** Command line argument error checking
	 ***
	 *** truffle exec lacks support for dynamic command line arguments:
	 *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
	 ******************************************/
	if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
		if (NUM_ARGS !== 7) {
			return console.error(
				"Error: invalid number of parameters, please try again."
			);
		}
	} else {
		if (NUM_ARGS !== 5) {
			return console.error(
				"Error: invalid number of parameters, please try again."
			);
		}
	}

	/*******************************************
	 *** Lock transaction parameters
	 ******************************************/
	let toChainName = "";
	let ethSender = "";
	let recipient = "";
	let coinDenom = "";
	let tokenId = "";

	if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
		toChainName = process.argv[6];
		ethSender = process.argv[7];
		recipient = process.argv[8];
		coinDenom = process.argv[9];
		tokenId = process.argv[10];
	} else {
		toChainName = process.argv[4];
		ethSender = process.argv[5];
		recipient = process.argv[6];
		coinDenom = process.argv[7];
		tokenId = process.argv[8];
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
	bankContract.setProvider(web3.currentProvider);
	tokenContract.setProvider(web3.currentProvider);

	try {
		// Send approve transaction
		const bankContractAddress = await bankContract.deployed().then(
			function(instance) {
				return instance.address;
			});

		let instance = await tokenContract.at(coinDenom)
		const { logs } = await instance.approve(bankContractAddress, tokenId, {
			from: ethSender,
			value: 0,
			gas: 300000 // 300,000 Gwei
		});
		console.log("Sent approval...");

		// Get event logs
		const eventA = logs.find(e => e.event === "Approval");

		// Parse event fields
		const approvalEvent = {
			owner: eventA.args.owner,
			approved: eventA.args.approved,
			tokenId: Number(eventA.args.tokenId)
		};
		console.log(approvalEvent);

		// Send lock transaction
		console.log("Connecting to bankContract....");
		const { logs: logs2 } = await bankContract.deployed().then(
			function (instance) {
				console.log("Connected to bankContract, sending lock...");
				return instance.lock(toChainName, recipient, coinDenom, tokenId, {
					from: ethSender,
					value: 0,
					gas: 300000 // 300,000 Gwei
				});
			});

		console.log("Sent lock...");

		// Get event logs
		const event = logs2.find(e => e.event === "LogLock");

		// Parse event fields
		const lockEvent = {
			chainName: event.args._chainName,
			from: event.args._from,
			to: event.args._to,
			token: event.args._token,
			symbol: event.args._symbol,
			tokenId: Number(event.args._tokenId),
			baseURI: event.args._baseURI,
			tokenURI: event.args._tokenURI,
			nonce: Number(event.args._nonce)
		};
		console.log(lockEvent);

	} catch (error) {
		console.error({ error });
	}
	return;
};
