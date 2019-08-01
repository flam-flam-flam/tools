package com.ukelink.imservice.server;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import org.springframework.stereotype.Component;

@Component
public class NettyServer {
    //static final int PORT = Integer.parseInt(System.getProperty("port", "8463"));

    public static void  init()
    {
        System.out.println("hello world init000  ------------------===============================");
        EventLoopGroup bossGroup = new NioEventLoopGroup();
        EventLoopGroup workerGroup = new NioEventLoopGroup();
        try {
            ServerBootstrap b = new ServerBootstrap();
            b.group(bossGroup, workerGroup)
                    .channel(NioServerSocketChannel.class)
                   // .handler(new LoggingHandler(LogLevel.INFO))
                    .childHandler(new NettyServerInitializer());

            ChannelFuture channelFuture = b.bind(8083).sync();
            channelFuture.channel().closeFuture().sync();
           //b.bind(8083).sync().channel();
        } catch (InterruptedException e) {
            e.printStackTrace();
        } finally {
            bossGroup.shutdownGracefully();
            workerGroup.shutdownGracefully();
        }
    }

}
