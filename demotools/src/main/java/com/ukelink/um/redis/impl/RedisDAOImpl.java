package com.ukelink.um.redis.impl;

import com.ukelink.um.redis.RedisDAO;
import org.apache.commons.lang.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import redis.clients.jedis.JedisCluster;

import java.util.List;
import java.util.Map;
import java.util.Set;

/**
 * @ClassName RedisDAOImpl
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-7-30 16:07
 * @Version 1.0
 */
@Repository
public class RedisDAOImpl implements RedisDAO {
    @Autowired
    private JedisCluster jc;


    @Override
    public String set(String logId, String key, String value) {
        String set;
        set = jc.set(key, value);
        return set;
    }

    @Override
    public String get(String logId, String key) {
        String get;
        get = jc.get(key);
        return get;
    }

    @Override
    public Long del(String logId, String key) {
        Long del;
        del = jc.del(key);
        return del;
    }

    @Override
    public Long hset(String logId, String key, String field, String value) {
        Long hset;
        hset = jc.hset(key, field, value);
        return hset;
    }

    @Override
    public String hmset(String logId, String key, Map<String, String> hash) {
        String hmset;
        hmset = jc.hmset(key, hash);
        return hmset;
    }

    @Override
    public String hget(String logId, String key, String field) {
        String hget;
        hget = jc.hget(key, field);
        return hget;
    }

    @Override
    public Map<String, String> hgetAll(String logId, String key) {
        Map<String, String> hgetAll;
        hgetAll = jc.hgetAll(key);
        return hgetAll;
    }

    @Override
    public Long hdel(String logId, String key, String... field) {
        Long hdel;
        hdel = jc.hdel(key, field);
        return hdel;
    }

    @Override
    public Boolean hexists(String logId, String key, String field) {
        Boolean hexists;
        hexists = jc.hexists(key, field);
        return hexists;
    }

    @Override
    public Set<String> hkeys(String logId, String key) {
        Set<String> hkeys;
        hkeys = jc.hkeys(key);
        return hkeys;
    }

    @Override
    public Long incr(String logId, String key) {
        Long incr;
        incr = jc.incr(key);
        return incr;
    }

    @Override
    public Long zadd(String logId, String key, Map<String, Double> scoreMembers) {
        Long zadd;
        zadd = jc.zadd(key, scoreMembers);
        return zadd;
    }

    @Override
    public Long zadd(String logId, String key, Double score, String member) {
        Long zadd;
        zadd = jc.zadd(key, score, member);
        return zadd;
    }

    @Override
    public Set<String> zrangeByScore(String logId, String key, String min, String max) {
        Set<String> zrangeByScore;
        zrangeByScore = jc.zrangeByScore(key, min, max);
        return zrangeByScore;
    }

    @Override
    public Set<String> zrange(String logId, String key, Long start, Long stop) {
        Set<String> zrange;
        zrange = jc.zrange(key, start, stop);
        return zrange;
    }

    @Override
    public Double zscore(String logId, String key, String member) {
        Double zscore;
        zscore = jc.zscore(key, member);
        return zscore;
    }
    @Override
//    public Object luaEval(String logId, String script,String sampleKey)
    public Object luaEval(String logId,String script, String sampleKey, List<String> key, List<String> argv)
    {
        Object result;
//        Result = jc.scriptLoad(script,sampleKey);
        result = jc.eval(script,key,argv);
        return result;
    }
    @Override
    public Object luaEvalSha1(String logId, String script, String sampleKey,List<String> key, List<String> argv)
    {
        Object result= null;
        String sha1 = jc.get(sampleKey);
        if(StringUtils.isEmpty(sha1))
        {
            System.out.println("sha1 is empty====================");
            sha1 = jc.scriptLoad(script,sampleKey);
            jc.set(sampleKey,sha1);
        }
        result = jc.evalsha(sha1, key, argv);
        return result;
    }
}
