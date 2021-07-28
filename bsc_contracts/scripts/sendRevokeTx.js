module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const stakingPoolContract = truffleContract(
		require("../build/contracts/StakingPool.json"));

	/*******************************************
	 *** Constants
	 ******************************************/
	// Config values
	const NETWORK_DEVELOP = process.argv[4] === "--network" && process.argv[5] === "develop";
	const NETWORK_ETHDEV = process.argv[4] === "--network" && process.argv[5] === "ethdev";
	const NUM_ARGS = process.argv.length - 4;

	/*******************************************
	 *** Command line argument error checking
	 ***
	 *** truffle exec lacks support for dynamic command line arguments:
	 *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
	 ******************************************/
	if (NETWORK_DEVELOP || NETWORK_ETHDEV) {
		if (NUM_ARGS !== 3) {
			return console.error("Error: invalid number of parameters, please try again.");
		}
	} else {
		if (NUM_ARGS !== 1) {
			return console.error("Error: invalid number of parameters, please try again.");
		}
	}

	/*******************************************
	 *** Stake transaction parameters
	 ******************************************/
	let staker = "";

	if (NETWORK_DEVELOP || NETWORK_ETHDEV) {
		staker = process.argv[6];
	} else {
		staker = process.argv[4];
	}

	/*******************************************
	 *** Web3 provider
	 *** Set contract provider based on --network flag
	 ******************************************/
	let provider;

	if (NETWORK_ETHDEV) {
		provider = new HDWalletProvider(
			process.env.MNEMONIC,
			process.env.HDWALLET_PROVIDER
		);
	} else {
		provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
	}

	const web3 = new Web3(provider);
	stakingPoolContract.setProvider(web3.currentProvider);

	try {
		/*******************************************
		 *** Contract interaction
		 ******************************************/

		// Send stake transaction
		console.log("Connecting to StakingPool contract ...");
		const { logs } = await stakingPoolContract.deployed().then(function (instance) {
			console.log("Sending revoke ...");
			return instance.revoke({
				chainId: 5777,
				from: staker,
				value: 0,
				gas: 300000 // 300,000 Gwei
			});
		});

		// Get event logs
		const event = logs.find(e => e.event === "LogRevoke");

		// Parse event fields
		const revokeEvent = {
			staker: event.args.staker,
			amount: Number(event.args.amount),
			interest: Number(event.args.interest),
			balance: Number(event.args.balance),
		};
		console.log(revokeEvent);

	} catch (error) {
		console.error({ error });
	}

	return;
};
