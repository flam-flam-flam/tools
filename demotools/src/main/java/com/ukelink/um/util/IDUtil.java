package com.ukelink.um.util;

import cn.hutool.core.util.IdUtil;

import java.util.UUID;
import java.util.concurrent.ThreadLocalRandom;
import java.util.concurrent.atomic.AtomicInteger;

public class IDUtil {
    //private  static  AtomicInteger a = new AtomicInteger();
    public static String getId() {
//        return String.valueOf(System.currentTimeMillis()) + ThreadLocalRandom.current().nextInt(100000000, 1000000000);
//        return UUID.randomUUID().toString();
        return IdUtil.objectId();
//        return String.valueOf(a.getAndIncrement());

    }
}
