module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	require("dotenv").config();
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");
	const BigNumber = require("bignumber.js")

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const stakingPoolContract = truffleContract(require("../build/contracts/StakingPool.json"));

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
	stakingPoolContract.setProvider(web3.currentProvider);

	try {
		/*******************************************
		 *** Contract interaction
		 ******************************************/
		let stakersCount = 0
		let stakers;
		await stakingPoolContract.deployed().then(async function (instance) {
			stakersCount = await instance.stakersCount();
			stakers = await instance.getTop(stakersCount);
		});

		if (stakersCount > 0) {
			for (let i = 0; i < stakersCount; i++) {
				await stakingPoolContract.deployed().then(async function (instance) {
					const weight = new BigNumber(await instance.weightOf(stakers[i]));
					let realWeight = weight.div(new BigNumber(10).pow(18));
					console.log("Validator:" + stakers[i] + ", Weight:", realWeight.toString(10));

					const times = await instance.getStakeTimes(stakers[i]);
					for (let idx = 0; idx < times; ++idx) {
						const record = await instance.getStakeRecord(stakers[i], idx);
						console.log("-> stake time:", new BigNumber(record[0]).toString(10), ", balance:",
							new BigNumber(record[1]).div(new BigNumber(10).pow(18)).toString(10));
					}
				});
				console.log("");
			}
		} else {
			console.log("No validators.");
		}
	} catch (error) {
		console.error({ error })
	}
};
