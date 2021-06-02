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
  const tokenContract = truffleContract(
      require("../build/contracts/BridgeToken.json")
  );

  const NULL_ADDRESS = "0x0000000000000000000000000000000000000000";

  /*******************************************
   *** Constants
   ******************************************/
      // Config values
  const NETWORK_ROPSTEN =
      process.argv[4] === "--network" && process.argv[5] === "ropsten";
  const NETWORK_DEVELOP =
      process.argv[4] === "--network" && process.argv[5] === "develop";
  const NETWORK_BSCDEV =
      process.argv[4] === "--network" && process.argv[5] === "bscdev";
  const NUM_ARGS = process.argv.length - 4;

  /*******************************************
   *** Command line argument error checking
   ***
   *** truffle exec lacks support for dynamic command line arguments:
   *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
   ******************************************/
  if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_BSCDEV) {
    if (NUM_ARGS !== 4) {
      return console.error(
          "Error: invalid number of parameters, please try again."
      );
    }
  } else {
    if (NUM_ARGS !== 2) {
      return console.error(
          "Error: must specify sender address, and amount."
      );
    }
  }

  /*******************************************
   *** Stake transaction parameters
   ******************************************/
  let ethSender = "";
  let amount = "";

  if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_BSCDEV) {
    ethSender = process.argv[6];
    amount = process.argv[7];
  } else {
    ethSender = process.argv[4];
    amount = process.argv[5];
  }

  /*******************************************
   *** Web3 provider
   *** Set contract provider based on --network flag
   ******************************************/
  let provider;
  if (NETWORK_ROPSTEN || NETWORK_BSCDEV) {
    provider = new HDWalletProvider(
        process.env.MNEMONIC,
        process.env.HDWALLET_PROVIDER
    );
  } else {
    provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
  }

  const web3 = new Web3(provider);
  stakingPoolContract.setProvider(web3.currentProvider);
  tokenContract.setProvider(web3.currentProvider);
  try {
    /*******************************************
     *** Contract interaction
     ******************************************/
    let _amount = web3.utils.toWei(amount)

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
    console.log("staking pool address: ", stakingPoolAddress, "\n");

    let instance = await tokenContract.at(tokenAddress)
    const { logs: approveLogs } = await instance.approve(stakingPoolAddress, _amount, {
      from: ethSender,
      value: 0,
      gas: 300000 // 300,000 Gwei
    });

    console.log("Sent approval...");

    // Get event logs
    const approveEvent = approveLogs.find(e => e.event === "Approval");

    // Parse event fields
    const approveLog = {
      owner: approveEvent.args.owner,
      spender: approveEvent.args.spender,
      value: Number(approveEvent.args.value)
    };

    console.log(approveLog);

    // Send stake transaction
    console.log("Connecting to contract....");
    const { logs: stakeLogs } = await stakingPoolContract.deployed().then(function (instance) {
      console.log("Connected to contract, sending stake...");
      return instance.stake(_amount, {
        from: ethSender,
        value: 0,
        gas: 300000 // 300,000 Gwei
      });
    });

    console.log("Sent stake...");

    // Get event logs
    const statkeEvent = stakeLogs.find(e => e.event === "LogStake");

    // Parse event fields
    const stakeLog = {
      staker: statkeEvent.args._staker,
      value: Number(statkeEvent.args._value),
      nonce: Number(statkeEvent.args._totalAmount)
    };

    console.log(stakeLog);
  } catch (error) {
    console.error({ error });
  }

  return;
};
