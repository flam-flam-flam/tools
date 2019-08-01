package com.ukelink.um.bean;

import lombok.Data;

import javax.annotation.sql.DataSourceDefinition;

@Data
public class VmpHeader {
    /** 标志位 */
    private Byte magic;

    /** 包头长度(高6位，低两位填充1） */
    private Byte headerLen;

    /** 压缩标志(1非压缩，2压缩) */
    private Byte encrypType;

    /** 加密算法(1 AES，2 RSA) */
    private Byte zipFlag;

    /** token字节长度 */
    private Byte tokenLen;

    /** 发送端使用的协议版本 */
    private Short version;

    /** 用户id */
    private Long userId;

    /** token */
    private String token;

    /** 消息路由到业务的id */
    private Short cgiId;

    /** 原始包体的长度 */
    private Integer lenOrgBody;

    /** 压缩后的包体长度 */
    private Integer lenCompressed;
}
