# BLOCK VARIABLES

Today we will cover the analogs of all the block variables from solidity. Not all of them have 1-1 analogs. In solidity, we have the following commonly used block variables:

- block.timestamp
- block.number
- block.hash()

And the lesser known ones:

- block.coinbase
- block.gaslimit
- block.basefee
- block.difficulty

Let's go through one by one, create a new program first.

```
anchor new day_11
```



## block.timestamp

By utilizing the `unix_timestamp` field within the [Clock sysvar](https://docs.solanalabs.com/runtime/sysvars), we can access the block timestamp Solana.

```rust
use anchor_lang::prelude::*;

declare_id!("2SCX7gj5ByjWvin16BNNDjZxEcsikX15DWMMccc7EVXD");

#[program]
pub mod day_11 {
    use super::*;

    pub fn get_timestamp(ctx: Context<Initialize>) -> Result<()> {
        let clock: Clock = Clock::get()?;
        msg!("Current Unix Timestamp: {}", clock.unix_timestamp);
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize {}

```

test:

![image-20240803101113018](./assets/image-20240803101113018.png)

let's div further to get the `day of the week`, we need to use a new  carte:  [chrono](https://docs.rs/chrono/latest/chrono/), it provides functionality for operations on dates and times in Rust.

```toml
# day_11/Cargo.toml
[dependencies]
anchor-lang = "0.30.1"
chrono = "0.4.31"
```

update code, be sure to import `use chrono::*`

```rust
pub fn get_day_of_week(ctx: Context<Initialize>) -> Result<()> {
        let clock: Clock = Clock::get()?;
        let time_stamp = clock.unix_timestamp;

        // let date_time = chrono::NaiveDateTime::from_timestamp(time_stamp, 0); // DEPRECATED!!
        let date_time = DateTime::from_timestamp(time_stamp, 0);
        msg!(
            "Day of the week: {:?}",
            date_time.expect("REASON").weekday() // ATTENTION!!
        );
        Ok(())
    }
```

output:

![image-20240803105052091](./assets/image-20240803105052091.png)



## block.number

solana has a notion of a `slot number` which is very related to the `block number`, we will defer a full discussion of how to get the block number in later tutorials.

## block.coinbase

In Ethereum, the `block.coinbase` represents the address of the minter who has successfully minted a block in Proof of Work. 

On the other hand, solana uses a `leader-based` consensus mechanism which is a combination of both Proof of History(**POH**) and Proof of Stake(**POS**), removing the concept of mining.

Instead, [a block or slot leader](https://docs.solanalabs.com/consensus/leader-rotation) is appointed to validate transactions and propose blocks during certain intervals, user a system known as the leader schedule. This schedule determines who will be the block producer at a certain time.

However, presently, there's **no specific way** to access the address of the block leader in Solana programs.

## block.hash

We include this section for completeness. There is no way to get hash in the future, the current function would be deprecated soon.

## block.gaslimit

Solana has a per-block compute unit limit of 48 million. Each transaction is by default limited to 200,000 compute units, though it can be raised to 1.4 million computes(disscuss later).

## block.basefee

In Ethereum, the basefee is dynamice due EIP-1559, it is a function of pervious block utilization.

In Solana, the base price of a transaction is static, so there is no need for a variable like this.

## block.difficulty

Block difficulty is a concept associated with Proof of Work(PoW) blockchains.

Solana, on the other hand, operates on a Proof of History(PoH) combined with Proof of Stake(Pos) consensus mechanism, which doesn't involve the concept of block difficulty.

## block.chainId

Solana doesn't have a chain id because it is not an EVM compatible chain.

Solana runs separate clusters for Devnet, Testnet, and Mainnet, but programs donot have a mechanism to know which one tye are on.

We can programatically adjust your code at deploy time [using the Rust cfg](https://solana.stackexchange.com/questions/848/how-to-have-a-different-program-id-depending-on-the-cluster) feature to have different features depending on which cluster it is deployed to.



## Key Takeaways

- block.timestamp => Clock
- block.number => slot number
- **block.coinbase => N/A**
- **block.hash => N/A**
- block.gaslimit => 48million compute units
- **block.basefee => N/A**
- **block.difficulty => N/A**
- **block.chainid => N/A**



## Links

- day_11 original article: https://www.rareskills.io/post/solana-clock
- source code: https://github.com/dukedaily/solana-expert-code/tree/day_11