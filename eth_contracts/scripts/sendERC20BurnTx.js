module.exports = async () => {
    /*******************************************
     *** Set up
     ******************************************/
    const Web3 = require("web3");
    const HDWalletProvider = require("@truffle/hdwallet-provider");

    // Contract abstraction
    const truffleContract = require("truffle-contract");
    const contract = truffleContract(
        require("../build/contracts/BridgeERC20Bank.json")
    );
    const tokenContract = truffleContract(
        require("../build/contracts/BridgeERC20Token.json")
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
    const NETWORK_ETHDEV =
        process.argv[4] === "--network" && process.argv[5] === "ethdev";
    const NUM_ARGS = process.argv.length - 4;

    /*******************************************
     *** Command line argument error checking
     ***
     *** truffle exec lacks support for dynamic command line arguments:
     *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
     ******************************************/
    if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
        if (NUM_ARGS !== 6) {
            return console.error(
                "Error: invalid number of parameters, please try again."
            );
        }
    } else {
        if (NUM_ARGS !== 4) {
            return console.error(
                "Error: must specify recipient address, token address, and amount."
            );
        }
    }

    /*******************************************
     *** Burn transaction parameters
     ******************************************/
    let bscSender = "";
    let ethRecipient = "";
    let coinDenom = "";
    let amount = "";

    if (NETWORK_ROPSTEN || NETWORK_DEVELOP || NETWORK_ETHDEV) {
        bscSender = process.argv[6];
        ethRecipient = process.argv[7];
        coinDenom = process.argv[8];
        amount = process.argv[9];
    } else {
        bscSender = process.argv[4];
        ethRecipient = process.argv[5];
        coinDenom = process.argv[6];
        amount = process.argv[7];
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
    } else if (NETWORK_ETHDEV) {
        provider = new HDWalletProvider(
            process.env.MNEMONIC,
            process.env.HDWALLET_PROVIDER
        );
    } else {
        provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
    }

    const web3 = new Web3(provider);
    contract.setProvider(web3.currentProvider);
    tokenContract.setProvider(web3.currentProvider);

    try {
        /*******************************************
         *** Contract interaction
         ******************************************/
        let _mount = web3.utils.toWei(amount)

        // Send approve transaction
        if (coinDenom === "eth") {
            coinDenom = NULL_ADDRESS;
        } else  {
            const bridgeContractAddress = await contract
                .deployed()
                .then(function(instance) {
                    return instance.address;
                });

            let instance = await tokenContract.at(coinDenom)
            const { logs } = await instance.approve(bridgeContractAddress, _mount, {
                from: bscSender,
                value: 0,
                gas: 300000 // 300,000 Gwei
            });

            console.log("Sent approval...");

            // Get event logs
            const eventA = logs.find(e => e.event === "Approval");

            // Parse event fields
            const approvalEvent = {
                owner: eventA.args.owner,
                spender: eventA.args.spender,
                value: Number(eventA.args.value)
            };

            console.log(approvalEvent);
        }

        // Send Burn transaction
        console.log("Connecting to contract....");
        const { logs: logs2 } = await contract.deployed().then(function (instance) {
            console.log("Connected to contract, sending burn...");
            return instance.burn(ethRecipient, coinDenom, _mount, {
                from: bscSender,
                value: coinDenom === NULL_ADDRESS ? _mount : 0,
                gas: 300000 // 300,000 Gwei
            });
        });

        console.log("Sent burn...");

        // Get event logs
        const eventB = logs2.find(e => e.event === "LogBurn");

        // Parse event fields
        const burnEvent = {
            chainName: eventB.args._chainName,
            from: eventB.args._from,
            to: eventB.args._to,
            token: eventB.args._token,
            symbol: eventB.args._symbol,
            value: Number(eventB.args._value),
            nonce: Number(eventB.args._nonce)
        };
        console.log(burnEvent);

    } catch (error) {
        console.error({ error });
    }
    return;
};

