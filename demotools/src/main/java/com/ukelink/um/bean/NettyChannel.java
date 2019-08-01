package com.ukelink.um.bean;

import io.netty.channel.ChannelHandlerContext;
import lombok.Data;

import java.util.HashMap;

@Data
public class NettyChannel {
    public ChannelHandlerContext channelHandlerContext;
}
