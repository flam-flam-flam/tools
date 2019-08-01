package com.ukelink.um.handle.netty;

import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.socket.SocketChannel;

public class NettyClientInit extends ChannelInitializer<SocketChannel> {
    @Override
    public void initChannel(SocketChannel ch){
        ChannelPipeline p = ch.pipeline();
        p.addLast("encode", new EncodeHandler());
        p.addLast("decode", new DecodeHandler());
        p.addLast("business",new BusinessHandler());
    }
}
