require("dotenv").config();

const Valset = artifacts.require("Valset");
const QuantiexERC20Bridge = artifacts.require("QuantiexERC20Bridge");
const QuantiexERC721Bridge = artifacts.require("QuantiexERC721Bridge");
const Oracle = artifacts.require("Oracle");
const BridgeERC20Bank = artifacts.require("BridgeERC20Bank");
const BridgeERC721Bank = artifacts.require("BridgeERC721Bank");
const BridgeRegistry = artifacts.require("BridgeRegistry");
const BridgeERC20Token = artifacts.require("BridgeERC20Token");
const BridgeERC721Token = artifacts.require("BridgeERC721Token");
const ERC721TokenFactory = artifacts.require("ERC721TokenFactory");
const StakingPool = artifacts.require("StakingPool");

module.exports = function(deployer, network, accounts) {
  /*******************************************
   *** Input validation of contract params
   ******************************************/
  let operator;
  let initialValidators = [];
  let consensusThreshold;

  // Input validation for general env variables
  if (process.env.CONSENSUS_THRESHOLD.length === 0) {
    return console.error(
      "Must provide consensus threshold as environment variable."
    );
  }
  consensusThreshold = process.env.CONSENSUS_THRESHOLD;

  // Input validation for local usage (develop, ganache)
  if (network === "develop" || network === "ganache") {
    // Operator
    if (process.env.LOCAL_OPERATOR.length === 0) {
      return console.error(
          "Must provide operator address as environment variable."
      );
    }
    // Initial validators
    if (process.env.LOCAL_INITIAL_VALIDATOR_ADDRESSES.length === 0) {
      return console.error(
          "Must provide initial validator addresses as environment variable."
      );
    }

    // Assign validated local input params
    operator = process.env.LOCAL_OPERATOR;
    initialValidators = process.env.LOCAL_INITIAL_VALIDATOR_ADDRESSES.split(",");
  } else {
    // Operator
    if (process.env.OPERATOR.length === 0) {
      return console.error(
        "Must provide operator address as environment variable."
      );
    }
    // Initial validators
    if (process.env.INITIAL_VALIDATOR_ADDRESSES.length === 0) {
      return console.error(
        "Must provide initial validator addresses as environment variable."
      );
    }

    // Assign validated testnet/mainnet input params
    operator = process.env.OPERATOR;
    initialValidators = process.env.INITIAL_VALIDATOR_ADDRESSES.split(",");
  }

  /*******************************************************
   *** Contract deployment summary
   ***
   *** Total deployments:       7 (includes Migrations.sol)
   *** Gas price (default):                       20.0 Gwei
   *** Final cost:                         0.25369878 Ether
   *******************************************************/
  deployer.then(async () => {
    // 1. Deploy BridgeERC20Token contract
    //    Gas used:        1,884,394 Gwei
    //    Total cost:    0.03768788 Ether
    await deployer.deploy(BridgeERC20Token, "ERC20COIN", {
      gas: 4612388,
      from: operator
    });

    // 2. Deploy BridgeERC721Token contract
    //    Gas used:        1,884,394 Gwei
    //    Total cost:    0.03768788 Ether
    await deployer.deploy(BridgeERC721Token, operator, "ERC721COIN", "http://xx.xx.com/", {
      gas: 4612388,
      from: operator
    });

    // 3. Deploy Valset contract:
    //    Gas used:          909,879 Gwei
    //    Total cost:    0.01819758 Ether
    await deployer.deploy(Valset, operator, initialValidators, {
      gas: 6721975,
      from: operator
    });

    // 4. Deploy QuantiexERC20Bridge contract:
    //    Gas used:       2,649,300 Gwei
    //    Total cost:     0.052986 Ether
    await deployer.deploy(QuantiexERC20Bridge, operator, Valset.address, {
      gas: 6721975,
      from: operator
    });

    // 5. Deploy QuantiexERC721Bridge contract:
    //    Gas used:       2,649,300 Gwei
    //    Total cost:     0.052986 Ether
    await deployer.deploy(QuantiexERC721Bridge, operator, Valset.address, {
      gas: 6721975,
      from: operator
    });

    // 6. Deploy StakingPool contract:
    //    Gas used:       2,649,300 Gwei
    //    Total cost:     0.052986 Ether
    await deployer.deploy(StakingPool, operator, {
      gas: 6721975,
      from: operator
    });

    // 7. Deploy Oracle contract:
    //    Gas used:        1,769,740 Gwei
    //    Total cost:     0.0353948 Ether
    await deployer.deploy(
      Oracle,
      operator,
      Valset.address,
      QuantiexERC20Bridge.address,
      QuantiexERC721Bridge.address,
      StakingPool.address,
      consensusThreshold,
      {
        gas: 6721975,
        from: operator
      }
    );

    // 8. Deploy BridgeERC20Bank contract:
    //    Gas used:        4,823,348 Gwei
    //    Total cost:    0.09646696 Ether
    await deployer.deploy(
      BridgeERC20Bank,
      operator,
      QuantiexERC20Bridge.address,
      {
        gas: 6721975,
        from: operator
      }
    );

    // 9. Deploy BridgeERC721Bank contract:
    //    Gas used:        4,823,348 Gwei
    //    Total cost:    0.09646696 Ether
    await deployer.deploy(
        ERC721TokenFactory,
        {
          gas: 6721975,
          from: operator
        }
    );

    // 10. Deploy BridgeERC721Bank contract:
    //    Gas used:        4,823,348 Gwei
    //    Total cost:    0.09646696 Ether
    await deployer.deploy(
        BridgeERC721Bank,
        operator,
        QuantiexERC721Bridge.address,
        ERC721TokenFactory.address,
        {
          gas: 6721975,
          from: operator
        }
    );

    // 11. Deploy BridgeRegistry contract:
    //    Gas used:          363,370 Gwei
    //    Total cost:     0.0072674 Ether
    return deployer.deploy(
      BridgeRegistry,
      QuantiexERC20Bridge.address,
      QuantiexERC721Bridge.address,
      BridgeERC20Bank.address,
      BridgeERC721Bank.address,
      Oracle.address,
      Valset.address,
      StakingPool.address,
      {
        gas: 6721975,
        from: operator
      }
    );
  });
};
