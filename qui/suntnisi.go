	// The transaction block can be used to read data from the transaction.
	// It must be closed when done.
	transactionBlock, err := db.NewTransaction(ctx)
	if err != nil {
		return err
	}
	defer transactionBlock.Close()  
