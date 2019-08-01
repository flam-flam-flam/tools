package com.ukelink.um.service;

import com.ukelink.um.handle.netty.NettyClientInit;
import io.netty.bootstrap.Bootstrap;
import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;
import io.netty.util.concurrent.Future;
import io.netty.util.concurrent.GenericFutureListener;
import org.springframework.stereotype.Component;

@Component
public class NettyService {
    public static void init() throws Exception {
        System.out.println("netty init");
        EventLoopGroup group = new NioEventLoopGroup();
        try {
            Bootstrap b = new Bootstrap();
            b.group(group)
                    .channel(NioSocketChannel.class)
                    .handler(new NettyClientInit());

//            Channel ch = b.connect("10.100.93.47",8801).sync().channel();
            ChannelFuture ch = b.connect("10.100.93.47", 8801);
            ch.addListener(future -> {
                if (future.isSuccess()) {
                    System.out.println("connect success");
                } else {
                    System.out.println("connect error");
                }
            });

//            ch.addListener(new GenericFutureListener<Future<? super Void>>() {
//                @Override
//                public void operationComplete(Future<? super Void> future) throws Exception {
//                    if (future.isSuccess()) {
//                        System.out.println("connect success");
//                    } else {
//                        System.out.println("connect error");
//                    }
//                }
//            });
//            ch.channel().closeFuture().sync();
//            ch.close();
//        } catch (InterruptedException e) {
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
//            group.shutdownGracefully();
        }
    }
}
