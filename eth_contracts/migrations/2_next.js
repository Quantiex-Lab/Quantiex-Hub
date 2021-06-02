require("dotenv").config();

const Valset = artifacts.require("Valset");
const QuantiexBridge = artifacts.require("QuantiexBridge");
const Oracle = artifacts.require("Oracle");
const BridgeBank = artifacts.require("BridgeBank");
const BridgeRegistry = artifacts.require("BridgeRegistry");
const BridgeToken = artifacts.require("BridgeToken");
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
    // 1. Deploy BridgeToken contract
    //    Gas used:        1,884,394 Gwei
    //    Total cost:    0.03768788 Ether
    await deployer.deploy(BridgeToken, "COIN", {
      gas: 4612388,
      from: operator
    });

    // 2. Deploy Valset contract:
    //    Gas used:          909,879 Gwei
    //    Total cost:    0.01819758 Ether
    await deployer.deploy(Valset, operator, initialValidators, {
      gas: 6721975,
      from: operator
    });

    // 3. Deploy QuantiexBridge contract:
    //    Gas used:       2,649,300 Gwei
    //    Total cost:     0.052986 Ether
    await deployer.deploy(QuantiexBridge, operator, Valset.address, {
      gas: 6721975,
      from: operator
    });

    // 4. Deploy StakingPool contract:
    //    Gas used:       2,649,300 Gwei
    //    Total cost:     0.052986 Ether
    await deployer.deploy(StakingPool, operator, {
      gas: 6721975,
      from: operator
    });

    // 5. Deploy Oracle contract:
    //    Gas used:        1,769,740 Gwei
    //    Total cost:     0.0353948 Ether
    await deployer.deploy(
      Oracle,
      operator,
      Valset.address,
      QuantiexBridge.address,
      StakingPool.address,
      consensusThreshold,
      {
        gas: 6721975,
        from: operator
      }
    );

    // 6. Deploy BridgeBank contract:
    //    Gas used:        4,823,348 Gwei
    //    Total cost:    0.09646696 Ether
    await deployer.deploy(
      BridgeBank,
      operator,
      QuantiexBridge.address,
      {
        gas: 6721975,
        from: operator
      }
    );

    // 7. Deploy BridgeRegistry contract:
    //    Gas used:          363,370 Gwei
    //    Total cost:     0.0072674 Ether
    return deployer.deploy(
      BridgeRegistry,
      QuantiexBridge.address,
      BridgeBank.address,
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
