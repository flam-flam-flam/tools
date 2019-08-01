package com.ukelink.um.bean;

import lombok.Data;

@Data
public class LongLinkHeader {
    /** 消息包大小 */
    private Integer packLen;
    /** 长连接消息头大小*/
    private Short headerLen;
    /** 版本号*/
    private Short version;
    /** CgiType*/
    private Short cmdId;
    /** 消息包序列号 */
    private Integer seq;
    /** 消息唯一标识 */
    private String deviceId;
}
