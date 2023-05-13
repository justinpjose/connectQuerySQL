# connectQuerySQL
This script allows us to connect to the remote databases locally and then utilise the methods already created to retrieve data or create new methods to carry out new SQL queries

## Instructions
1. Provide the database configurations and if applicable, the ssh configuration

2. Using the configurations, you can then call the method `ConnectToDB` to connect to the database
```
    import (
	    "connectQuerySQL/backbonev2"
	    "connectQuerySQL/config"
	    "log"
    )
    
    dbCfg := config.DBConfig{
		Name:     "",
		Host:     "",
		Password: "",
		Port:     000,
		Username: "",
	}

	sshCfg := config.SSHConfig{
		Host:               "",
		Port:               00,
		User:               "",
		PrivateKeyPath:     "",
		PrivateKeyPassword: "",
	}

	backbonev2, closeConn, err := backbonev2.ConnectToDB(dbCfg, sshCfg)
	if err != nil {
		log.Fatalf("failed to connect to backbone v2 db - err: %v", err)
	}
    defer closeConn()
```

3. This will return a `Service` struct which will contain the following information
```
type Service struct {
	Db     *sql.DB
	Tables dbTables
}
```

4. You can use the above to begin a transaction on the DB
```
tx, err := backbonev2.Db.BeginTx(ctx, nil)
if err != nil {
	log.Fatal(postgresql.TxBeginError(err))
}
defer postgresql.Commit(tx)
```

5. Then using an existing method from the tables
```
style, err := backbonev2.Tables.Injector.FindByProductNumber(ctx, productNo, tx)
	if err != nil {
		log.Fatalf("failed to get style by product number - err: %v", err)
	}
```

6. OR execute your own SQL query to query to the database. Examples can found in https://github.com/River-Island/product-backbone-v2/blob/master/pkg/shared/store/injector/injector.go

7. Run with the command `go run main.go run.go`

## Limitations
- As of now, you can only SSH using a privateKey and not a password. This is because we provide a private key to connect to the injector and backbone database instead of a password

## Notes
- `rootLogger := testutil.NopLogger()`