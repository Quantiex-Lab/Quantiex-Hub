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
    if (NUM_ARGS !== 3) {
      return console.error(
          "Error: invalid number of parameters, please try again."
      );
    }
  } else {
    if (NUM_ARGS !== 1) {
      return console.error(
          "Error: must specify token address."
      );
    }
  }

  /*******************************************
   *** Lock transaction parameters
   ******************************************/
  let tokenAddress;

  if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_BSCDEV) {
    tokenAddress = process.argv[6];
  } else {
    tokenAddress = process.argv[4];
  }

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
  stakingPoolContract.setProvider(web3.currentProvider);
  try {

    /*******************************************
     *** Contract interaction
     ******************************************/
    // Set Oracle
    const { logs: setTokenAddressLogs } = await stakingPoolContract
      .deployed()
      .then(function(instance) {
        return instance.setTokenAddress(tokenAddress, {
          from: operator,
          value: 0,
          gas: 300000 // 300,000 Gwei
        });
      });
    // Get event logs
    const setTokenAddressEvent = setTokenAddressLogs.find(e => e.event === "LogSetTokenAddress");
    console.log("Staking token address:", setTokenAddressEvent.args.token);

  } catch (error) {
    console.error({error})
  }

  return;
};
