# TxFee Compute Unit

In Ethereum, the fee of a transaction is computed as `gasPrice * gasUsed`. This tells us how much in Eth will be needed, and there is a cost ceiling, gasLimit, which will protect the caller from drain his wallet by any kind of issues, such as infinite loops, the transaction will revert if it runs out of gas.

## Compute units

Unlike on EVM chains, Solana opcodes(instructions) consume `compute units`, not gas, and each transaction is limited to 200,000 compute unites and reverts if exceed.



**Different:**

- In Ethereum, gas costs of computing are treated the same as gas costs associated with storage.
- In Solana, storage is handled differently, so the pricing of persistent data in solana is a different topic of discussion.

**Same:**

- From the perspective of pricing running opcodes, however, Ethereum and Solana behave similarly.

Both chains execute compiled bytecode and charge a fee for each instruction executed. Ethereum uses EVM bytecode, but Solana runs a modified version of **berkeley packet filter**, called Solana packet filter.

Ethereum charges different prices for different opcodes depending on how long they take to execute, ranging from one gas to thousands of gas, **however**, each opcode costs one compute unit in solana.





## Key Takeaways

- 
- 

## Links

- original article: https://www.rareskills.io/post/solana-compute-unit-price
- source code: https://github.com/dukedaily/solana-expert-code/tree/day_15
