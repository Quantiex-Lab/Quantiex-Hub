require("dotenv").config();

var HDWalletProvider = require("@truffle/hdwallet-provider");

module.exports = {
  // contracts_directory: "./flat",
  networks: {
    develop: {
      host: "localhost",
      port: 8546, // Match default network 'ganache'
      network_id: 5777,
      gas: 6721975, // Truffle default development block gas limit
      gasPrice: 200000000000,
      solc: {
        version: "0.5.0",
        optimizer: {
          enabled: true,
          runs: 200
        }
      }
    },
    ropsten: {
      provider: function () {
        return new HDWalletProvider(
          process.env.MNEMONIC,
            process.env.HDWALLET_PROVIDER
        );
      },
      network_id: 3,
      gas: 6000000
    },
    bscdev: {
      provider: function () {
        return new HDWalletProvider(
            process.env.MNEMONIC,
            process.env.HDWALLET_PROVIDER
        );
      },
      network_id: 16,
      gas: 6000000,
      gasPrice: 1000000000
    },
    xdai: {
      provider: function () {
        return new HDWalletProvider(
          process.env.MNEMONIC,
          "https://dai.poa.network"
        );
      },
      network_id: 100,
      gas: 6000000
    }
  },
  rpc: {
    host: "localhost",
    post: 8080
  },
  mocha: {
    useColors: true
  }
};
