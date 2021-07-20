module.exports = async () => {
    /*******************************************
     *** Set up
     ******************************************/
    require("dotenv").config();
    const Web3 = require("web3");
    const HDWalletProvider = require("@truffle/hdwallet-provider");
    try {

        // Contract abstraction
        const truffleContract = require("truffle-contract");
        const tokenContract = truffleContract(
          require("../build/contracts/BridgeERC20Token.json")
        );

        /*******************************************
         *** Constants
         ******************************************/
        const NETWORK_ROPSTEN =
            process.argv[4] === "--network" && process.argv[5] === "ropsten";
        const NETWORK_DEVELOP =
            process.argv[4] === "--network" && process.argv[5] === "develop";

        let sender;
        let recipient
        let token;
        let amount;

        if (NETWORK_ROPSTEN || NETWORK_DEVELOP) {
            sender = process.argv[6];
            recipient = process.argv[7];
            token = process.argv[8];
            amount = process.argv[9];
        } else {
            sender = process.argv[4];
            recipient = process.argv[5];
            token = process.argv[6];
            amount = process.argv[7];
        }

        if (!sender) {
            console.log("Please provide an Ethereum address to check their balance")
            return
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
        /*******************************************
         *** Contract interaction
         ******************************************/
        let _amount = web3.utils.toWei(amount)

        if (token === "eth")
        {
            let txObj = {
                "from":sender,
                "to":recipient,
                "value":_amount
            }

            web3.eth.sendTransaction(txObj, function (error, result) {
                if (error) {
                    console.log("Transaction error:", error);
                } else {
                    console.log("Result:", result);
                }
            })
        } else {
            tokenContract.setProvider(web3.currentProvider);

            let instance = await tokenContract.at(token)
            const { logs } = await instance.transfer(recipient, _amount, {
                from: sender,
                value: 0,
                gas: 300000 // 300,000 Gwei
            });

            console.log("Sent transfer...");

            // Get event logs
            const transferEvent = logs.find(e => e.event === "Transfer");

            // Parse event fields
            const transferLog = {
                sender: transferEvent.args.sender,
                recipient: transferEvent.args.recipient,
                amount: Number(transferEvent.args.amount)
            };

            console.log(transferLog);
        }

    } catch (error) {
        console.error({error})
    }

    return null;
};
  