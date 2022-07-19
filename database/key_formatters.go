package database

func formatUserKey(userId string) string {
	return "user:" + userId
}

func formatTransactionKey(txnId string) string {
	return "transaction:" + txnId
}
