package com.ukelink.um.redis;

/**
 * @ClassName LuaScript
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-7-31 14:34
 * @Version 1.0
 */
public class LuaScript {
   public static final String TESTLUA = "redis.call('set', KEYS[1], ARGV[1])" +
            "redis.call('incr', KEYS[1]) " +
            "local value = redis.call('get',KEYS[1])" +
            "return value" ;
}
