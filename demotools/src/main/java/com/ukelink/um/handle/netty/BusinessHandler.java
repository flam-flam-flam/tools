package com.ukelink.um.handle.netty;

import com.ukelink.um.bean.LongLinkHeader;
import com.ukelink.um.bean.LongLinkMessage;
import io.netty.channel.*;
import io.netty.util.concurrent.GenericFutureListener;
import okio.Buffer;

import java.io.ByteArrayOutputStream;

import static java.nio.charset.StandardCharsets.UTF_8;


public class BusinessHandler extends ChannelInboundHandlerAdapter {
    private static LongLinkHeader longLinkHeader = new LongLinkHeader();
    private static LongLinkMessage longLinkMessage = new LongLinkMessage();

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        System.out.println("====================channelRead===================");
    }

    @Override
    public void channelActive(ChannelHandlerContext ctx) {
        ctx.fireChannelActive();
        Buffer buffer = new Buffer();
        DefaultChannelId channelId = (DefaultChannelId) ctx.channel().id();

        ByteArrayOutputStream loginStream = new ByteArrayOutputStream();

        String token = "12345678";
        buffer.writeByte(0xEF);
        buffer.writeByte(0);
        buffer.writeByte(0);
        buffer.writeByte(token.length());
        buffer.write(token.getBytes(UTF_8));
        buffer.writeShort(10);
        buffer.writeLong(112233);
        buffer.writeShort(11);
        byte[] bytes = loginStream.toByteArray();
        buffer.writeInt(bytes.length);
        buffer.writeInt(bytes.length);

        ChannelFuture future = ctx.channel().writeAndFlush(buffer);
        future.addListener((GenericFutureListener) future1 -> {
            if (future1.isSuccess()) {
                System.out.println("send success");
            } else {
                System.out.println("send error");
            }
        });

        System.out.println("====================channelActive222===================");
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        System.out.println("====================exceptionCaught===================");
        cause.printStackTrace();
        ctx.close();
    }
}
