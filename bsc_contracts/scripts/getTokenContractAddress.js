module.exports = async () => {
  /*******************************************
   *** Set up
   ******************************************/
  require("dotenv").config();
  const Web3 = require("web3");
  const HDWalletProvider = require("@truffle/hdwallet-provider");

  // Contract abstraction
  const truffleContract = require("truffle-contract");
  const erc20Contract = truffleContract(
    require("../build/contracts/BridgeERC20Token.json")
  );
  const erc721Contract = truffleContract(
      require("../build/contracts/BridgeERC721Token.json")
  );

  /*******************************************
   *** Constants
   ******************************************/
  const NETWORK_ROPSTEN =
    process.argv[4] === "--network" && process.argv[5] === "ropsten";
  const NETWORK_BSCDEV =
      process.argv[4] === "--network" && process.argv[5] === "bscdev";

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
  erc20Contract.setProvider(web3.currentProvider);
  erc721Contract.setProvider(web3.currentProvider);

  /*******************************************
   *** Contract interaction
   ******************************************/
  const erc20Address = await erc20Contract.deployed().then(function(instance) {
    return instance.address;
  });
  console.log("Token erc20Contract address: ", erc20Address);

  const erc721Address = await erc721Contract.deployed().then(function(instance) {
    return instance.address;
  });
  console.log("Token erc721Contract address: ", erc721Address);

  return true;
};
