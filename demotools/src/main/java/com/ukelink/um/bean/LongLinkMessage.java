package com.ukelink.um.bean;

import io.netty.buffer.ByteBuf;
import lombok.Data;

@Data
public class LongLinkMessage {
    LongLinkHeader longLinkHeader;
    ByteBuf longLinkBody;
}
