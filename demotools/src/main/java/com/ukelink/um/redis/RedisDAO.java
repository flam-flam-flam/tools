package com.ukelink.um.redis;

import java.util.List;
import java.util.Map;
import java.util.Set;

/**
 * @InterfaceName RedisDAO
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-7-30 16:04
 * @Version 1.0
 */
public interface RedisDAO {
    /**
     * redis中的set方法
     *
     * @param logId 日志id
     * @param key key
     * @param value value
     * @return 操作结果
     */
    String set(String logId, String key, String value);

    /**
     * redis中的get方法
     *
     * @param logId 日志id
     * @param key key
     * @return 操作结果
     */
    String get(String logId, String key);

    /**
     * redis中的del方法
     *
     * @param logId 日志id
     * @param key key
     * @return 操作结果
     */
    Long del(String logId, String key);

    /**
     * redis中的hset方法
     *
     * @param logId 日志id
     * @param key key
     * @param field field
     * @param value value
     * @return 操作结果
     */
    Long hset(String logId, String key, String field, String value);

    /**
     * redis中的hmset方法
     *
     * @param logId 日志id
     * @param key key
     * @param hash hash
     * @return 操作结果
     */
    String hmset(String logId, String key, Map<String, String> hash);

    /**
     * redis中的hget方法
     *
     * @param logId 日志id
     * @param key key
     * @param field field
     * @return 操作结果
     */
    String hget(String logId, String key, String field);

    /**
     * redis中的hgetAll方法
     *
     * @param logId 日志id
     * @param key key
     * @return 操作结果
     */
    Map<String, String> hgetAll(String logId, String key);

    /**
     * redis中的hdel方法
     *
     * @param logId 日志id
     * @param key key
     * @param field field
     * @return 操作结果
     */
    Long hdel(String logId, String key, String... field);

    /**
     * redis中的hexists方法
     *
     * @param logId 日志id
     * @param key key
     * @param field field
     * @return 操作结果
     */
    Boolean hexists(String logId, String key, String field);

    /**
     * redis中的hkeys方法
     *
     * @param logId 日志id
     * @param key key
     * @return 操作结果
     */
    Set<String> hkeys(String logId, String key);

    /**
     * redis中的incr方法
     *
     * @param logId 日志id
     * @param key key
     * @return 操作结果
     */
    Long incr(String logId, String key);

    /**
     * redis中的zadd方法
     *
     * @param logId 日志id
     * @param key key
     * @param scoreMembers 批量存储
     * @return 操作结果
     */
    Long zadd(String logId, String key, Map<String, Double> scoreMembers);

    /**
     * redis中的zadd方法
     *
     * @param logId 日志id
     * @param key key
     * @param score score
     * @param member member
     * @return 操作结果
     */
    Long zadd(String logId, String key, Double score, String member);

    Set<String> zrangeByScore(String logId, String key, String min, String max);

    Set<String> zrange(String logId, String key, Long start, Long stop);

    Double zscore(String logId, String key, String member);

    Object luaEval(String logId, String script,String sampleKey, List<String> key, List<String> argv);

    Object luaEvalSha1(String logId, String script, String sampleKey,List<String> key, List<String> argv);
    //Object luaEval(String logId, String script,String sampleKey);
}
