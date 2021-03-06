module theodo.red/creditcompanion/packages/tokens

go 1.15

replace theodo.red/creditcompanion/packages/teatime => ../teatime

replace theodo.red/creditcompanion/packages/logging => ../logging

replace theodo.red/creditcompanion/packages/database => ../database

require (
	github.com/aws/aws-sdk-go v1.36.31
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f
	github.com/stretchr/testify v1.7.0
	theodo.red/creditcompanion/packages/database v1.0.0
	theodo.red/creditcompanion/packages/logging v1.0.0
	theodo.red/creditcompanion/packages/teatime v1.0.0
)
