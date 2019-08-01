package com.ukelink.um.handle.netty;

import com.ukelink.um.bean.LongLinkMessage;
import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.ByteToMessageCodec;
import io.netty.handler.codec.ByteToMessageDecoder;
import io.netty.handler.codec.MessageToByteEncoder;
import okio.Buffer;

import java.nio.charset.Charset;
import java.util.List;

import static java.nio.charset.StandardCharsets.UTF_8;

public class EncodeHandler extends MessageToByteEncoder<Buffer> {
    public void encode(ChannelHandlerContext ctx, Buffer msg, ByteBuf out) throws Exception {
        //out.add(doEncode(ctx.alloc(), msg));
        System.out.println("2222 endcode");
//        buffer.writeShort(12);
//        buffer.writeShort(18);
//        buffer.writeInt(40);
//        buffer.writeShort(11);
//        buffer.writeInt(22);
//        buffer.writeString("1111", Charset.forName("utf-8"));
        out.writeInt(40);
        out.writeShort(18);
        out.writeShort(11);
        out.writeShort(3);
        out.writeInt(22);
        out.writeBytes("1111".getBytes(UTF_8));
//        out.
        //out.writeBytes(msg.getLongLinkBody());
    }
}
