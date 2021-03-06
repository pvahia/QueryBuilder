<h1 align="center">
  <a href="https://github.com/Snehal1112/QueryBuilder"><img src="./querybuilder.png" alt="QueryBuilder" height="250px"></a>
  <br>
  QueryBuilder
  <br>
</h1>

Query Builder is a flexible tool developed using Go, that is used to help in creating all kinds of SQL queries. 
Query Builder lets you create lengthy queries and complexity without the need to know the underlying database model.

First of all, we need to follow the below steps to set the env variable and create SQL Builder instance,

#### SET ENV VARIABLE : 

```
  os.Setenv("DB_USER", "root")
  os.Setenv("DB_PASSWORD", "Snehal@1977")
  os.Setenv("DB_DATABASE", "querybuilder")
  os.Setenv("DRIVER", "mysql")
```

#### CREATE SQL BUILDER INSTANCE :

```
  builders := builder.NewSQLBuilder("mysql")
```

After doing this, we can now use ```builders``` for creating any query.



# Supported sql query
  - [ ] DDL (Data Definition Language)
    - [X] Create (table, database)
    
      - Example :
       ```
          ddlQuery := builders.NewDDL()
          createDatabase := ddlQuery.Create().Database("TestDatabase")
          result, err := createDatabase.Execute()
          if err != nil {
            log.Fatal(err)
          }
          _, err = result.LastInsertId()
          if err != nil {
            log.Fatal(err)
          }
          log.Println("TestDatabase is created")
        ```
          
    - [X] Drop (table, database)
    
      - Example :
       ```
          drop := builders.NewDDL().Drop()
          table := drop.Table([]string{"products"})
          result, err := table.Execute()
          if err != nil {
            log.Fatal(err)
          }
          output, er := result.RowsAffected()
          if er != nil {
            log.Fatal(err)
          }
          log.Println("products table drop successfully")
        }
        ```
        
    - [X] Alter (table, database)
    - [ ] Truncate
  - [ ] DML (Data Manipulation Language)
    - [ ] Insert query 
    - [ ] Update query
    - [ ] Delete query
  - [ ] DCL (Data Control Language)
    - [ ] Grant
    - [ ] Revoke Grant
  - [ ] TCL (Transaction Control Language)
    - [ ] Commit
    - [ ] Rollback
    - [ ] Save point
  - [ ] DQL (Data Query language)
    - [ ] Select
    
 For more examples, have a look at the [examples](https://github.com/Snehal1112/QueryBuilder/tree/master/example) folder.