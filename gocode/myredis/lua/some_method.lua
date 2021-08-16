local k1=KEYS[1]
local k2=KEYS[2]
local k3=KEYS[3]
local k4=KEYS[4]
local a1=ARGV[1]
local a2=ARGV[2]
local a3=ARGV[3]
local a4=ARGV[4]
local result = {}
--local result=greet.." "..name
--local obj = {key=KEYS[1], argv1=ARGV[1], value=name }
--obj = {name = 'felord.cn', age = 18}
--print(obj)
--return obj
result = {1,2,{3,'Hello World!'}}
result = {k1,k2,k3,k4,a1,a2,a3,a4}
return result