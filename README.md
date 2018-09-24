## EOS.go

A library to interact with EOS blockchain. Contains some other helpful functions.

## UNDER DEVELOPMENT

The library is quite raw and contains very basic features. Active development has just begun.

## Progress

### Chain API

- [x] get_info (`GetInfo`)
- [x] get_block (`GetBlockByNumber` & `GetBlockByID`)
- [ ] get_block_header_state
- [ ] get_account
- [ ] get_code
- [ ] get_abi
- [ ] get_raw_code_and_abi
- [ ] get_table_rows
- [ ] get_currency_balance
- [ ] get_currency_stats
- [ ] get_producers
- [ ] get_producer_schedule
- [ ] get_scheduled_transactions
- [ ] abi_json_to_bin
- [ ] abi_bin_to_json
- [ ] get_required_keys
- [ ] get_transaction_id
- [ ] push_block
- [ ] push_transaction
- [ ] push_transactions

### Wallet API

- [ ] set_timeout
- [ ] sign_transaction
- [ ] sign_digest
- [ ] create
- [ ] open
- [ ] lock_all
- [ ] lock
- [ ] unlock
- [ ] import_key
- [ ] remove_key
- [ ] create_key
- [ ] list_wallets
- [ ] list_keys
- [ ] get_public_keys

### Helper functions

- [ ] Some useful functions from `cleos`
- [ ] Quick call of system contract's actions (like `deletagebw`, `buyrambytes`, etc.)
- [ ] Quick call of token actions (like `transfer`, `issue`, etc.)
