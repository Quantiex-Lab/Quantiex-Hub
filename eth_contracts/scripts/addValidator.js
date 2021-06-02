module.exports = async () => {
  /*******************************************
   *** Set up
   ******************************************/
  const Web3 = require("web3");
  const HDWalletProvider = require("@truffle/hdwallet-provider");

  // Contract abstraction
  const truffleContract = require("truffle-contract");
  const valsetContract = truffleContract(
      require("../build/contracts/Valset.json")
  );

  /*******************************************
   *** Constants
   ******************************************/
      // Config values
  const NETWORK_ROPSTEN =
      process.argv[4] === "--network" && process.argv[5] === "ropsten";
  const NETWORK_DEVELOP =
      process.argv[4] === "--network" && process.argv[5] === "develop";
  const NUM_ARGS = process.argv.length - 4;

  /*******************************************
   *** Command line argument error checking
   ***
   *** truffle exec lacks support for dynamic command line arguments:
   *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
   ******************************************/
  if (NETWORK_ROPSTEN || NETWORK_DEVELOP) {
    if (NUM_ARGS !== 3) {
      return console.error(
          "Error: invalid number of parameters, please try again."
      );
    }
  } else {
    if (NUM_ARGS !== 1) {
      return console.error(
          "Error: must specify recipient address, token address, and amount."
      );
    }
  }

  /*******************************************
   *** Stake transaction parameters
   ******************************************/
  let validator = "";

  if (NETWORK_ROPSTEN || NETWORK_DEVELOP) {
    validator = process.argv[6];
  } else {
    validator = process.argv[4];
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
  } else {
    provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
    operator = process.env.LOCAL_OPERATOR;
  }

  const web3 = new Web3(provider);
  valsetContract.setProvider(web3.currentProvider);
  try {
    /*******************************************
     *** Contract interaction
     ******************************************/
    console.log("Connecting to contract....");
    const { logs } = await valsetContract.deployed().then(function (instance) {
      console.log("Connected to contract, sending stake...");
      return instance.addValidator(validator, {
        from: operator,
        value: 0,
        gas: 300000 // 300,000 Gwei
      });
    });

    console.log("Sent validator...");

    // Get event logs
    const event = logs.find(e => e.event === "LogValidatorAdded");

    // Parse event fields
    const addEvent = {
      validator: event.args._validator,
      currentValsetVersion: Number(event.args._currentValsetVersion),
      validatorCount: Number(event.args._validatorCount)
    };

    console.log(addEvent);
  } catch (error) {
    console.error({ error });
  }
  return;
};
