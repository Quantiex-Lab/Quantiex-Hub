module.exports = async () => {
  /*******************************************
   *** Set up
   ******************************************/
  const Web3 = require("web3");
  const HDWalletProvider = require("@truffle/hdwallet-provider");

  // Contract abstraction
  const truffleContract = require("truffle-contract");
  const contract = truffleContract(
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
  let staker = "";
  let amount = "";

  if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
    staker = process.argv[6];
    amount = process.argv[7];
  } else {
    staker = process.argv[4];
    amount = process.argv[5];
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
    operator = process.env.LOCAL_OPERATOR;
  } else {
    provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
    operator = process.env.LOCAL_OPERATOR;
  }

  const web3 = new Web3(provider);
  contract.setProvider(web3.currentProvider);
  tokenContract.setProvider(web3.currentProvider);
  try {
    /*******************************************
     *** Contract interaction
     ******************************************/
    let _amount = web3.utils.toWei(amount)

    // Send stake transaction
    console.log("Connecting to contract....");
    const { logs } = await contract.deployed().then(function (instance) {
      console.log("Connected to contract, sending revoke...");
      return instance.revoke(staker, _amount, {
        chainId: 5777,
        from: operator,
        value: 0,
        gas: 300000 // 300,000 Gwei
      });
    });

    console.log("Sent revoke...");

    // Get event logs
    const event = logs.find(e => e.event === "LogRevoke");

    // Parse event fields
    const revokeEvent = {
      staker: event.args._staker,
      value: Number(event.args._value),
      nonce: Number(event.args._totalAmount)
    };

    console.log(revokeEvent);

  } catch (error) {
    console.error({ error });
  }
  return;
};
