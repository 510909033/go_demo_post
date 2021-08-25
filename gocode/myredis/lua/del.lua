-- lua删除锁：
-- KEYS和ARGV分别是以集合方式传入的参数，对应上文的Test和uuid。
-- 如果对应的value等于传入的uuid。
if redis.call('get', KEYS[1]) == ARGV[1]
then
    -- 执行删除操作
    return redis.call('del', KEYS[1])
else
    -- 不成功，返回0
    return 0
end