package rpc

import (
	"fmt"
	"TVStorageManager/json"
	"net"
	"TVStorageManager/logic"
	"TVStorageManager/network"
)


// passThroughTable
type RpcInterFace struct {
	Method      	string
	Func 			func(request *json.Json, tvConn net.Conn)(response string)
	IsPassThrough	bool
}

var rpcInterFaceTable map[string]RpcInterFace
func InitRpcInterFaces() {

	rpcInterFaceTable = make(map[string]RpcInterFace)

	rpcInterFaceTable["UploadFileToIPFS"] = RpcInterFace{"UploadFileToIPFS", UploadFileToIPFS, false}
	rpcInterFaceTable["DeclareUploadFile"] = RpcInterFace{"DeclareUploadFile", DeclareUploadFile, false}
	rpcInterFaceTable["ListUploadDeclaration"] = RpcInterFace{"ListUploadDeclaration", ListUploadDeclaration, false}
	rpcInterFaceTable["PinAddFileToLocal"] = RpcInterFace{"PinAddFileToLocal", PinAddFileToLocal, false}
	rpcInterFaceTable["DeclarePieceSaved"] = RpcInterFace{"DeclarePieceSaved", DeclarePieceSaved, false}
	rpcInterFaceTable["ListFileStoreRequest"] = RpcInterFace{"ListFileStoreRequest", ListFileStoreRequest, false}
	rpcInterFaceTable["ConfirmFileSaved"] = RpcInterFace{"ConfirmFileSaved", ConfirmFileSaved, false}
	rpcInterFaceTable["ListConfirmFileSaved"] = RpcInterFace{"ListConfirmFileSaved", ListConfirmFileSaved, false}

	// the tv rpc interface for pass through
	rpcInterFaceTable["about"] = RpcInterFace{"about", RpcPassThrough, true}
	rpcInterFaceTable["add_event_handler"] = RpcInterFace{"add_event_handler", RpcPassThrough, true}
	rpcInterFaceTable["add_script"] = RpcInterFace{"add_script", RpcPassThrough, true}
	rpcInterFaceTable["batch"] = RpcInterFace{"batch", RpcPassThrough, true}
	rpcInterFaceTable["batch_authenticated"] = RpcInterFace{"batch_authenticated", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_broadcast_transaction"] = RpcInterFace{"blockchain_broadcast_transaction", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_btc_address_convert"] = RpcInterFace{"blockchain_btc_address_convert", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_calculate_supply"] = RpcInterFace{"blockchain_calculate_supply", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_check_signature"] = RpcInterFace{"blockchain_check_signature", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_dump_state"] = RpcInterFace{"blockchain_dump_state", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_export_fork_graph"] = RpcInterFace{"blockchain_export_fork_graph", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_generate_issuance_map"] = RpcInterFace{"blockchain_generate_issuance_map", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_generate_snapshot"] = RpcInterFace{"blockchain_generate_snapshot", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get__upload_requests"] = RpcInterFace{"blockchain_get__upload_requests", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_account"] = RpcInterFace{"blockchain_get_account", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_account_public_balance"] = RpcInterFace{"blockchain_get_account_public_balance", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_all_contracts"] = RpcInterFace{"blockchain_get_all_contracts", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_asset"] = RpcInterFace{"blockchain_get_asset", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_balance"] = RpcInterFace{"blockchain_get_balance", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_block"] = RpcInterFace{"blockchain_get_block", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_block_count"] = RpcInterFace{"blockchain_get_block_count", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_block_signee"] = RpcInterFace{"blockchain_get_block_signee", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_block_transactions"] = RpcInterFace{"blockchain_get_block_transactions", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_delegate_slot_entrys"] = RpcInterFace{"blockchain_get_delegate_slot_entrys", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_file_authorizing_contract"] = RpcInterFace{"blockchain_get_file_authorizing_contract", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_file_save_node"] = RpcInterFace{"blockchain_get_file_save_node", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_forever_contracts"] = RpcInterFace{"blockchain_get_forever_contracts", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_info"] = RpcInterFace{"blockchain_get_info", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_node_vm_enabled"] = RpcInterFace{"blockchain_get_node_vm_enabled", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_pretty_contract_transaction"] = RpcInterFace{"blockchain_get_pretty_contract_transaction", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_pretty_transaction"] = RpcInterFace{"blockchain_get_pretty_transaction", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_security_state"] = RpcInterFace{"blockchain_get_security_state", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_slate"] = RpcInterFace{"blockchain_get_slate", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_transaction"] = RpcInterFace{"blockchain_get_transaction", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_get_transaction_rpc"] = RpcInterFace{"blockchain_get_transaction_rpc", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_is_synced"] = RpcInterFace{"blockchain_is_synced", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_accounts"] = RpcInterFace{"blockchain_list_accounts", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_active_delegates"] = RpcInterFace{"blockchain_list_active_delegates", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_address_balances"] = RpcInterFace{"blockchain_list_address_balances", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_address_transactions"] = RpcInterFace{"blockchain_list_address_transactions", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_assets"] = RpcInterFace{"blockchain_list_assets", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_balances"] = RpcInterFace{"blockchain_list_balances", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_blocks"] = RpcInterFace{"blockchain_list_blocks", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_delegates"] = RpcInterFace{"blockchain_list_delegates", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_file_save_declare"] = RpcInterFace{"blockchain_list_file_save_declare", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_file_saved"] = RpcInterFace{"blockchain_list_file_saved", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_forks"] = RpcInterFace{"blockchain_list_forks", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_key_balances"] = RpcInterFace{"blockchain_list_key_balances", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_missing_block_delegates"] = RpcInterFace{"blockchain_list_missing_block_delegates", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_pending_transactions"] = RpcInterFace{"blockchain_list_pending_transactions", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_pub_all_address"] = RpcInterFace{"blockchain_list_pub_all_address", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_recently_registered_accounts"] = RpcInterFace{"blockchain_list_recently_registered_accounts", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_list_recently_updated_accounts"] = RpcInterFace{"blockchain_list_recently_updated_accounts", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_set_node_vm_enabled"] = RpcInterFace{"blockchain_set_node_vm_enabled", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_unclaimed_genesis"] = RpcInterFace{"blockchain_unclaimed_genesis", RpcPassThrough, true}
	rpcInterFaceTable["blockchain_verify_signature"] = RpcInterFace{"blockchain_verify_signature", RpcPassThrough, true}
	rpcInterFaceTable["builder_finalize_and_sign"] = RpcInterFace{"builder_finalize_and_sign", RpcPassThrough, true}
	rpcInterFaceTable["call_contract"] = RpcInterFace{"call_contract", RpcPassThrough, true}
	rpcInterFaceTable["call_contract_offline"] = RpcInterFace{"call_contract_offline", RpcPassThrough, true}
	rpcInterFaceTable["call_contract_testing"] = RpcInterFace{"call_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["compile_contract"] = RpcInterFace{"compile_contract", RpcPassThrough, true}
	rpcInterFaceTable["compile_script"] = RpcInterFace{"compile_script", RpcPassThrough, true}
	rpcInterFaceTable["compile_script"] = RpcInterFace{"compile_script", RpcPassThrough, true}
	rpcInterFaceTable["confirm_piece_saved"] = RpcInterFace{"confirm_piece_saved", RpcPassThrough, true}
	rpcInterFaceTable["debug_get_client_name"] = RpcInterFace{"debug_get_client_name", RpcPassThrough, true}
	rpcInterFaceTable["declare_piece_saved"] = RpcInterFace{"declare_piece_saved", RpcPassThrough, true}
	rpcInterFaceTable["delegate_blacklist_add_operation"] = RpcInterFace{"delegate_blacklist_add_operation", RpcPassThrough, true}
	rpcInterFaceTable["delegate_blacklist_add_transaction"] = RpcInterFace{"delegate_blacklist_add_transaction", RpcPassThrough, true}
	rpcInterFaceTable["delegate_blacklist_remove_operation"] = RpcInterFace{"delegate_blacklist_remove_operation", RpcPassThrough, true}
	rpcInterFaceTable["delegate_blacklist_remove_transaction"] = RpcInterFace{"delegate_blacklist_remove_transaction", RpcPassThrough, true}
	rpcInterFaceTable["delegate_get_config"] = RpcInterFace{"delegate_get_config", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_block_max_size"] = RpcInterFace{"delegate_set_block_max_size", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_block_max_transaction_count"] = RpcInterFace{"delegate_set_block_max_transaction_count", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_imessage_fee_coe"] = RpcInterFace{"delegate_set_imessage_fee_coe", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_network_min_connection_count"] = RpcInterFace{"delegate_set_network_min_connection_count", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_soft_max_imessage_length"] = RpcInterFace{"delegate_set_soft_max_imessage_length", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_transaction_canonical_signatures_required"] = RpcInterFace{"delegate_set_transaction_canonical_signatures_required", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_transaction_max_size"] = RpcInterFace{"delegate_set_transaction_max_size", RpcPassThrough, true}
	rpcInterFaceTable["delegate_set_transaction_min_fee"] = RpcInterFace{"delegate_set_transaction_min_fee", RpcPassThrough, true}
	rpcInterFaceTable["delete_event_handler"] = RpcInterFace{"delete_event_handler", RpcPassThrough, true}
	rpcInterFaceTable["destroy_contract"] = RpcInterFace{"destroy_contract", RpcPassThrough, true}
	rpcInterFaceTable["destroy_contract_testing"] = RpcInterFace{"destroy_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["disable_script"] = RpcInterFace{"disable_script", RpcPassThrough, true}
	rpcInterFaceTable["disk_usage"] = RpcInterFace{"disk_usage", RpcPassThrough, true}
	rpcInterFaceTable["download_validation"] = RpcInterFace{"download_validation", RpcPassThrough, true}
	rpcInterFaceTable["enable_script"] = RpcInterFace{"enable_script", RpcPassThrough, true}
	rpcInterFaceTable["execute_command_line"] = RpcInterFace{"execute_command_line", RpcPassThrough, true}
	rpcInterFaceTable["execute_script"] = RpcInterFace{"execute_script", RpcPassThrough, true}
	rpcInterFaceTable["export_script_db"] = RpcInterFace{"export_script_db", RpcPassThrough, true}
	rpcInterFaceTable["generate_download_validation"] = RpcInterFace{"generate_download_validation", RpcPassThrough, true}
	rpcInterFaceTable["get_contract_balance"] = RpcInterFace{"get_contract_balance", RpcPassThrough, true}
	rpcInterFaceTable["get_contract_info"] = RpcInterFace{"get_contract_info", RpcPassThrough, true}
	rpcInterFaceTable["get_contract_info_from_gpc_file"] = RpcInterFace{"get_contract_info_from_gpc_file", RpcPassThrough, true}
	rpcInterFaceTable["get_contract_registered_in_transaction"] = RpcInterFace{"get_contract_registered_in_transaction", RpcPassThrough, true}
	rpcInterFaceTable["get_events_bound"] = RpcInterFace{"get_events_bound", RpcPassThrough, true}
	rpcInterFaceTable["get_file_access"] = RpcInterFace{"get_file_access", RpcPassThrough, true}
	rpcInterFaceTable["get_info"] = RpcInterFace{"get_info", RpcPassThrough, true}
	rpcInterFaceTable["get_request_trx_id"] = RpcInterFace{"get_request_trx_id", RpcPassThrough, true}
	rpcInterFaceTable["get_result_trx_id"] = RpcInterFace{"get_result_trx_id", RpcPassThrough, true}
	rpcInterFaceTable["get_script_info"] = RpcInterFace{"get_script_info", RpcPassThrough, true}
	rpcInterFaceTable["get_transaction_id_contract_registered"] = RpcInterFace{"get_transaction_id_contract_registered", RpcPassThrough, true}
	rpcInterFaceTable["help"] = RpcInterFace{"help", RpcPassThrough, true}
	rpcInterFaceTable["http_start_server"] = RpcInterFace{"http_start_server", RpcPassThrough, true}
	rpcInterFaceTable["import_script_db"] = RpcInterFace{"import_script_db", RpcPassThrough, true}
	rpcInterFaceTable["list_event_handler"] = RpcInterFace{"list_event_handler", RpcPassThrough, true}
	rpcInterFaceTable["list_scripts"] = RpcInterFace{"list_scripts", RpcPassThrough, true}
	rpcInterFaceTable["load_contract_to_file"] = RpcInterFace{"load_contract_to_file", RpcPassThrough, true}
	rpcInterFaceTable["meta_help"] = RpcInterFace{"meta_help", RpcPassThrough, true}
	rpcInterFaceTable["network_add_node"] = RpcInterFace{"network_add_node", RpcPassThrough, true}
	rpcInterFaceTable["network_broadcast_transaction"] = RpcInterFace{"network_broadcast_transaction", RpcPassThrough, true}
	rpcInterFaceTable["network_get_advanced_node_parameters"] = RpcInterFace{"network_get_advanced_node_parameters", RpcPassThrough, true}
	rpcInterFaceTable["network_get_block_propagation_data"] = RpcInterFace{"network_get_block_propagation_data", RpcPassThrough, true}
	rpcInterFaceTable["network_get_blocked_ips"] = RpcInterFace{"network_get_blocked_ips", RpcPassThrough, true}
	rpcInterFaceTable["network_get_connection_count"] = RpcInterFace{"network_get_connection_count", RpcPassThrough, true}
	rpcInterFaceTable["network_get_info"] = RpcInterFace{"network_get_info", RpcPassThrough, true}
	rpcInterFaceTable["network_get_peer_info"] = RpcInterFace{"network_get_peer_info", RpcPassThrough, true}
	rpcInterFaceTable["network_get_transaction_propagation_data"] = RpcInterFace{"network_get_transaction_propagation_data", RpcPassThrough, true}
	rpcInterFaceTable["network_get_upnp_info"] = RpcInterFace{"network_get_upnp_info", RpcPassThrough, true}
	rpcInterFaceTable["network_list_potential_peers"] = RpcInterFace{"network_list_potential_peers", RpcPassThrough, true}
	rpcInterFaceTable["network_set_advanced_node_parameters"] = RpcInterFace{"network_set_advanced_node_parameters", RpcPassThrough, true}
	rpcInterFaceTable["ntp_update_time"] = RpcInterFace{"ntp_update_time", RpcPassThrough, true}
	rpcInterFaceTable["register_contract"] = RpcInterFace{"register_contract", RpcPassThrough, true}
	rpcInterFaceTable["register_contract_testing"] = RpcInterFace{"register_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["remove_script"] = RpcInterFace{"remove_script", RpcPassThrough, true}
	rpcInterFaceTable["rpc_set_password"] = RpcInterFace{"rpc_set_password", RpcPassThrough, true}
	rpcInterFaceTable["rpc_set_username"] = RpcInterFace{"rpc_set_username", RpcPassThrough, true}
	rpcInterFaceTable["rpc_start_server"] = RpcInterFace{"rpc_start_server", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_account_balance"] = RpcInterFace{"sandbox_account_balance", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_call_contract"] = RpcInterFace{"sandbox_call_contract", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_call_contract_testing"] = RpcInterFace{"sandbox_call_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_close"] = RpcInterFace{"sandbox_close", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_compile_contract"] = RpcInterFace{"sandbox_compile_contract", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_destroy_contract"] = RpcInterFace{"sandbox_destroy_contract", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_destroy_contract_testing"] = RpcInterFace{"sandbox_destroy_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_get_contract_balance"] = RpcInterFace{"sandbox_get_contract_balance", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_get_contract_info"] = RpcInterFace{"sandbox_get_contract_info", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_list_my_addresses"] = RpcInterFace{"sandbox_list_my_addresses", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_load_contract_to_file"] = RpcInterFace{"sandbox_load_contract_to_file", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_open"] = RpcInterFace{"sandbox_open", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_register_contract"] = RpcInterFace{"sandbox_register_contract", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_register_contract_testing"] = RpcInterFace{"sandbox_register_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_transfer_to_contract"] = RpcInterFace{"sandbox_transfer_to_contract", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_transfer_to_contract_testing"] = RpcInterFace{"sandbox_transfer_to_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_upgrade_contract"] = RpcInterFace{"sandbox_upgrade_contract", RpcPassThrough, true}
	rpcInterFaceTable["sandbox_upgrade_contract_testing"] = RpcInterFace{"sandbox_upgrade_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["stop"] = RpcInterFace{"stop", RpcPassThrough, true}
	rpcInterFaceTable["store_file_piece"] = RpcInterFace{"store_file_piece", RpcPassThrough, true}
	rpcInterFaceTable["store_file_to_network"] = RpcInterFace{"store_file_to_network", RpcPassThrough, true}
	rpcInterFaceTable["store_reject"] = RpcInterFace{"store_reject", RpcPassThrough, true}
	rpcInterFaceTable["upgrade_contract"] = RpcInterFace{"upgrade_contract", RpcPassThrough, true}
	rpcInterFaceTable["upgrade_contract_testing"] = RpcInterFace{"upgrade_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["validate_address"] = RpcInterFace{"validate_address", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_balance"] = RpcInterFace{"wallet_account_balance", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_balance_ids"] = RpcInterFace{"wallet_account_balance_ids", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_balance_rpc"] = RpcInterFace{"wallet_account_balance_rpc", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_create"] = RpcInterFace{"wallet_account_create", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_delete"] = RpcInterFace{"wallet_account_delete", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_historic_balance"] = RpcInterFace{"wallet_account_historic_balance", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_list_public_keys"] = RpcInterFace{"wallet_account_list_public_keys", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_register"] = RpcInterFace{"wallet_account_register", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_rename"] = RpcInterFace{"wallet_account_rename", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_retract"] = RpcInterFace{"wallet_account_retract", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_set_approval"] = RpcInterFace{"wallet_account_set_approval", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_transaction_history"] = RpcInterFace{"wallet_account_transaction_history", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_update_active_key"] = RpcInterFace{"wallet_account_update_active_key", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_update_private_data"] = RpcInterFace{"wallet_account_update_private_data", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_update_registration"] = RpcInterFace{"wallet_account_update_registration", RpcPassThrough, true}
	rpcInterFaceTable["wallet_account_vote_summary"] = RpcInterFace{"wallet_account_vote_summary", RpcPassThrough, true}
	rpcInterFaceTable["wallet_active_delegate_salary"] = RpcInterFace{"wallet_active_delegate_salary", RpcPassThrough, true}
	rpcInterFaceTable["wallet_address_create"] = RpcInterFace{"wallet_address_create", RpcPassThrough, true}
	rpcInterFaceTable["wallet_allow_store_request"] = RpcInterFace{"wallet_allow_store_request", RpcPassThrough, true}
	rpcInterFaceTable["wallet_asset_create"] = RpcInterFace{"wallet_asset_create", RpcPassThrough, true}
	rpcInterFaceTable["wallet_asset_issue"] = RpcInterFace{"wallet_asset_issue", RpcPassThrough, true}
	rpcInterFaceTable["wallet_asset_issue_to_addresses"] = RpcInterFace{"wallet_asset_issue_to_addresses", RpcPassThrough, true}
	rpcInterFaceTable["wallet_backup_create"] = RpcInterFace{"wallet_backup_create", RpcPassThrough, true}
	rpcInterFaceTable["wallet_backup_restore"] = RpcInterFace{"wallet_backup_restore", RpcPassThrough, true}
	rpcInterFaceTable["wallet_balance_set_vote_info"] = RpcInterFace{"wallet_balance_set_vote_info", RpcPassThrough, true}
	rpcInterFaceTable["wallet_cancel_scan"] = RpcInterFace{"wallet_cancel_scan", RpcPassThrough, true}
	rpcInterFaceTable["wallet_change_passphrase"] = RpcInterFace{"wallet_change_passphrase", RpcPassThrough, true}
	rpcInterFaceTable["wallet_check_address"] = RpcInterFace{"wallet_check_address", RpcPassThrough, true}
	rpcInterFaceTable["wallet_check_passphrase"] = RpcInterFace{"wallet_check_passphrase", RpcPassThrough, true}
	rpcInterFaceTable["wallet_check_vote_status"] = RpcInterFace{"wallet_check_vote_status", RpcPassThrough, true}
	rpcInterFaceTable["wallet_close"] = RpcInterFace{"wallet_close", RpcPassThrough, true}
	rpcInterFaceTable["wallet_collect_genesis_balances"] = RpcInterFace{"wallet_collect_genesis_balances", RpcPassThrough, true}
	rpcInterFaceTable["wallet_create"] = RpcInterFace{"wallet_create", RpcPassThrough, true}
	rpcInterFaceTable["wallet_delegate_pay_balance_query"] = RpcInterFace{"wallet_delegate_pay_balance_query", RpcPassThrough, true}
	rpcInterFaceTable["wallet_delegate_set_block_production"] = RpcInterFace{"wallet_delegate_set_block_production", RpcPassThrough, true}
	rpcInterFaceTable["wallet_delegate_withdraw_pay"] = RpcInterFace{"wallet_delegate_withdraw_pay", RpcPassThrough, true}
	rpcInterFaceTable["wallet_dump_account_private_key"] = RpcInterFace{"wallet_dump_account_private_key", RpcPassThrough, true}
	rpcInterFaceTable["wallet_dump_private_key"] = RpcInterFace{"wallet_dump_private_key", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_account"] = RpcInterFace{"wallet_get_account", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_account_owner_publickey"] = RpcInterFace{"wallet_get_account_owner_publickey", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_account_public_address"] = RpcInterFace{"wallet_get_account_public_address", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_all_approved_accounts"] = RpcInterFace{"wallet_get_all_approved_accounts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_contracts"] = RpcInterFace{"wallet_get_contracts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_delegate_statue"] = RpcInterFace{"wallet_get_delegate_statue", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_info"] = RpcInterFace{"wallet_get_info", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_my_access"] = RpcInterFace{"wallet_get_my_access", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_my_store_confirmed"] = RpcInterFace{"wallet_get_my_store_confirmed", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_my_store_rejected"] = RpcInterFace{"wallet_get_my_store_rejected", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_my_store_request"] = RpcInterFace{"wallet_get_my_store_request", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_my_upload_requests"] = RpcInterFace{"wallet_get_my_upload_requests", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_name"] = RpcInterFace{"wallet_get_name", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_setting"] = RpcInterFace{"wallet_get_setting", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_transaction"] = RpcInterFace{"wallet_get_transaction", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_transaction_fee"] = RpcInterFace{"wallet_get_transaction_fee", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_transaction_imessage_fee_coe"] = RpcInterFace{"wallet_get_transaction_imessage_fee_coe", RpcPassThrough, true}
	rpcInterFaceTable["wallet_get_transaction_imessage_soft_max_length"] = RpcInterFace{"wallet_get_transaction_imessage_soft_max_length", RpcPassThrough, true}
	rpcInterFaceTable["wallet_import_private_key"] = RpcInterFace{"wallet_import_private_key", RpcPassThrough, true}
	rpcInterFaceTable["wallet_list"] = RpcInterFace{"wallet_list", RpcPassThrough, true}
	rpcInterFaceTable["wallet_list_accounts"] = RpcInterFace{"wallet_list_accounts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_list_my_accounts"] = RpcInterFace{"wallet_list_my_accounts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_list_my_addresses"] = RpcInterFace{"wallet_list_my_addresses", RpcPassThrough, true}
	rpcInterFaceTable["wallet_list_store_request_for_my_file"] = RpcInterFace{"wallet_list_store_request_for_my_file", RpcPassThrough, true}
	rpcInterFaceTable["wallet_list_unregistered_accounts"] = RpcInterFace{"wallet_list_unregistered_accounts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_lock"] = RpcInterFace{"wallet_lock", RpcPassThrough, true}
	rpcInterFaceTable["wallet_login_finish"] = RpcInterFace{"wallet_login_finish", RpcPassThrough, true}
	rpcInterFaceTable["wallet_login_start"] = RpcInterFace{"wallet_login_start", RpcPassThrough, true}
	rpcInterFaceTable["wallet_open"] = RpcInterFace{"wallet_open", RpcPassThrough, true}
	rpcInterFaceTable["wallet_publish_slate"] = RpcInterFace{"wallet_publish_slate", RpcPassThrough, true}
	rpcInterFaceTable["wallet_publish_version"] = RpcInterFace{"wallet_publish_version", RpcPassThrough, true}
	rpcInterFaceTable["wallet_rebroadcast_transaction"] = RpcInterFace{"wallet_rebroadcast_transaction", RpcPassThrough, true}
	rpcInterFaceTable["wallet_recover_accounts"] = RpcInterFace{"wallet_recover_accounts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_regenerate_keys"] = RpcInterFace{"wallet_regenerate_keys", RpcPassThrough, true}
	rpcInterFaceTable["wallet_remove_contact_account"] = RpcInterFace{"wallet_remove_contact_account", RpcPassThrough, true}
	rpcInterFaceTable["wallet_remove_transaction"] = RpcInterFace{"wallet_remove_transaction", RpcPassThrough, true}
	rpcInterFaceTable["wallet_repair_entrys"] = RpcInterFace{"wallet_repair_entrys", RpcPassThrough, true}
	rpcInterFaceTable["wallet_rescan_blockchain"] = RpcInterFace{"wallet_rescan_blockchain", RpcPassThrough, true}
	rpcInterFaceTable["wallet_scan_contracts"] = RpcInterFace{"wallet_scan_contracts", RpcPassThrough, true}
	rpcInterFaceTable["wallet_scan_transaction"] = RpcInterFace{"wallet_scan_transaction", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_automatic_backups"] = RpcInterFace{"wallet_set_automatic_backups", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_node_id"] = RpcInterFace{"wallet_set_node_id", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_setting"] = RpcInterFace{"wallet_set_setting", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_transaction_expiration_time"] = RpcInterFace{"wallet_set_transaction_expiration_time", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_transaction_fee"] = RpcInterFace{"wallet_set_transaction_fee", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_transaction_imessage_fee_coe"] = RpcInterFace{"wallet_set_transaction_imessage_fee_coe", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_transaction_imessage_soft_max_length"] = RpcInterFace{"wallet_set_transaction_imessage_soft_max_length", RpcPassThrough, true}
	rpcInterFaceTable["wallet_set_transaction_scanning"] = RpcInterFace{"wallet_set_transaction_scanning", RpcPassThrough, true}
	rpcInterFaceTable["wallet_sign_hash"] = RpcInterFace{"wallet_sign_hash", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transaction_history_splite"] = RpcInterFace{"wallet_transaction_history_splite", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transfer_to_address"] = RpcInterFace{"wallet_transfer_to_address", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transfer_to_address_rpc"] = RpcInterFace{"wallet_transfer_to_address_rpc", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transfer_to_contract"] = RpcInterFace{"wallet_transfer_to_contract", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transfer_to_contract_testing"] = RpcInterFace{"wallet_transfer_to_contract_testing", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transfer_to_public_account"] = RpcInterFace{"wallet_transfer_to_public_account", RpcPassThrough, true}
	rpcInterFaceTable["wallet_transfer_to_public_account_rpc"] = RpcInterFace{"wallet_transfer_to_public_account_rpc", RpcPassThrough, true}
	rpcInterFaceTable["wallet_unlock"] = RpcInterFace{"wallet_unlock", RpcPassThrough, true}
	rpcInterFaceTable["wallet_verify_titan_deposit"] = RpcInterFace{"wallet_verify_titan_deposit", RpcPassThrough, true}
	rpcInterFaceTable["wallet_withdraw_from_address"] = RpcInterFace{"wallet_withdraw_from_address", RpcPassThrough, true}

	// abbreviation
	rpcInterFaceTable["login"] = RpcInterFace{"login", RpcPassThrough, true}
	rpcInterFaceTable["unlock"] = RpcInterFace{"unlock", RpcPassThrough, true}
	rpcInterFaceTable["balance"] = RpcInterFace{"balance", RpcPassThrough, true}
	rpcInterFaceTable["info"] = RpcInterFace{"info", RpcPassThrough, true}
	rpcInterFaceTable["open"] = RpcInterFace{"open", RpcPassThrough, true}
}


func RpcPassThrough(request *json.Json, tvConn net.Conn)(response string) {
	
 	if request != nil {
		//id := request.Get("id")
		method, _ := request.Get("method").String()
		params, _ := request.Get("params").Array()
	    fmt.Println("RpcPassThrough called :", method , "[")
	    for i := 0; i < len(params); i++ {
			fmt.Println( params[i])
	    }
		fmt.Println( "]")

	    requestStr := json.GenerateJsonString(method, params)
	    response = network.CallRpc(requestStr, tvConn)
	    fmt.Println(response)
	}
	return
}

func UploadFileToIPFS(request *json.Json, tvConn net.Conn)(response string) {
	fmt.Println("UploadFileToIPFS called ")
	params, _ := request.Get("params").Array()
	response, _ = logic.UploadFileToIPFS(params[0].(string), tvConn)
	return
}

func DeclareUploadFile(request *json.Json, tvConn net.Conn)(response string) {
	fmt.Println("DeclareUploadFile called ")
	method, _ := request.Get("method").String()
	params, _ := request.Get("params").Array()
	if method == "DeclareUploadFile" {
		method = "store_file_to_network"
	}
	requestStr := json.GenerateJsonString(method, params)
	response, _ = logic.DeclareUploadFile(requestStr, tvConn)
	return
}

func ListUploadDeclaration(request *json.Json, tvConn net.Conn)(response string) {
	fmt.Println("ListUploadDeclaration called ")
	method, _ := request.Get("method").String()
	params, _ := request.Get("params").Array()
	if method == "ListUploadDeclaration" {
		method = "blockchain_get__upload_requests"
	}
	requestStr := json.GenerateJsonString(method, params)
	response, _ = logic.DeclareUploadFile(requestStr, tvConn)
	return
}

func PinAddFileToLocal(request *json.Json, tvConn net.Conn)(response string) {
	fmt.Println("PinAddFileToLocal called ")
	params, _ := request.Get("params").Array()
	response, _ = logic.PinAddFileToLocal(params[0].(string), tvConn)
	return
}

func DeclarePieceSaved(request *json.Json, tvConn net.Conn)(response string) {
	fmt.Println("DeclarePieceSaved called ")
	method, _ := request.Get("method").String()
	params, _ := request.Get("params").Array()
	if method == "DeclarePieceSaved" {
		method = "declare_piece_saved"
	}
	requestStr := json.GenerateJsonString(method, params)
	response, _ = logic.DeclarePieceSaved(requestStr, tvConn)
	return
}

func ListFileStoreRequest(request *json.Json, tvConn net.Conn) (response string) {
	fmt.Println("ListFileStoreRequest called ")
	if request != nil {
		//id := request.Get("id")
		method, _ := request.Get("method").String()
		params, _ := request.Get("params").Array()

		if (method == "ListFileStoreRequest") {
			requestStr := json.GenerateJsonString("blockchain_list_file_save_declare", params)
			response = logic.ListStoreRequest(requestStr, tvConn)
			//fmt.Println(response)
		}
	}
	return response
}

func ConfirmFileSaved(request *json.Json, tvConn net.Conn) (response string) {
	fmt.Println("ConfirmFileSaved called ")
	if request != nil {
		//id := request.Get("id")
		method, _ := request.Get("method").String()
		params, _ := request.Get("params").Array()

		if (method == "ConfirmFileSaved") {
			requestStr := json.GenerateJsonString("confirm_piece_saved", params)
			response = logic.ConfirmPieceSaved(requestStr, tvConn)
			fmt.Println(response)
		}
	}
	return response
}

func ListConfirmFileSaved(request *json.Json, tvConn net.Conn) (response string) {
	fmt.Println("ListConfirmFileSaved called ")
	if request != nil {
		//id := request.Get("id")
		method, _ := request.Get("method").String()
		params, _ := request.Get("params").Array()

		if (method == "ListConfirmFileSaved") {
			requestStr := json.GenerateJsonString("blockchain_list_file_saved", params)
			response = logic.ListConfirmSavedFile(requestStr, tvConn)
			fmt.Println(response)
		}
	}
	return response
}

