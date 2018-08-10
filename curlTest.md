# 用户登录
curl -XPOST -H "Content-Type: application/json" http://10.96.9.227:8090/login -d'{"username":"skr","password":"123456"}'

# 用户创建
curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user -d'{"username":"skr2","password":"123456"}'

# 用户查询
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user/skr

# 用户删除
curl -XDELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user/5

# 用户修改
curl -XPUT -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user/4 -d'{"username":"likarl","password":"123456"}'

# 用户列表
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user

#-------------------------------————#

# HTTPS
# 用户登录
curl -XPOST -H "Content-Type: application/json" https://10.96.9.227:8081/login -d'{"username":"skr","password":"123456"}' --cacert conf/server.crt --cert conf/server.crt --key conf/server.key

# 用户创建
curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user -d'{"username":"skr2","password":"123456"}'

# 用户查询
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user/skr

# 用户删除
curl -XDELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user/5

# 用户修改
curl -XPUT -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" http://10.96.9.227:8090/v1/user/4 -d'{"username":"likarl","password":"123456"}'

# 用户列表
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM4NzA0ODgsImlkIjozLCJuYmYiOjE1MzM4NzA0ODgsInVzZXJuYW1lIjoic2tyIn0.ej_avc-Sr27N6eihZq5lB0_rAyClAHRx1pw_e_IXMuE" -H "Content-Type: application/json" https://10.96.9.227:8081/v1/user --cacert conf/server.crt --cert conf/server.crt --key conf/server.key

