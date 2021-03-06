package credtrack

import "theodo.red/creditcompanion/packages/money"

type CreditTransaction struct {
	Id                   string                     `dynamodbav:id`
	CreditSourceId       string                     `dynamodbav:creditSourceId`
	CreatedAt            string                     `dynamodbav:createdAt`
	LinkedClients        map[string]float32         `dynamodbav:linkedClients`
	Total                money.MonetaryAmount       `dynamodbav:total`
	TransactionDirection money.TransactionDirection `dynamodbav:transactionDirection`
	Description          string                     `dynamodbav:description`
	TransactedAt         string                     `dynamodbav:transactedAt`
}
