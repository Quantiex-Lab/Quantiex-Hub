module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const stakingPoolContract = truffleContract(
		require("../build/contracts/StakingPool.json")
	);
	const erc20TokenContract = truffleContract(
		require("../build/contracts/BridgeERC20Token.json")
	);

	const NULL_ADDRESS = "0x0000000000000000000000000000000000000000";

	/*******************************************
	 *** Constants
	 ******************************************/
		// Config values
	const NETWORK_ROPSTEN = process.argv[4] === "--network" && process.argv[5] === "ropsten";
	const NETWORK_DEVELOP = process.argv[4] === "--network" && process.argv[5] === "develop";
	const NETWORK_ETHDEV = process.argv[4] === "--network" && process.argv[5] === "ethdev";
	const NUM_ARGS = process.argv.length - 4;

	/*******************************************
	 *** Command line argument error checking
	 ***
	 *** truffle exec lacks support for dynamic command line arguments:
	 *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
	 ******************************************/
	if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
		if (NUM_ARGS !== 4) {
			return console.error(
				"Error: invalid number of parameters, please try again."
			);
		}
	} else {
		if (NUM_ARGS !== 2) {
			return console.error(
				"Error: must specify recipient address, token address, and amount."
			);
		}
	}

	/*******************************************
	 *** Stake transaction parameters
	 ******************************************/
	let sender = "";
	let amount = "";

	if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
		sender = process.argv[6];
		amount = process.argv[7];
	} else {
		sender = process.argv[4];
		amount = process.argv[5];
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
	stakingPoolContract.setProvider(web3.currentProvider);
	erc20TokenContract.setProvider(web3.currentProvider);
	
	try {
		/*******************************************
		 *** Contract interaction
		 ******************************************/
		let _amount = web3.utils.toWei(amount)

		{
			//get stake token address
			const tokenAddress = await stakingPoolContract.deployed().then(function(instance) {
				return instance.tokenAddress();
			});
			if (tokenAddress === NULL_ADDRESS) {
				console.log("Stake token address is NULL.");
				return;
			}
			console.log("Stake token address: ", tokenAddress, "\n");

			// Send approve transaction
			const stakingPoolAddress = await stakingPoolContract.deployed().then(function(instance) {
				return instance.address;
			});

			let instance = await erc20TokenContract.at(tokenAddress);
			const { logs } = await instance.approve(stakingPoolAddress, _amount, {
				chainId: 5777,
				from: sender,
				value: 0,
				gas: 300000 // 300,000 Gwei
			});

			console.log("Sent approval...");

			// Get event logs
			const approveEvent = logs.find(e => e.event === "Approval");

			// Parse event fields
			const approveLog = {
				owner: approveEvent.args.owner,
				spender: approveEvent.args.spender,
				value: Number(approveEvent.args.value)
			};

			console.log(approveLog);
		}

		// Send stake transaction
		console.log("Connecting to StakingPool contract....");
		const { logs } = await stakingPoolContract.deployed().then(function(instance) {
			console.log("Call depositInterestFund...");
			return instance.depositInterestFund(_amount, {
				chainId: 5777,
				from: sender,
				value: 0,
				gas: 300000 // 300,000 Gwei
			});
		});

		// Get event logs
		const event = logs.find(e => e.event === "LogDepositInterestFund");

		// Parse event fields
		const depositEvent = {
			sender: event.args.sender,
			amount: Number(event.args.amount),
			interestFund: Number(event.args.interestFund)
		};
		console.log(depositEvent);
		
	} catch (error) {
		console.error({ error });
	}
	return;
};
