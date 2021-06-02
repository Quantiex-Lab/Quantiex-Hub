module.exports = async () => {
  /*******************************************
   *** Set up
   ******************************************/
  require("dotenv").config();
  const Web3 = require("web3");
  const BigNumber = require("bignumber.js")
  const HDWalletProvider = require("@truffle/hdwallet-provider");
  try {

  // Contract abstraction
  const truffleContract = require("truffle-contract");
  const contract = truffleContract(
    require("../build/contracts/StakingPool.json")
  );

  /*******************************************
   *** Constants
   ******************************************/
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
            "Error: must specify validator address."
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
  if (NETWORK_ROPSTEN) {
    provider = new HDWalletProvider(
      process.env.MNEMONIC,
      "https://ropsten.infura.io/v3/".concat(process.env.INFURA_PROJECT_ID)
    );
  } else {
    provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
  }

  const web3 = new Web3(provider);
  contract.setProvider(web3.currentProvider);
  /*******************************************
   *** Contract interaction
   ******************************************/
  const tokenAmount = await contract.deployed().then(function(instance) {
    return instance.weightOf(validator);
  });

  return console.log("Staking token amount: ", (new BigNumber(tokenAmount).div(new BigNumber(10).pow(18))).toString(10));
} catch (error) {
  console.error({error})
}
};
