package com.ukelink.um.util;

import sun.security.provider.MD5;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 * @ClassName Md5Util
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-7-29 11:14
 * @Version 1.0
 */
public class Md5Util  {
    public static String md5(String userName, String password) throws NoSuchAlgorithmException {

       // String password = "456:mytest:456";
        String userDigest = userName + ":" + "mytest:" + password;
        System.out.println("userDigest: " + userDigest);
        MessageDigest md = MessageDigest.getInstance("MD5");
        byte[] hashInBytes = md.digest(userDigest.getBytes(StandardCharsets.UTF_8));

        StringBuilder sb = new StringBuilder();
        for (byte b : hashInBytes) {
            sb.append(String.format("%02x", b));
        }
        System.out.println(sb.toString());
        return sb.toString();
    }
}
