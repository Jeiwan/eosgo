## EOS.go

A library to interact with EOS blockchain. Contains some other helpful functions.

## UNDER DEVELOPMENT

The library is quite raw and contains very basic features. Active development has just begun.

## Progress

### Chain API

- [ ] abi_bin_to_json
- [x] abi_json_to_bin (`ABIJSONtoBin`)
- [ ] get_abi
- [ ] get_account
- [x] get_block (`GetBlockByNumber` & `GetBlockByID`)
- [ ] get_block_header_state
- [ ] get_code
- [x] get_currency_balance (`GetCurrencyBalance`)
- [ ] get_currency_stats
- [x] get_info (`GetInfo`)
- [ ] get_producer_schedule
- [ ] get_producers
- [ ] get_raw_code_and_abi
- [ ] get_required_keys
- [ ] get_scheduled_transactions
- [ ] get_table_rows
- [ ] get_transaction_id
- [ ] push_block
- [x] push_transaction (`PushNotification`)
- [ ] push_transactions

### Wallet API

- [ ] create
- [ ] create_key
- [ ] get_public_keys
- [ ] import_key
- [ ] list_keys
- [ ] list_wallets
- [ ] lock
- [ ] lock_all
- [ ] open
- [ ] remove_key
- [ ] set_timeout
- [ ] sign_digest
- [x] sign_transaction (`SignTransaction`)
- [ ] unlock

### Helper functions

- [ ] Some useful functions from `cleos`
- [ ] Quick call of system contract's actions (like `deletagebw`, `buyrambytes`, etc.)
- [ ] Quick call of token actions (like `transfer`, `issue`, etc.)
