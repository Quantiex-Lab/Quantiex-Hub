module.exports = async () => {
    /*******************************************
     *** Set up
     ******************************************/
    require("dotenv").config();
    const Web3 = require("web3");
    const HDWalletProvider = require("@truffle/hdwallet-provider");
    const BigNumber = require("bignumber.js")

    // Contract abstraction
    const truffleContract = require("truffle-contract");
    const contract = truffleContract(
        require("../build/contracts/Valset.json")
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
    contract.setProvider(web3.currentProvider);
    try {
        // Get current accounts
        const accounts = await web3.eth.getAccounts();

        /*******************************************
         *** Contract interaction
         ******************************************/
        await contract.deployed().then(async function (instance) {
            for (let i = 0; i < accounts.length; i++) {
                console.log("Trying " + accounts[i] + "...")
                const isValidator = await instance.isActiveValidator(accounts[i], {
                    from: operator,
                    value: 0,
                    gas: 300000 // 300,000 Gwei
                });
                if (isValidator) {
                    const power = new BigNumber(await instance.getValidatorPower(accounts[i], {
                        from: operator,
                        value: 0,
                        gas: 300000 // 300,000 Gwei
                    }));
                    console.log("Validator " + accounts[i] + " is active! Power:", power.c[0])
                }
            }
        });
    } catch (error) {
        console.error({ error })
    }
};
