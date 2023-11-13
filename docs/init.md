ent 使用

0、初始化 ent
go run entgo.io/ent/cmd/ent new

1、从数据库中创建 schema
go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:root@tcp(127.0.0.1:3306)/pay-system"

2、生成代码
go generate ./ent

```
func main() {
    drv, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pay-system?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))
    defer client.Close()
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    QueryUser(ctx, client.Debug())
}
```

```
func QueryUser(ctx context.Context, client *ent.Client) (*ent.ProductApp, error) {
	u, err := client.ProductApp.
		Query().
		Where(productapp.AppID("123")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
```
