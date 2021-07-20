module.exports = async () => {
	/*******************************************
	 *** Set up
	 ******************************************/
	const Web3 = require("web3");
	const HDWalletProvider = require("@truffle/hdwallet-provider");

	// Contract abstraction
	const truffleContract = require("truffle-contract");
	const erc20BridgeContract = truffleContract(
		require("../build/contracts/QuantiexERC20Bridge.json")
	);
	const erc721BridgeContract = truffleContract(
		require("../build/contracts/QuantiexERC721Bridge.json")
	);
	const oracleContract = truffleContract(
		require("../build/contracts/Oracle.json")
	);
	const erc20BankContract = truffleContract(
		require("../build/contracts/BridgeERC20Bank.json")
	);
	const erc721BankContract = truffleContract(
		require("../build/contracts/BridgeERC721Bank.json")
	);

	/*******************************************
	 *** Constants
	 ******************************************/
		// Config values
	const NETWORK_ROPSTEN =
		process.argv[4] === "--network" && process.argv[5] === "ropsten";
	const NETWORK_BSCDEV =
		process.argv[4] === "--network" && process.argv[5] === "bscdev";

	/*******************************************
	 *** Web3 provider
	 *** Set contract provider based on --network flag
	 ******************************************/
	let provider;
	let operator;
	if (NETWORK_ROPSTEN || NETWORK_BSCDEV) {
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

	erc20BridgeContract.setProvider(web3.currentProvider);
	erc721BridgeContract.setProvider(web3.currentProvider);
	oracleContract.setProvider(web3.currentProvider);
	erc20BankContract.setProvider(web3.currentProvider);
	erc721BankContract.setProvider(web3.currentProvider);

	try {
		/*******************************************
		 *** Contract interaction
		 ******************************************/
			// Get deployed Oracle's address
		const oracleContractAddress = await oracleContract
				.deployed()
				.then(function(instance) {
					return instance.address;
				});

		// Set Oracle
		{
			const { logs: setOracleLogs } = await erc20BridgeContract
				.deployed()
				.then(function(instance) {
					return instance.setOracle(oracleContractAddress, {
						from: operator,
						value: 0,
						gas: 300000 // 300,000 Gwei
					});
				});
			const setOracleEvent = setOracleLogs.find(e => e.event === "LogOracleSet");
			console.log("QuantiexERC20Bridge's Oracle set:", setOracleEvent.args._oracle);
		}
		{
			const { logs: setOracleLogs } = await erc721BridgeContract
				.deployed()
				.then(function(instance) {
					return instance.setOracle(oracleContractAddress, {
						from: operator,
						value: 0,
						gas: 300000 // 300,000 Gwei
					});
				});
			const setOracleEvent = setOracleLogs.find(e => e.event === "LogOracleSet");
			console.log("QuantiexERC721Bridge's Oracle set:", setOracleEvent.args._oracle);
		}

		// Set BridgeERC20Bank
		{
			// Get deployed BridgeERC20Bank's address
			const erc20BankContractAddress = await erc20BankContract
				.deployed()
				.then(function(instance) {
					return instance.address;
				});
			const { logs: setBridgeBankLogs } = await erc20BridgeContract.deployed().then(
				function(instance) {
					return instance.setBridgeBank(erc20BankContractAddress, {
						from: operator,
						value: 0,
						gas: 300000 // 300,000 Gwei
					});
				});
			// Get event logs
			const setBridgeBankEvent = setBridgeBankLogs.find(
				e => e.event === "LogBridgeBankSet"
			);
			console.log(
				"QuantiexERC20Bridge's BridgeERC20Bank set:",
				setBridgeBankEvent.args._bridgeBank
			);
		}

		// Set BridgeERC721Bank
		{
			// Get deployed BridgeERC721Bank's address
			const erc721BankContractAddress = await erc721BankContract
				.deployed()
				.then(function(instance) {
					return instance.address;
				});
			const { logs: setBridgeBankLogs } = await erc721BridgeContract.deployed().then(
				function(instance) {
					return instance.setBridgeBank(erc721BankContractAddress, {
						from: operator,
						value: 0,
						gas: 300000 // 300,000 Gwei
					});
				});
			// Get event logs
			const setBridgeBankEvent = setBridgeBankLogs.find(
				e => e.event === "LogBridgeBankSet"
			);
			console.log(
				"QuantiexERC721Bridge's BridgeERC721Bank set:",
				setBridgeBankEvent.args._bridgeBank
			);
		}

		return;
	} catch (error) {
		console.error({error})
	}
};
