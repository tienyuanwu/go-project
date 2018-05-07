安裝
1. 安裝 go https://golang.org/doc/install
2. 執行 make
3. 執行 ./source

提供API
1. http://localhost:8080/record (get) 
取得 record id 列表

2. http://localhost:8080/record (post)
新增紀錄（資料格式請參考test/test.json)
會回傳新增的 record id

3. http://localhost:8080//chart3d?id=<id>&table=<table name>
取得3d圖資料的API，目前table只有1，id則參考1, 2 API回傳的結果。
  
目前系統尚未有 database，所以資料都只儲存在記憶體中，重開資料會消失。
